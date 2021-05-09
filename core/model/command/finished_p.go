package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
)

const (
	TypePFinish = "p:finish"
)

func NewPFinishTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id": caseID,
	}
	return cmdx.NewTask(TypePFinish, payload)
}

// return bool and error
func HandlePFinishTask(ctx context.Context, t *cmdx.Task) (bool, error) {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return false, err
	}
	var wfCase = &model.WfCase{}
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =?", caseID))
	if wfCase.State == 4 {
		return true, nil
	}
	var place = &model.WfPlace{}
	endPlace, err := place.FindBy(model.NewSqlCnd().Where("workflow_id =? and palce_type =?", wfCase.WorkflowID, 2))
	if err != nil {
		return false, err
	}
	var token = &model.WfToken{}
	endPlaceToken, err := token.Find(model.NewSqlCnd().Where("case_id =? and place_id =?", caseID, endPlace.ID))
	if err != nil {
		return false, err
	}
	if len(endPlaceToken) == 0 {
		return false, nil
	}
	freeAndLockedTokens, err := token.Find(model.NewSqlCnd().Where("place_id =? and case_id =? and state in(?)", endPlace.ID, caseID, []int{0, 1}))
	if err != nil {
		return false, err
	}
	if len(freeAndLockedTokens) > 0 {
		panic("The workflow net is misconstructed: Some parallel executions have not finished.")
	}
	task := NewTokenConsumeTask(caseID, endPlace.ID)
	HandleTokenConsumeTask(context.Background(), task)
	if wfCase.State != 4 {
		wfCase.ChangeState("finished")
	    //	TODO: add sub workflow instance
	}
	return true, nil
}
