package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
	"time"
)

const (
	TypeTokenLock = "token:lock"
)

func NewTokenLockTask(caseID, placeID, workitemID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id":     caseID,
		"place_id":    placeID,
		"workitem_id": workitemID,
	}
	return cmdx.NewTask(TypeTokenLock, payload)
}

func HandleTokenLockTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	placeID, err := t.Payload.GetInt("place_id")
	if err != nil {
		return err
	}
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Println(caseID, placeID, workitemID)
	var token = &model.WfToken{}
	token, err = token.FindBy(model.NewSqlCnd().Where("case_id =? and state =? and place_id", caseID, 0, placeID))
	if err != nil {
		return err
	}
	token.Updates(map[string]interface{}{
		"state": 1,
		"locked_at": time.Now(),
		"locked_workitem_id": workitemID,
	})
	return nil
}
