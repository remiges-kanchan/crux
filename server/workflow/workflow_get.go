package workflow

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/remiges-tech/alya/service"
	"github.com/remiges-tech/alya/wscutils"
	"github.com/remiges-tech/crux/db/sqlc-gen"
	"github.com/remiges-tech/crux/server"
)

// WorkflowGet will be responsible for processing the /workflowget request that comes through as a POST
func WorkflowGet(c *gin.Context, s *service.Service) {
	lh := s.LogHarbour
	lh.Log("WorkflowGet request received")

	// var response schemaGetResp
	var request WorkflowGetReq
	err := wscutils.BindJSON(c, &request)
	if err != nil {
		lh.LogActivity("error while binding json request error:", err.Error)
		return
	}

	valError := wscutils.WscValidate(request, func(err validator.FieldError) []string { return []string{} })
	if len(valError) > 0 {
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, valError))
		lh.LogActivity("validation error:", valError)
		return
	}

	query, ok := s.Dependencies["queries"].(*sqlc.Queries)
	if !ok {
		lh.Log("Error while getting query instance from service Dependencies")
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_InternalErr, server.ErrCode_DatabaseError))
		return
	}

	dbResponse, err := query.Workflowget(c, sqlc.WorkflowgetParams{
		Slice:   request.Slice,
		App:     request.App,
		Class:   request.Class,
		Setname: request.Name,
	})
	if err != nil {
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(server.MsgId_Invalid, server.ErrCode_Invalid, nil)}))
		lh.LogActivity("failed to get data from DB:", err.Error)
		return
	}

	tempData := responseBinding(dbResponse)

	err = json.Unmarshal(dbResponse.Flowrules, &tempData.Flowrules)
	if err != nil {
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(server.MsgId_Invalid, server.ErrCode_Invalid, nil)}))
		lh.LogActivity("failed to unmarshal data:", err.Error)
		return
	}
	lh.Log(fmt.Sprintf("Record found: %v", map[string]any{"response": tempData}))
	wscutils.SendSuccessResponse(c, wscutils.NewSuccessResponse(tempData))
}

func responseBinding(dbResponse sqlc.WorkflowgetRow) WorkflowgetRow {
	tempData := WorkflowgetRow{
		ID:         dbResponse.ID,
		Slice:      dbResponse.Slice,
		App:        dbResponse.App,
		Class:      dbResponse.Class,
		Name:       dbResponse.Name,
		IsActive:   dbResponse.IsActive.Bool,
		IsInternal: dbResponse.IsInternal,
		Createdat:  dbResponse.Createdat,
		Createdby:  dbResponse.Createdby,
		Editedat:   dbResponse.Editedat,
		Editedby:   dbResponse.Editedby,
	}
	return tempData
}