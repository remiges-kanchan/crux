// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	ActivateRecord(ctx context.Context, arg ActivateRecordParams) error
	AddWFNewInstances(ctx context.Context, arg AddWFNewInstancesParams) ([]Wfinstance, error)
	AllRuleset(ctx context.Context) ([]Ruleset, error)
	AllSchemas(ctx context.Context) ([]Schema, error)
	AppDelete(ctx context.Context, arg AppDeleteParams) error
	AppExist(ctx context.Context, app string) (int32, error)
	AppNew(ctx context.Context, arg AppNewParams) ([]App, error)
	AppUpdate(ctx context.Context, arg AppUpdateParams) error
	CapGet(ctx context.Context, arg CapGetParams) ([]CapGetRow, error)
	CapList(ctx context.Context, arg CapListParams) ([]CapListRow, error)
	CloneRecordInConfigBySliceID(ctx context.Context, arg CloneRecordInConfigBySliceIDParams) (pgconn.CommandTag, error)
	CloneRecordInRealmSliceBySliceID(ctx context.Context, arg CloneRecordInRealmSliceBySliceIDParams) (int32, error)
	CloneRecordInRulesetBySliceID(ctx context.Context, arg CloneRecordInRulesetBySliceIDParams) (pgconn.CommandTag, error)
	CloneRecordInSchemaBySliceID(ctx context.Context, arg CloneRecordInSchemaBySliceIDParams) (pgconn.CommandTag, error)
	ConfigGet(ctx context.Context, realm string) ([]ConfigGetRow, error)
	ConfigSet(ctx context.Context, arg ConfigSetParams) error
	DeactivateRecord(ctx context.Context, arg DeactivateRecordParams) error
	DeleteCapGranForApp(ctx context.Context, arg DeleteCapGranForAppParams) error
	DeleteWFInstanceListByParents(ctx context.Context, arg DeleteWFInstanceListByParentsParams) ([]Wfinstance, error)
	DeleteWFInstances(ctx context.Context, arg DeleteWFInstancesParams) error
	DeleteWfinstanceByID(ctx context.Context, arg DeleteWfinstanceByIDParams) ([]Wfinstance, error)
	GetApp(ctx context.Context, arg GetAppParams) (string, error)
	GetAppList(ctx context.Context, realm string) ([]GetAppListRow, error)
	GetAppName(ctx context.Context, arg GetAppNameParams) ([]App, error)
	GetCapGrantForApp(ctx context.Context, arg GetCapGrantForAppParams) ([]Capgrant, error)
	GetClass(ctx context.Context, arg GetClassParams) (string, error)
	GetRealmSliceListByRealm(ctx context.Context, realm string) ([]GetRealmSliceListByRealmRow, error)
	GetSchemaWithLock(ctx context.Context, arg GetSchemaWithLockParams) (GetSchemaWithLockRow, error)
	GetUserRealm(ctx context.Context, userid string) ([]string, error)
	GetWFActiveStatus(ctx context.Context, arg GetWFActiveStatusParams) (pgtype.Bool, error)
	GetWFINstance(ctx context.Context, arg GetWFINstanceParams) (int64, error)
	GetWFInstanceCounts(ctx context.Context, arg GetWFInstanceCountsParams) (int64, error)
	GetWFInstanceCurrent(ctx context.Context, arg GetWFInstanceCurrentParams) (Wfinstance, error)
	GetWFInstanceFromId(ctx context.Context, id int32) (Wfinstance, error)
	GetWFInstanceList(ctx context.Context, arg GetWFInstanceListParams) ([]Wfinstance, error)
	GetWFInstanceListByParents(ctx context.Context, id []int32) ([]Wfinstance, error)
	GetWFInstanceListForMarkDone(ctx context.Context, arg GetWFInstanceListForMarkDoneParams) ([]Wfinstance, error)
	GetWFInternalStatus(ctx context.Context, arg GetWFInternalStatusParams) (bool, error)
	GetWorkflowNameForStep(ctx context.Context, step string) (GetWorkflowNameForStepRow, error)
	GrantAppCapability(ctx context.Context, arg GrantAppCapabilityParams) error
	GrantRealmCapability(ctx context.Context, arg GrantRealmCapabilityParams) error
	InsertNewRecordInRealmSlice(ctx context.Context, arg InsertNewRecordInRealmSliceParams) (int32, error)
	RealmSliceActivate(ctx context.Context, arg RealmSliceActivateParams) (Realmslice, error)
	RealmSliceAppsList(ctx context.Context, id int32) ([]RealmSliceAppsListRow, error)
	RealmSliceDeactivate(ctx context.Context, arg RealmSliceDeactivateParams) (Realmslice, error)
	RealmSlicePurge(ctx context.Context, realm string) (pgconn.CommandTag, error)
	RulesetRowLock(ctx context.Context, arg RulesetRowLockParams) (Ruleset, error)
	SchemaDelete(ctx context.Context, id int32) (int32, error)
	SchemaGet(ctx context.Context, arg SchemaGetParams) ([]SchemaGetRow, error)
	SchemaNew(ctx context.Context, arg SchemaNewParams) (int32, error)
	SchemaUpdate(ctx context.Context, arg SchemaUpdateParams) error
	UpdateCapGranForUser(ctx context.Context, userid string) error
	UpdateWFInstanceDoneat(ctx context.Context, arg UpdateWFInstanceDoneatParams) error
	UpdateWFInstanceStep(ctx context.Context, arg UpdateWFInstanceStepParams) error
	UserActivate(ctx context.Context, arg UserActivateParams) (Capgrant, error)
	UserDeactivate(ctx context.Context, arg UserDeactivateParams) (Capgrant, error)
	WfPatternSchemaGet(ctx context.Context, arg WfPatternSchemaGetParams) ([]byte, error)
	WfSchemaGet(ctx context.Context, arg WfSchemaGetParams) (Schema, error)
	WfSchemaList(ctx context.Context, arg WfSchemaListParams) ([]WfSchemaListRow, error)
	Wfschemadelete(ctx context.Context, arg WfschemadeleteParams) error
	Wfschemaget(ctx context.Context, arg WfschemagetParams) (WfschemagetRow, error)
	WorkFlowNew(ctx context.Context, arg WorkFlowNewParams) error
	WorkFlowUpdate(ctx context.Context, arg WorkFlowUpdateParams) (pgconn.CommandTag, error)
	WorkflowDelete(ctx context.Context, arg WorkflowDeleteParams) (pgconn.CommandTag, error)
	WorkflowList(ctx context.Context, arg WorkflowListParams) ([]WorkflowListRow, error)
	Workflowget(ctx context.Context, arg WorkflowgetParams) (WorkflowgetRow, error)
}

var _ Querier = (*Queries)(nil)
