package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfUser 用户模型
type WfUser struct {
	Model
	Account      string    `json:"account"`
	Name         string    `json:"name"`
	Party        WfParty   `json:"party" gorm:"FOREIGNKEY:UUID;ASSOCIATION_FOREIGNKEY:PartyID"`
	GroupID      int       `json:"group_id"`
	Group        WfGroup   `json:"gruop" gorm:"FOREIGNKEY:UUID;ASSOCIATION_FOREIGNKEY:GroupID"`
	DepartmentID int       `json:"department_id"`
	PositionID   int       `json:"position_id"`
	RoleID       int       `json:"role_id"`
	TagID        int       `json:"tag_id"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `json:"column:updated_at"`
}

/**
 * @Author canweiyao
 * @Description  查找符合条件的用户
 * @Date 6:46 PM 2021/4/4
 * @Param
 * @return
 **/
func (user *WfUser) Find(cond *SqlCond) ([]WfUser, error) {
	wfUserList := make([]WfUser, 0)
	err := cond.Find(db.WfDb, &wfUserList)
	return wfUserList, err
}

/**
 * @Author canweiyao
 * @Description 根据条件查找单个用户
 * @Date 6:50 PM 2021/4/4
 * @Param
 * @return
 **/
func (user *WfUser) FindBy(cond *SqlCond) (*WfUser, error) {
	err := cond.Preload("Group").FindBy(db.WfDb, user)
	if err != nil {
		return &WfUser{}, err
	}
	return user, nil
}
