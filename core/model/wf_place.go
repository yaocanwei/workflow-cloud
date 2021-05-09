package model

import (
	"flowpipe-server/core/db"
	"time"
	// "github.com/jinzhu/gorm"
)

// WfPlace 库所 圆形节点
type WfPlace struct {
	Model
	WorkflowID  int       `json:"workflow_id" gorm:"index"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" gorm:"type:text"`
	SortOrder   int       `json:"sort_order"`
	PlaceType   int       `json:"palce_type"` // 0: star; 1: normal; 2: end
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	Arcs        []WfArc   `gorm:"foreignkey:PlaceID;association_foreignkey:ID"`
}

// PlaceType 库所类型
type PlaceType int

// Start 开始
// Normal 正常
// End 结束
const (
	Start  PlaceType = 0
	Normal PlaceType = 1
	End    PlaceType = 2
)

// Create 创建库所
func (place *WfPlace) Create() (*WfPlace, error) {
	err := db.WfDb.Create(&place).Error
	return place, err
}

//
func (place PlaceType) placeType() string {
	switch place {
	case Start:
		return "start"
	case Normal:
		return "normal"
	case End:
		return "end"
	default:
		return "unknown"
	}
}

// Delete 删除库所
func (place *WfPlace) Delete(ID int) error {
	err := db.WfDb.Where("id = ?", ID).Delete(&place).Error
	return err
}

// Updates 更新库所
func (place *WfPlace) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(&place).Error
}

/**
 * @Author canweiyao
 * @Description Updates 更新库所
 * @Date 7:19 AM 2021/4/6
 * @Param
 * @return
 **/
func (place *WfPlace) Update() (err error) {
	err = db.WfDb.Save(place).Error
	return
}

/**
 * @Author canweiyao
 * @Description  根据条件查找符合条件的库所
 * @Date 7:09 PM 2021/4/4
 * @Param
 * @return
 **/
func (place *WfPlace) Find(cond *SqlCond) ([]WfPlace, error) {
	wfPlaceList := make([]WfPlace, 0)
	err := cond.Find(db.WfDb, &wfPlaceList)
	return wfPlaceList, err
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:09 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfPlace *WfPlace) FindBy(cond *SqlCond) (*WfPlace, error) {
	err := cond.FindBy(db.WfDb, wfPlace)
	if err != nil {
		return &WfPlace{}, err
	}
	return wfPlace, nil
}
