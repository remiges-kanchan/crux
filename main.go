package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remiges-tech/alya/service"
	"github.com/remiges-tech/alya/wscutils"
	"github.com/remiges-tech/crux/db"
	"github.com/remiges-tech/crux/db/sqlc-gen"
	crux "github.com/remiges-tech/crux/matching-engine"
	"github.com/remiges-tech/crux/server/app"
	"github.com/remiges-tech/crux/server/auth"
	"github.com/remiges-tech/crux/server/config"
	"github.com/remiges-tech/crux/server/realmslice"
	"github.com/remiges-tech/crux/server/schema"
	"github.com/remiges-tech/crux/server/wfinstance"
	"github.com/remiges-tech/crux/server/workflow"
	"github.com/remiges-tech/logharbour/logharbour"
	"github.com/remiges-tech/rigel"
	"github.com/remiges-tech/rigel/etcd"
)

func main() {

	// logger setup
	fallbackWriter := logharbour.NewFallbackWriter(os.Stdout, os.Stdout)
	lctx := logharbour.NewLoggerContext(logharbour.Debug0)
	l := logharbour.NewLogger(lctx, "crux", fallbackWriter)

	rigelAppName := flag.String("appName", "crux", "The name of the application")
	rigelModuleName := flag.String("moduleName", "WFE", "The name of the module")
	rigelVersionNumber := flag.Int("versionNumber", 1, "The number of the version")
	rigelConfigName := flag.String("configName", "devConfig", "The name of the configuration")
	etcdEndpoints := flag.String("etcdEndpoints", "localhost:2379", "Comma-separated list of etcd endpoints")

	flag.Parse()
	// Create a new EtcdStorage instance
	etcdStorage, err := etcd.NewEtcdStorage([]string{*etcdEndpoints})
	if err != nil {
		l.LogActivity("Error while Creating new instance of EtcdStorage", err)
		log.Fatalf("Failed to create EtcdStorage: %v", err)
	}
	l.LogActivity("Creates a new instance of EtcdStorage with endpoints", "localhost:2379")

	// Create a new Rigel instance
	rigel := rigel.New(etcdStorage, *rigelAppName, *rigelModuleName, *rigelVersionNumber, *rigelConfigName)
	l.LogActivity("Creates a new instance of rigel", rigel)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbHost, err := rigel.GetString(ctx, "db_host")
	if err != nil {
		l.LogActivity("Error while getting data from rigel", err)
		log.Fatalf("Failed to get data from rigel: %v", err)
	}
	dbPort, err := rigel.GetInt(ctx, "db_port")
	if err != nil {
		l.LogActivity("Error while getting data from rigel", err)
		log.Fatalf("Failed to get data from rigel: %v", err)
	}
	dbUser, err := rigel.GetString(ctx, "db_user")
	if err != nil {
		l.LogActivity("Error while getting data from rigel", err)
		log.Fatalf("Failed to get data from rigel: %v", err)
	}
	dbPassword, err := rigel.GetString(ctx, "db_password")
	if err != nil {
		l.LogActivity("Error while getting data from rigel", err)
		log.Fatalf("Failed to get data from rigel: %v", err)
	}
	dbName, err := rigel.GetString(ctx, "db_name")
	if err != nil {
		l.LogActivity("Error while getting data from rigel", err)
		log.Fatalf("Failed to get data from rigel: %v", err)
	}
	appServerPort, err := rigel.GetInt(ctx, "app_server_port")
	if err != nil {
		l.LogActivity("Error while getting data from rigel", err)
		log.Fatalf("Failed to get data from rigel: %v", err)
	}
	l.Log("Retrieves the configuration data from rigel")

	// Database connection
	connURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	connPool, err := db.NewProvider(connURL)
	if err != nil {
		l.LogActivity("Error while establishes a connection with database", err)
		log.Fatalln("Failed to establishes a connection with database", err)
	}
	queries := sqlc.New(connPool)

	// load all rulesets and schema's in cache
	err = crux.Load(queries, ctx)
	if err != nil {
		fmt.Println("error at loading cache")
		return
	}

	// Define a custom validation tag-to-message ID map
	customValidationMap := map[string]int{
		"required":  101,
		"gt":        102,
		"alpha":     103,
		"lowercase": 104,
		"max":       105,
		"lt":        106,
	}
	// Custom validation tag-to-error code map
	customErrCodeMap := map[string]string{
		"required":  "required",
		"gt":        "greater",
		"alpha":     "alphabet",
		"lowercase": "lowercase",
		"max":       "exceed the maximum value allowed",
		"lt":        "exceed the limit value allowed",
	}
	// Register the custom map with wscutils
	wscutils.SetValidationTagToMsgIDMap(customValidationMap)
	wscutils.SetValidationTagToErrCodeMap(customErrCodeMap)

	// Set default message ID and error code if needed
	wscutils.SetDefaultMsgID(100)
	wscutils.SetDefaultErrCode("validation_error")

	// router
	r := gin.Default()

	// r.Use(corsMiddleware())

	// schema services
	s := service.NewService(r).
		WithLogHarbour(l).
		WithDatabase(connPool).
		WithDependency("queries", queries)

	apiV1Group := r.Group("/api/v1/")

	// Schema
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/wfschemaget", schema.SchemaGet)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodDelete, "/wfschemadelete", schema.SchemaDelete)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/wfschemaList", schema.SchemaList)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/wfschemaNew", schema.SchemaNew)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPut, "/wfschemaUpdate", schema.SchemaUpdate)
	// Workflow
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/workflowget", workflow.WorkflowGet)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/workflowlist", workflow.WorkflowList)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/workflowNew", workflow.WorkFlowNew)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPut, "/workflowUpdate", workflow.WorkFlowUpdate)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodDelete, "/workflowdelete", workflow.WorkflowDelete)
	//wfinstance
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/wfinstancenew", wfinstance.GetWFinstanceNew)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/wfinstanceabort", wfinstance.GetWFInstanceAbort)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/wfinstancelist", wfinstance.GetWFInstanceList)
	//app
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/appnew", app.AppNew)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPut, "/appupdate", app.AppUpdate)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/appdelete/:name", app.AppDelete)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodGet, "/applist", app.AppList)

	// Realm-slice management
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/realmSliceNew", realmslice.RealmSliceNew)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/realmsliceactivate", realmslice.RealmSliceActivate)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/realmslicedeactivate", realmslice.RealmSliceDeactivate)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodGet, "/realmSliceApps/:id", realmslice.RealmSliceApps)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/RealmSlicePurge", realmslice.RealmSlicePurge)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodGet, "/realmslicelist", realmslice.RealmSliceList)
	// auth
	s.RegisterRouteWithGroup(apiV1Group, http.MethodGet, "/useractivate/:userid", auth.UserActivate)
	s.RegisterRouteWithGroup(apiV1Group, http.MethodGet, "/userdeactivate/:userid", auth.UserDeactivate)
	// config management
	s.RegisterRouteWithGroup(apiV1Group, http.MethodPost, "/configset", config.ConfigSet)

	appServerPortStr := strconv.Itoa(appServerPort)
	r.Run(":" + appServerPortStr)
	if err != nil {
		l.LogActivity("Failed to start server", err)
		log.Fatalf("Failed to start server: %v", err)
	}

}
