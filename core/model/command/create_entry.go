package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
	"strconv"
)

const (
	TypeEntryCreate = "entry:create"
)

func NewEntryCreateTask(formID, workitemID, userID int, params map[string]interface{}) *cmdx.Task {
	payload := map[string]interface{}{
		"form_id":     formID,
		"workitem_id": workitemID,
		"user_id":     userID,
		"params":      params,
	}
	return cmdx.NewTask(TypeTokenConsume, payload)
}

// HandleEntryCreateTask 处理创建entry逻辑
func HandleEntryCreateTask(ctx context.Context, t *cmdx.Task) error {
	formID, err := t.Payload.GetInt("form_id")
	if err != nil {
		return err
	}
	workitemID, err := t.Payload.GetInt("workitem_id")
	userID, err := t.Payload.GetInt("user_id")
	params, err := t.Payload.GetStringMap("params")
	fmt.Println(formID, workitemID, userID, params)
	var entry = &model.WfEntry{}
	entry, err = entry.FindOrCreate(model.NewSqlCnd().Where("form_id =? and user_id =? and workitem_id =?", formID, userID, workitemID))
	if err != nil {
		return err
	}
	var wfFieldValue = &model.WfFieldValue{}
	for fieldID, fieldValue := range params {
		wfFieldValue, err = wfFieldValue.FindBy(model.NewSqlCnd().Where("form_id =? and field_id =? and entry_id =?", formID, fieldID, entry.ID))
		wfFieldValue.Value = fieldValue.(string)
		if err == nil {
			wfFieldValue.Updates(map[string]interface{}{"id": wfFieldValue.ID})
		} else {
			wfFieldValue.FormID = formID
			fieldIDFormat, _ := strconv.Atoi(fieldID)
			wfFieldValue.FieldID = fieldIDFormat
			wfFieldValue.EntryID = entry.ID
			wfFieldValue.Create()
		}
	}
	// entry.Updates(map[string]interface{}{"payload": })
	// TODO: update payload
	return nil
}
