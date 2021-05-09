package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
	"time"
)

const (
	TypeWorkitemFinish = "workitem:finish"
)

func NewWorkitemFinishTask(workitemID int) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
	}
	return cmdx.NewTask(TypeWorkitemFinish, payload)
}

func HandleWorkitemFinishTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Println(workitemID)
	var workitem = &model.WfWorkitem{}
	workitem, err = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if err != nil {
		return err
	}
	if workitem.Forked {
		workitem.Updates(map[string]interface{}{
			"finished_at": time.Now(),
			"state":       3,
		})
		// TODO: use self join
		// workitem.Parent.Updates(map[string]interface{}{
		// 	"children_finished_count": workitem.Parent.ChildrenFinishedCount + 1,
		// })
		var parentWorkitem = &model.WfWorkitem{}
		parentWorkitem, err = parentWorkitem.FindBy(model.NewSqlCnd().Where("parent_id =?", workitem.ParentID))
		if err != nil {
			return err
		}
		parentWorkitem.Updates(map[string]interface{}{
			"children_finished_count": parentWorkitem.ChildrenFinishedCount + 1,
		})
		// if parentWorkitem.ChildrenFinishedCount >= parentWorkitem.ChildrenCount || workitem.Transition.FinishCondition
	}
	return nil
}
