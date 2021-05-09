package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"time"
)

const (
	TypeTokenConsume = "token:consume"
)

func NewTokenConsumeTask(caseID, placeID int, lockedItemID ...int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id":        caseID,
		"place_id":       placeID,
		"locked_item_id": lockedItemID[0],
	}
	return cmdx.NewTask(TypeTokenConsume, payload)
}

// HandleTokenConsumeTask token 消耗
func HandleTokenConsumeTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	placeID, err := t.Payload.GetInt("place_id")
	if err != nil {
		return err
	}
	lockedItemID, err := t.Payload.GetInt("locked_item_id")
	if err != nil {
		return err
	}
	var token = &model.WfToken{}
	var workitem = &model.WfWorkitem{}
	workitem, _ = workitem.FindBy(model.NewSqlCnd().Where("id =?", lockedItemID))
	// TODO: transaction
	// 如果存在workitem
	token.ConsumedAt = time.Now()
	token.State = 3
	var condition = map[string]interface{}{}
	if err == nil {
		condition = map[string]interface{}{
			"case_id":            caseID,
			"place_id":           placeID,
			"state":              1,
			"locked_workitem_id": lockedItemID,
		}
		token.Updates(condition)
	} else {
		freeToken, _ := token.FindBy(model.NewSqlCnd().Where("case_id =? and place_id =? and state =?", caseID, placeID, 0))
		condition = map[string]interface{}{
			"id": freeToken.ID,
		}
		token.Updates(condition)
	}
	return nil
}
