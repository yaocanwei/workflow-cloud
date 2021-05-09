package model

import(
    "flowpipe-server/core/db"
    "time"
)

type WfTransitionStaticAssignment struct {
    Model
    PartyID int `json:"party_id"`
    Party   WfParty  `json:"party" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:PartyID"`
    TransitionID int `json:"transition_id"`
    WorkflowID string `json:"workflow_id"`
    CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt time.Time `json:"column:updated_at"`
    // TODO: index
}

// Create 创建
func (staticAssignment *WfTransitionStaticAssignment) Create() (*WfTransitionStaticAssignment, error) {
    err := db.WfDb.Create(&staticAssignment).Error
    return staticAssignment, err
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:05 PM 2021/4/4
 * @Param
 * @return
 **/
func (staticAssignment *WfTransitionStaticAssignment) Find(cond *SqlCond) ([]WfTransitionStaticAssignment, error) {
    staticAssignmentList := make([]WfTransitionStaticAssignment, 0)
    err := cond.Find(db.WfDb, &staticAssignmentList)
    return staticAssignmentList, err
}

/**
 * @Author canweiyao
 * @Description 返回form
 * @Date 7:04 PM 2021/4/4
 * @Param
 * @return
 **/
func (staticAssignment *WfTransitionStaticAssignment) FindBy(cond *SqlCond) (*WfTransitionStaticAssignment, error) {
    err := cond.FindBy(db.WfDb, staticAssignment)
    if err != nil {
        return &WfTransitionStaticAssignment{}, err
    }
    return staticAssignment, nil
}

/**
 * @Author canweiyao
 * @Description 根据id删除guard
 * @Date 11:28 AM 2021/4/11
 * @Param
 * @return
 **/
func (staticAssignment *WfTransitionStaticAssignment) Delete(id int) error {
    return db.WfDb.Where("id = ?", id).Delete(&staticAssignment).Error
}