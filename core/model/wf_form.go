package model

import (
	"flowpipe-server/core/db"
	"time"
	// "github.com/jinzhu/gorm"
)

// WfForm 流程表单
type WfForm struct {
	Model
	Name        string    `json: "name"`
	Description string    `json: "description" gorm:"type:text"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at" json:"updated_at"`
}

// Create 创建
func (wfForm *WfForm) Create() (*WfForm, error) {
	err := db.WfDb.Create(&wfForm).Error
	return wfForm, err
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:05 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfForm *WfForm) Find(cond *SqlCond) ([]WfForm, error) {
	wfFormList := make([]WfForm, 0)
	err := cond.Find(db.WfDb, &wfFormList)
	return wfFormList, err
}

/**
 * @Author canweiyao
 * @Description 返回form
 * @Date 7:04 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfForm *WfForm) FindBy(cond *SqlCond) (*WfForm, error) {
	err := cond.FindBy(db.WfDb, wfForm)
	if err != nil {
		return &WfForm{}, err
	}
	return wfForm, nil
}

// Count 统计form
func (wfForm *WfForm) Count() (count int, err error) {
	if err = db.WfDb.Model(&wfForm).Unscoped().Count(&count).Error; err != nil {
		return
	}
	return
}

/**
 * @Author canweiyao
 * @Description Updates 更新form
 * @Date 7:19 AM 2021/4/6
 * @Param
 * @return
 **/
func (wfForm *WfForm) Update() (err error) {
	err = db.WfDb.Save(wfForm).Error
	return
}


// Delete 根据id删除guard
func (wfForm *WfForm) Delete(id int) error {
	return db.WfDb.Where("id = ?", id).Delete(&wfForm).Error
}