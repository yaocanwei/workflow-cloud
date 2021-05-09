package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeCaseStart = "case:start"
)

// NewCaseStartTask NewCaseStartTask
func NewCaseStartTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id": caseID,
	}
	return cmdx.NewTask(TypeCaseStart, payload)
}

// HandleCaseStartTask HandleCaseStartTask
func HandleCaseStartTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	fmt.Println(caseID)
	var wfCase = &model.WfCase{}
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =?", caseID))
	err = wfCase.ChangeState("active")
	if err != nil {
		return err
	}
	var place = &model.WfPlace{}
	place, err = place.FindBy(model.NewSqlCnd().Where("workflow_id =? and palce_type =?", wfCase.WorkflowID, 0))
	if err != nil {
		return err
	}
	task := NewTokenAddTask(caseID, place.ID, wfCase.WorkflowID)
	err = HandleTokenAddTask(context.Background(), task)
	if err != nil {
		return err
	}
	return nil
}
