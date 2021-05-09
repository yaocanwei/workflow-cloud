package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeCaseCancel = "case:cancel"
)

func NewCaseCancelTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{"case_id": caseID}
	return cmdx.NewTask(TypeWorkitemCancel, payload)
}

// HandleCaseCancelTask
func HandleCaseCancelTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	fmt.Print(caseID)
	var wfCase *model.WfCase
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =?", caseID))
	if err != nil {
		return err
	}
	if wfCase.State != 2 && wfCase.State != 1 {
		panic("在用或暂停的实例才能取消")
	}
	wfCase.ChangeState("canceled")
	return nil
}
