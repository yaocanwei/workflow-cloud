package model

import(
    "flowpipe-server/core/db"
    "time"
)

type WfGroup struct {
    Model
    Name int `json: "name"`
    CreatedAt time.Time `gorm:"created_at" json:"created_at"`
    UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
    // TODO index
}

/**
 * @Author canweiyao
 * @Description  查找符合条件的租
 * @Date 11:47 PM 2021/4/11
 * @Param
 * @return
 **/
func (group *WfGroup) Find(cond *SqlCond) ([]WfUser, error) {
    wfUserList := make([]WfUser, 0)
    err := cond.Find(db.WfDb, &wfUserList)
    return wfUserList, err
}

/**
 * @Author canweiyao
 * @Description 根据条件查找单个用户
 * @Date 11:47 PM 2021/4/11
 * @Param
 * @return
 **/
func (group *WfGroup) FindBy(cond *SqlCond) (*WfGroup, error) {
    err := cond.FindBy(db.WfDb, group)
    if err != nil {
        return &WfGroup{}, err
    }
    return group, nil
}
