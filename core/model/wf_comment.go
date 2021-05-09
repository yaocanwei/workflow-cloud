package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfComment 节点评论
type WfComment struct {
	Model
	WorkitemID int       `json:"workitem_id"`
	UserID     int       `json:"user_id"`
	Body       string    `json:"body" gorm:"type:text"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at" json:"updated_at"`
	// TODO index
}

// Create 创建评论
func (comment *WfComment) Create() (*WfComment, error) {
	err := db.WfDb.Create(&comment).Error
	return comment, err
}

/**
 * @Author canweiyao
 * @Description  查找符合条件的
 * @Date 6:46 PM 2021/4/4
 * @Param
 * @return
 **/
func (comment *WfComment) Find(cond *SqlCond) ([]WfComment, error) {
	wfCommentList := make([]WfComment, 0)
	err := cond.Find(db.WfDb, &wfCommentList)
	return wfCommentList, err
}

/**
 * @Author canweiyao
 * @Description 根据条件查找comment
 * @Date 6:50 PM 2021/4/4
 * @Param
 * @return
 **/
func (comment *WfComment) FindBy(cond *SqlCond) (*WfComment, error) {
	err := cond.FindBy(db.WfDb, comment)
	if err != nil {
		return &WfComment{}, err
	}
	return comment, nil
}

