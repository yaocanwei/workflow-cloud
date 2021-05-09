//Package model 提供通用部分
package model

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key" json:"id,omitempty"`
}
