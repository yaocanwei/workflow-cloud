package model

import (
	db "flowpipe-server/core/db"
	"time"
)

// WfTranstion 流程变迁, 用于控制变迁发生方式和读写变量
type WfTranstion struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description" gorm:"type:text"`
	WorkflowID  int    `json:"workflow_id"`
	SortOrder   int    `json:"sort_order"`
	// use with timed trigger, after x minitues, trigger exec
	TriggerLimit int `json:"trigger_limit"`
	// 0-user,1-automatic, 2-message,3-time
	TriggerType int `json:"trigger_type"`
	FormID      int `json:"form_id"`
	// func EnableDefault(); default: EnableDefault
	EnableCallback string `json:"enable_callback"`
	// func FireDefault(); default: FireDefault
	FireCallback string `json:"fire_callback"`
	// func NotificationDefault(); default: NotificationDefault
	NotificationCallback string `json:"notification_callback"`
	// func TimeDefault(); default: TimeDefault
	TimeCallback string `json:"time_callback"`
	// func DeadlineDefault(); default: DeadlineDefault
	DeadlineCallback string `json:"deadline_callback"`
	// func HoldTimeoutDefault(); default: HoldTimeoutDefault
	HoldTimeoutCallback string `json:"hold_timeout_callback"`
	// func AssignmentDefault(); default: AssignmentDefault
	AssignmentCallback string `json:"assignment_callback"`
	// func UnassignmentDefault(); Default: UnassignmentDefault
	UnassignmentCallback string `json:"unassignment_callback"`
	FormType             string `json:"form_type"`
	// SubWorkflowID string `json:"sub_workflow_id"`
	MultipleInstance  bool      `json:"multiple_instance"`
	FinishCondition   string    `json:"finish_condition"`
	DynamicAssignByID int       `json:"dynamic_assign_by_id"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
	Arcs []WfArc `gorm:"foreignkey:TransitionID;association_foreignkey:ID"`
	TransitionStaticAssignments []WfTransitionStaticAssignment
}

// Create 创建transition
func (wfTransition *WfTranstion) Create() (*WfTranstion, error) {
	err := db.WfDb.Create(&wfTransition).Error
	return wfTransition, err
}

// Delete 删除transition
func (wfTransition *WfTranstion) Delete(ID int) error {
	err := db.WfDb.Where("id = ?", ID).Delete(&wfTransition).Error
	return err
}

// Updates 更新transition
func (wfTransition *WfTranstion) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(&wfTransition).Error
}

/**
 * @Author canweiyao
 * @Description 根据条件查找记录
 * @Date 8:31 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfTransition *WfTranstion) FindBy(cond *SqlCond) (*WfTranstion, error) {
	err := cond.FindBy(db.WfDb, wfTransition)
	if err != nil {
		return &WfTranstion{}, err
	}
	return wfTransition, nil
}

/**
 * @Author canweiyao
 * @Description  返回transitions数组
 * @Date 8:29 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfTransition *WfTranstion) Find(cond *SqlCond) ([]WfTranstion, error) {
	wfTransitionList := make([]WfTranstion, 0)
	err := cond.Find(db.WfDb, &wfTransitionList)
	return wfTransitionList, err
}

// ExplicitOrSplit 判断是否走不同分支
func (wfTransition *WfTranstion) ExplicitOrSplit() {
	db.WfDb.Preload("Arcs", "direction = ?", 1).Find(&wfTransition)
}

// Count 统计transition
func (wfTransition *WfTranstion) Count() (count int, err error) {
	if err = db.WfDb.Model(&wfTransition).Unscoped().Count(&count).Error; err != nil {
		return
	}
	return
}
