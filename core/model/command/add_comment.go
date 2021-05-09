package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
)

// TypeCommentAdd 定义添加评论
const (
	TypeCommentAdd = "comment:add"
)

// NewCommentAddTask 初始化评论任务
func NewCommentAddTask(worktemID, userID int, body string) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": worktemID,
		"body":        body,
		"user_id":     userID,
	}
	return cmdx.NewTask(TypeCommentAdd, payload)
}

// HandleCommentAddTask 评论逻辑处理
func HandleCommentAddTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	body, err := t.Payload.GetString("body")
	if err != nil {
		return err
	}
	userID, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	comment := &model.WfComment{
		WorkitemID: workitemID,
		UserID:     userID,
		Body:       body,
	}
	comment.Create()
	return nil
}
