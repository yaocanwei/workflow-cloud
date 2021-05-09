package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeCaseSuspend = "case:suspend"
)

// NewCaseSuspendTask NewCaseSuspendTask
func NewCaseSuspendTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id": caseID,
	}
	return cmdx.NewTask(TypeCaseSuspend, payload)
}

/**
 * @Author canweiyao
 * @Description //TODO 
 * @Date 3:36 AM 2021/3/26
 * @Param 
 * @return 
 **/
func HandleCaseSuspendTask(ctx context.Context, t *cmdx.Task) error {
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
	if wfCase.State != 1 {
		panic("Only active or suspended case can be canceled")
	}
	err = wfCase.Updates(map[string]interface{}{"state": 2})
	if err != nil {
		return err
	}
	return nil
}
