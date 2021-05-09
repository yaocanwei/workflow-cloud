package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeCaseResume = "case:resume"
)

func NewCaseResumeTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id": caseID,
	}
	return cmdx.NewTask(TypeCaseResume, payload)
}

func HandleCaseResumeTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	fmt.Println(caseID)
	var wfCase = &model.WfCase{}
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =?", caseID))
	if err != nil {
		return err
	}
	if wfCase.State != 2 && wfCase.State != 3 {
		panic("Only suspended or canceled cases can be resumed")
	}
	err = wfCase.ChangeState("active")
	return err
}
