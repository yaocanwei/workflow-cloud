package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfField 字段控制
type WfField struct {
	Model
	Name          string    `json:"name"`
	FormID        int       `json:"form_id"`
	Position      int       `json:"position" gorm:"column:position default: 0"` // 排序
	FieldType     int       `json:"field_type"`
	FieldTypeName string    `json:"field_type_name"`
	DefaultValue  string    `json:"default_value"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at" json:"updated_at"`
	// TODO index
}

// Create 创建field
func (field *WfField) Create() (*WfField, error) {
	err := db.WfDb.Create(field).Error
	return field, err
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:05 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfField *WfField) Find(cond *SqlCond) ([]WfField, error) {
	wfFieldList := make([]WfField, 0)
	err := cond.Find(db.WfDb, &wfFieldList)
	return wfFieldList, err
}

/**
 * @Author canweiyao
 * @Description 返回field
 * @Date 7:04 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfFeild *WfField) FindBy(cond *SqlCond) (*WfField, error) {
	err := cond.FindBy(db.WfDb, wfFeild)
	if err != nil {
		return &WfField{}, err
	}
	return wfFeild, nil
}

/**
 * @Author canweiyao
 * @Description Updates 更新transition
 * @Date 7:19 AM 2021/4/6
 * @Param
 * @return
 **/
func (wfFeild *WfField) Update() (err error) {
	err = db.WfDb.Save(wfFeild).Error
	return
}