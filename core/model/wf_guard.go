package model

import (
	db "flowpipe-server/core/db"
	"time"
)

// WfGuard 弧守卫函数, 用于消解冲突, 从而选择一条唯一执行路径
type WfGuard struct {
	Model
	ArcID         int       `json:"arc_id"`
	WorkflowID    int       `json:"workflow_id"`
	FieldableType string    `json:"fieldable_type"`
	FieldableID   int    `json:"fieldable_id"`
	Op            string    `json:"op"`
	Value         string    `json:"value"`
	Exp           string    `json:"exp"` //表达式
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at" json:"updated_at"`
}

/**
TODO: validate guard workflowid equals arcs workflowid
 */

/**
 * @Author canweiyao
 * @Description 创建弧守卫函数
 * @Date 8:37 AM 2021/4/10
 * @Param
 * @return
 **/
func (guard *WfGuard) Create() (*WfGuard, error) {
	err := db.WfDb.Create(&guard).Error
	return guard, err
}

/**
 * @Author canweiyao
 * @Description Updates 更新弧守卫函数
 * @Date 7:19 AM 2021/4/6
 * @Param
 * @return
 **/
func (guard *WfGuard) Update() (err error) {
	err = db.WfDb.Save(guard).Error
	return
}


/**
 * @Author canweiyao
 * @Description 返回弧守卫函数
 * @Date 7:04 PM 2021/4/4
 * @Param
 * @return
 **/
func (guard *WfGuard) FindBy(cond *SqlCond) (*WfGuard, error) {
	err := cond.FindBy(db.WfDb, guard)
	if err != nil {
		return &WfGuard{}, err
	}
	return guard, nil
}

// Delete 根据id删除guard
func (guard *WfGuard) Delete(id int) error {
	return db.WfDb.Where("id = ?", id).Delete(&guard).Error
}