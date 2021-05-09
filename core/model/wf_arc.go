package model

import (
	db "flowpipe-server/core/db"
	"time"
	// "github.com/jinzhu/gorm"
)

// WfArc 是 Transition 到 Place 的连线，ArcGuard 决定连线是否可走，比如 请假时间 > 10 天 则放行之类的逻辑
// 是库所和变迁之间的有向弧
// 有向弧是有方向的
// 两个库所或变迁之间不允许有弧
// 库所可以拥有任意数量的令牌
type WfArc struct {
	Model
	// 流程定义ID
	WorkflowID   int `json:"workflow_id" binding:"required"`
	TransitionID int `json:"transition_id" binding:"required"`
	PlaceID      int `json:"place_id" binding:"required"`
	// default: 0; 0-in, 1-out
	Direction   int       `json:"direction" binding:"required"`
	CreatedAt   time.Time `gorm:"column:create_at" json:"create_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Guards      []WfGuard `gorm:"foreignkey:ArcID;association_foreignkey:ID"`
	GuardsCount int       `gorm:"column:guards_count" json:"guards_count"`
}

// Create 创建有向弧
func (arc *WfArc) Create() (*WfArc, error) {
	err := db.WfDb.Create(arc).Error
	return arc, err
}

/**
 * @Author canweiyao
 * @Description  查找符合条件的记录
 * @Date 6:46 PM 2021/4/4
 * @Param
 * @return
 **/
func (arc *WfArc) Find(cond *SqlCond) ([]WfArc, error) {
	wfArcList := make([]WfArc, 0)
	err := cond.Find(db.WfDb, &wfArcList)
	return wfArcList, err
}

/**
 * @Author canweiyao
 * @Description 根据条件查找单条arc记录
 * @Date 6:50 PM 2021/4/4
 * @Param
 * @return
 **/
func (arc *WfArc) FindBy(cond *SqlCond) (*WfArc, error) {
	err := cond.FindBy(db.WfDb, arc)
	if err != nil {
		return &WfArc{}, err
	}
	return arc, nil
}

/**
 * @Author canweiyao
 * @Description Updates 更新transition
 * @Date 7:19 AM 2021/4/6
 * @Param
 * @return
 **/
func (arc *WfArc) Update() (err error) {
	err = db.WfDb.Save(arc).Error
	return
}

func (arc *WfArc) Updates(id int, columns map[string]interface{}) error {
	return nil
}

func (arc *WfArc) UpdateColumn(id int, name string, value interface{}) (err error) {
	return db.WfDb.Model(&WfArc{}).Where("id =?", id).UpdateColumn(name, value).Error
}

func (arc *WfArc) Delete(id int) error {
	return db.WfDb.Delete(&WfArc{}, "id = ?", id).Error
}

func (arc *WfArc) Get(id int) (*WfArc, error) {
	if err := db.WfDb.First(arc, "id = ?", id).Error; err != nil {
		return &WfArc{}, err
	}
	return arc, nil
}

func (arc *WfArc) Count(cond *SqlCond) (int, error) {
	return cond.Count(db.WfDb, &WfArc{})
}
