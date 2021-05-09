package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

// TypeTokenAdd 创建用户访问token
const (
	TypeTokenAdd = "token:add"
)

// NewTokenAddTask 初始化token创建任务
func NewTokenAddTask(caseID, placeID, WorkflowID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id":     caseID,
		"place_id":    placeID,
		"workflow_id": WorkflowID,
	}
	return cmdx.NewTask(TypeTokenAdd, payload)
}

// HandleTokenAddTask token创建逻辑处理
func HandleTokenAddTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	placeID, err := t.Payload.GetInt("place_id")
	if err != nil {
		return err
	}
	WorkflowID, err := t.Payload.GetInt("workflow_id")
	fmt.Println(caseID, placeID)
	token := &model.WfToken{
		WorkflowID: WorkflowID,
		CaseID:     caseID,
		PlaceID:    placeID,
	}
	token.Create()
	return nil
}
