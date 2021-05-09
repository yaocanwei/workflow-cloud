package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
	"time"
)

const (
	TypeTokenRelease = "token:release"
)

func NewTokenReleaseTask(workitemID int) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
	}
	return cmdx.NewTask(TypeTokenLock, payload)
}

func HandleTokenReleaseTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Println(workitemID)
	var token = &model.WfToken{}
	lockedTokens, err := token.Find(model.NewSqlCnd().Where("locked_workitem_id =? and state =?", workitemID, 1))
	if err != nil {
		return err
	}
	for _, lockedToken := range lockedTokens {
		task := NewTokenAddTask(lockedToken.CaseID, lockedToken.PlaceID, lockedToken.WorkflowID)
		HandleTokenAddTask(context.Background(), task)
		lockedToken.Updates(map[string]interface{}{
			"state": 2,
			"canceled_at": time.Now(),
		})
	}
	return nil
}
