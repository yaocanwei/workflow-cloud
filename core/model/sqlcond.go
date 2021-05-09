/**
  @author: cheney
  @date: 2021/4/4
  @note:
 **/
package model

import "github.com/jinzhu/gorm"

type SqlCond struct {
	queryCols []string // 要查询的字段，如果为空，表示查询所有字段
	WhereParams    []ParamPair  // where参数
	JoinParams []ParamPair // 连表查询参数
	PreloadColumns []PreloadColumnPair // 预加载列
	Orders    []OrderByCol // 排序
	Paging    *Paging      // 分页`
}

type ParamPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

type PreloadColumnPair struct {
	Column string // preload的列
	Args []interface{} // 参数
}

// 排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

// 分页请求数据
type Paging struct {
	Page  int   `json:"page"`  // 页码
	Limit int   `json:"limit"` // 每页条数
	Total int64 `json:"total"` // 总数据条数
}

func NewSqlCnd() *SqlCond {
	return &SqlCond{}
}


func (s *SqlCond) Build(db *gorm.DB) *gorm.DB {
	ret := db

	if len(s.queryCols) > 0 {
		ret = ret.Select(s.queryCols)
	}

	// where
	if len(s.WhereParams) > 0 {
		for _, param := range s.WhereParams {
			ret = ret.Where(param.Query, param.Args...)
		}
	}

	// joins
	if len(s.JoinParams) > 0 {
		for _, param := range s.JoinParams {
			ret = ret.Joins(param.Query, param.Args...)
		}
	}

	// preload
	if len(s.PreloadColumns) > 0 {
		for _, column := range s.PreloadColumns {
			ret = ret.Preload(column.Column, column.Args)
		}
	}

	// order
	if len(s.Orders) > 0 {
		for _, order := range s.Orders {
			if order.Asc {
				ret = ret.Order(order.Column + " ASC")
			} else {
				ret = ret.Order(order.Column + " DESC")
			}
		}
	}

	// limit
	if s.Paging != nil && s.Paging.Limit > 0 {
		ret = ret.Limit(s.Paging.Limit)
	}

	// offset
	if s.Paging != nil && s.Paging.Offset() > 0 {
		ret = ret.Offset(s.Paging.Offset())
	}
	return ret
}


func (s *SqlCond) Find(db *gorm.DB, out interface{}) error {
	return s.Build(db).Find(out).Error
}

func (s *SqlCond) FindBy(db *gorm.DB, out interface{}) error {
	if err := s.Limit(1).Build(db).First(out).Error; err != nil {
		return err
	}
	return nil
}

func (s *SqlCond) Joins(query string, args ...interface{}) *SqlCond {
	s.JoinParams = append(s.JoinParams, ParamPair{query, args})
	return s
}

func (s *SqlCond) Order(column string, asc bool) *SqlCond {
	s.Orders = append(s.Orders, OrderByCol{column, asc})
	return s
}

func (s *SqlCond) Where(query string, args ...interface{}) *SqlCond {
	s.WhereParams = append(s.WhereParams, ParamPair{query, args})
	return s
}

func (s *SqlCond) Preload(column string, conditions ...interface{}) *SqlCond {
	s.PreloadColumns = append(s.PreloadColumns, PreloadColumnPair{column, conditions})
	return s
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

func (s *SqlCond) Limit(limit int) *SqlCond {
	s.Page(1, limit)
	return s
}

func (s *SqlCond) Page(page, limit int) *SqlCond {
	if s.Paging == nil {
		s.Paging = &Paging{Page: page, Limit: limit}
	} else {
		s.Paging.Page = page
		s.Paging.Limit = limit
	}
	return s
}

func (s *SqlCond) Count(db *gorm.DB, model interface{}) (int, error) {
	var err error
	ret := db.Model(model)

	// where
	if len(s.WhereParams) > 0 {
		for _, query := range s.WhereParams {
			ret = ret.Where(query.Query, query.Args...)
		}
	}

	var count int
	err = ret.Count(&count).Error
	return count, err
}