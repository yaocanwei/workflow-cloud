package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
	"time"
)

const (
	TypeTransitionsEnable = "transitions:enable"
)

func NewTransitionEnableTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{"case_id": caseID}
	return cmdx.NewTask(TypeTransitionsEnable, payload)
}

func HandleTransitionEnableTask(ctx context.Context, t *cmdx.Task) error {
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
	var workitem = &model.WfWorkitem{}
	enabledItems, _ := workitem.Find(model.NewSqlCnd().Where("case_id =? and state =?", caseID, 0))
	for _, item := range enabledItems {
		if !wfCase.CanFire(workitem.TransitionID) {
			item.Updates(map[string]interface{}{
				"state":         4,
				"overridden_at": time.Now(),
			})
		}
	}

	transitions := wfCase.Workflow.Transitions
	for _, transition := range transitions{
		workitems, err := workitem.Find(model.NewSqlCnd().Where("case_id =? and state in(?)", caseID, []int{0, 1}))
		if wfCase.CanFire(transition.ID) && (len(workitems) == 0 || err != nil ){
			continue
		}
		if transition.TriggerType == 3 && (transition.TriggerLimit != 0 ) {
			triggerTime := time.Now().Local().Add(time.Minute * time.Duration(transition.TriggerLimit))
			workitem = &model.WfWorkitem{
				WorkflowID: wfCase.WorkflowID,
				TransitionID: transition.ID,
				State: 0,
				TriggerTime: triggerTime,
			}
			workitem, _ = workitem.Create()
		//	TODO: add background
		}
	}

	return nil
}
