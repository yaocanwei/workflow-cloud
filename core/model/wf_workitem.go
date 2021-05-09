package model

import (
	"flowpipe-server/core/db"
	"flowpipe-server/helper"
	"time"
)

// WfWorkitem 流程任务
type WfWorkitem struct {
	Model
	CaseID                int                    `json:"case_id"`
	Case                  WfCase                 `json:"wfcase" gorm:"FOREIGNKEY:UUID;ASSOCIATION_FOREIGNKEY:CaseID"` // 工作流实例
	WorkflowID            int                    `json:"workflow_id"`
	TransitionID          int                    `json:"transition_id"`
	Transition            WfTranstion            `json:"wftransition" gorm:"FOREIGNKEY:UUID;ASSOCIATION_FOREIGNKEY:TransitionID"`
	Parties               []WfParty              `json:"parties" gorm:"many2many:wf_workitem_assignments;"`
	WorkitemAssignments   []WfWorkitemAssignment `json:"wf_workitem_assignments"`
	State                 int                    `json:"state"` // 0-enabled, 1-started, 2-canceled, 3-finished,4-overridden
	EnabledAt             time.Time              `json:"enabled_at" gorm:"column:enabled_at"`
	StartedAt             time.Time              `json:"started_at" gorm:"column:started_at"`
	CanceledAt            time.Time              `gorm:"column:canceled_at" json:"canceled_at"`
	FinishedAt            time.Time              `gorm:"column:finished_at" json:"finished_at"`
	OverriddenAt          time.Time              `gorm:"column:overridden_at" json:"overridden_at"` // 拒绝时间
	Deadline              time.Time              `gorm:"column:deadline" json:"deadline"`
	CreatedAt             time.Time              `gorm:"column:created_at" json:"created_at"`
	UpdatedAt             time.Time              `gorm:"column:updated_at" json:"updated_at"`
	TriggerTime           time.Time              `gorm:"column:trigger_time" json:"trigger_time"` // set when transition_trigger=TIME & trigger_limit present
	HoldingUserID         int                    `json:"holding_user_id"`
	ChildrenCount         int                    `json:"children_count" gorm:"default:0"`
	ChildrenFinishedCount int                    `json:"children_finished_count" gorm:"default:0"`
	Forked                bool                   `json:"forked"`
	ParentID              int                    `json:"parent_id"`
	// Parent                *WfWorkitem            `gorm:"foreignkey:ParentID"`
	Children []*WfWorkitem `gorm:"foreignkey:ParentID"`
}

/**
 * @Author
 * @Description 获取任务
 * @Date 9:41 AM 2021/4/4
 * @Param
 * @return
 **/
func (workItem *WfWorkitem) Find(cond *SqlCond) ([]WfWorkitem, error) {
	wfWorkitemList := make([]WfWorkitem, 0)
	err := cond.Find(db.WfDb, &wfWorkitemList)
	return wfWorkitemList, err
}

/**
 * @Author canweiyao
 * @Description  返回workitem
 * @Date 6:59 PM 2021/4/4
 * @Param
 * @return
 **/
func (workItem *WfWorkitem) FindBy(cond *SqlCond) (*WfWorkitem, error) {
	err := cond.FindBy(db.WfDb, workItem)
	if err != nil {
		return &WfWorkitem{}, err
	}
	return workItem, nil
}

// HasChild 是否有子任务
func (workItem *WfWorkitem) HasChild(id int) bool {
	if err := db.WfDb.First(&workItem, "parent_id=?", id).Error; err != nil {
		return false
	}
	return true
}

// Create 创建
func (workItem *WfWorkitem) Create() (*WfWorkitem, error) {
	err := db.WfDb.Create(&workItem).Error
	return workItem, err
}

// Updates 更新任务
func (workItem *WfWorkitem) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(&workItem).Error
}

// OwnBy 判断是否任务指派人
func (workItem *WfWorkitem) OwnBy(userID int) bool {
	parties := make([]WfParty, 0)
	db.WfDb.Table("wf_parties t").Select("t.*").Joins("wf_workitem_assignments t1 ON t1.party_id = t.id left join wf_workitems t2 ON t2.id = t1.workitem_id left join wf_transitions t3 ON t3.id = t2.transition_id left join wf_cases t4 ON wf_cases.id = t2.case_id").
		Where("t3.trigger_type = ? and t4.state = ? and t2.state in(?) and t.id = ?", 0, 1, []int{1, 0}, 1).
		Find(&parties)
	userIDs := make([]int, 0)
	// TODO: flatten refactor
	for _, party := range parties {
		if len(party.Users) > 0 {
			for _, user := range party.Users {
				userIDs = append(userIDs, user.ID)
			}
		}
	}
	return helper.Contains(userIDs, userID)
}

func (workItem *WfWorkitem) Todo(userID int, state string) ([]WfWorkitem, error) {
	var err error
	workitemIDs := []int{userID}
	wfUser := &WfUser{}
	wfUser, err = wfUser.FindBy(NewSqlCnd().Where("id =?", userID))
	if err != nil {
		return nil, err
	}
	groupID := wfUser.GroupID
	if groupID != 0 {
		workitemIDs = append(workitemIDs, groupID)
	}
	joinCon := "INNER JOIN `wf_workitem_assignments ON wf_workitem_assignments.workitem_id = wf_workitems.id"
	workitems, err := workItem.Find(NewSqlCnd().Joins(joinCon).
		Where("wf_workitems.forked =? and wf_workitem_assignments.party_id in(?)", false, workitemIDs))
	if state != "" {
		workitems, err = workItem.Find(NewSqlCnd().Joins(joinCon).
			Where("wf_workitems.forked =? and wf_workitem_assignments.party_id in(?) and state =?", false, workitemIDs, state))
	}
	return workitems, err
}
