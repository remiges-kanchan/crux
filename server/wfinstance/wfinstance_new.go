package wfinstance

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/remiges-tech/alya/service"
	"github.com/remiges-tech/alya/wscutils"
	crux "github.com/remiges-tech/crux/matching-engine"
	"github.com/remiges-tech/crux/server"
	"github.com/remiges-tech/crux/types"
)

// Incoming request format
type WFInstanceNewRequest struct {
	Slice    int32             `json:"slice" validate:"required,gt=0,lt=85"`
	App      string            `json:"app" validate:"required,lt=15"`
	EntityID string            `json:"entityid" validate:"required,gt=0,lt=40"`
	Entity   map[string]string `json:"entity" validate:"required"`
	Workflow string            `json:"workflow" validate:"required,gt=0,lt=20"`
	Trace    *int              `json:"trace,omitempty"`
	Parent   *int32            `json:"parent,omitempty"`
}

// WFInstanceNew response format
type WFInstanceNewResponse struct {
	Tasks     []map[string]int32 `json:"tasks,omitempty"`
	Nextstep  *string            `json:"nextstep,omitempty"`
	Loggedat  pgtype.Timestamp   `json:"loggedat,omitempty"`
	Subflows  *map[string]string `json:"subflows,omitempty"`
	Tracedata *map[string]string `json:"tracedata,omitempty"`
	Done      string             `json:"done,omitempty"`
	ID        string             `json:"id,omitempty"` //wfinstance id
}

// GetWFinstanceNew will be responsible for processing the /wfinstanceNew request that comes through as a POST
func GetWFinstanceNew(c *gin.Context, s *service.Service) {
	lh := s.LogHarbour.WithClass("wfinstance")
	lh.Log("GetWFinstanceNew request received")

	var (
		wfinstanceNewreq WFInstanceNewRequest
		actionSet        crux.ActionSet
		seenRuleSets     = make(map[string]struct{})
		response         WFInstanceNewResponse
		attribute        map[string]string
		done, nextStep   string
		steps            []string
		myRuleSet        *crux.Ruleset_t
	)
	// USERID, err := server.ExtractUserNameFromJwt(c)
	// if err != nil {
	// 	lh.Info().Log("unable to extract userID from token")
	// 	wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Missing, server.ERRCode_Token_Data_Missing))
	// 	errRes := append(errRes, wscutils.BuildErrorMessage(server.MsgId_Missing, server.ERRCode_Token_Data_Missing, nil))
	// 	return false, errRes
	// }

	// REALM, err := server.ExtractRealmFromJwt(c)
	// if err != nil {
	// 	lh.Info().Log("unable to extract realm from token")
	// 	wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Missing, server.ERRCode_Token_Data_Missing))
	// 	errRes := append(errRes, wscutils.BuildErrorMessage(server.MsgId_Missing, server.ERRCode_Token_Data_Missing, nil))
	// 	return false, errRes
	// }

	isCapable, _ := server.Authz_check(types.OpReq{
		User: USERID,
	}, false)

	if !isCapable {
		lh.Info().LogActivity("unauthorized user:", USERID)
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Unauthorized, server.ErrCode_Unauthorized))
		return
	}

	// Bind request
	err := wscutils.BindJSON(c, &wfinstanceNewreq)
	if err != nil {
		lh.Error(err).Log("GetWFinstanceNew||error while binding json request error")
		return
	}
	// Standard validation of Incoming Request
	validationErrors := wscutils.WscValidate(wfinstanceNewreq, func(err validator.FieldError) []string { return []string{} })
	if len(validationErrors) > 0 {
		lh.Debug0().LogActivity("GetWFinstanceNew||validation error:", validationErrors)
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, validationErrors))
		return
	}
	// Validate request
	existingEntity := wfinstanceNewreq.Entity
	isValidReq, errStr := validateWFInstanceNewReq(wfinstanceNewreq, s, c)
	if len(errStr) > 0 || !isValidReq {
		lh.Debug0().LogActivity("GetWFinstanceNew||request validation error:", errStr)
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, errStr))
		return

	} else {
		// Additional attributes to append
		existingEntity[STEP] = START
		existingEntity[STEPFALED] = FALSE
	}
	lh.Debug0().LogActivity("wfinstanceNewRequest after adding additional attributes :", wfinstanceNewreq)

	//  doMatch() Processing

	// To get Entity
	entity := getEntityStructure(wfinstanceNewreq)

	// To get workflow rulesets from RuleSetCache
	ruleSets, err := crux.RetrieveWorkflowRulesetFromCache(REALM, wfinstanceNewreq.App, wfinstanceNewreq.Entity[CLASS], int(wfinstanceNewreq.Slice))
	if err != nil {
		lh.Error(err).Log("GetWFinstanceNew||error while getting workflow rulesets from RuleSetCache ")
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Invalid, err.Error()))
		return
	}

	for _, ruleSet := range ruleSets {
		if ruleSet.SetName == wfinstanceNewreq.Workflow {
			myRuleSet = ruleSet
		}

	}

	// call DoMatch()
	actionSet, _, err = crux.DoMatch(entity, myRuleSet, actionSet, seenRuleSets)
	if err != nil {
		lh.Error(err).Log("GetWFinstanceNew||error while calling doMatch Method")
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Invalid, err.Error()))
		return
	}

	//To verify actionSet Properties and get their values
	attribute, error := getValidPropertyAttr(actionSet)
	if error != nil {
		lh.Debug0().LogActivity("GetWFinstanceNew||error while verifying actionset properties :", error.Error())
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Invalid, server.ErrCode_Invalid_property_attributes))
		return
	}
	if attribute[DONE] == TRUE {
		done = attribute[DONE]
	} else {
		nextStep = attribute[NEXTSTEP]
	}

	if done == TRUE {
		response := make(map[string]bool)
		response[DONE] = true
		lh.Log("Finished execution of GetWFinstanceNew")
		wscutils.SendSuccessResponse(c, wscutils.NewSuccessResponse(response))
	}

	// To add records in table
	steps = actionSet.Tasks

	// if tasks of actionset contains only one task
	if len(actionSet.Tasks) == 1 && done == "" {
		addTaskRequest := AddTaskRequest{
			Steps:    steps,
			Nextstep: steps[0],
			Request:  wfinstanceNewreq,
		}
		response, err = AddTasks(addTaskRequest, s, c)
		if err != nil {
			lh.LogActivity("GetWFinstanceNew||error while adding single step in wfinstance table :", err.Error())
			wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_InternalErr, server.ErrCode_DatabaseError))
			return
		}
		lh.Log("Finished execution of GetWFinstanceNew")
		wscutils.SendSuccessResponse(c, wscutils.NewSuccessResponse(response))
		return
	}
	// if tasks of actionset contains multiple tasks
	if len(actionSet.Tasks) > 1 && done == "" {
		addTaskRequest := AddTaskRequest{
			Steps:    steps,
			Nextstep: nextStep,
			Request:  wfinstanceNewreq,
		}
		response, err = AddTasks(addTaskRequest, s, c)
		if err != nil {
			lh.LogActivity("GetWFinstanceNew||error while adding multiple steps in wfinstance table :", error.Error())
			wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_InternalErr, server.ErrCode_DatabaseError))
			return
		}
		lh.Log("Finished execution of GetWFinstanceNew")
		wscutils.SendSuccessResponse(c, wscutils.NewSuccessResponse(response))
	}

}

// To verify whether actionset.properties attributes valid and get their values
func getValidPropertyAttr(a crux.ActionSet) (map[string]string, error) {
	attribute := make(map[string]string)
	attributes := a.Properties

	isDoneOrNextStepPresent := false
	for attr, val := range attributes {
		if attr == DONE {
			attribute[attr] = val
			isDoneOrNextStepPresent = true
		} else if attr == NEXTSTEP {
			attribute[attr] = val
			isDoneOrNextStepPresent = true
		}
	}

	if !isDoneOrNextStepPresent {
		return nil, fmt.Errorf("property attributes does not contain either done or nextstep %v", attribute)
	}

	return attribute, nil
}
