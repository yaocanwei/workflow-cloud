package wf

import (
	"errors"
	"flowpipe-server/core/model"
	"flowpipe-server/handlers"
	"flowpipe-server/util"
	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  创建表单实例
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type GuardsController struct {
}

/**
 * @Author canweiyao
 * @Description 创建表单
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (guards *GuardsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var err error
	guard := model.WfGuard{}
	if err := c.ShouldBind(&guard); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}
	ardID := c.Param("arc_id")
	if ardID == "" {
		handlers.JSONResponse(c, 500, "arc_id is required.", "err")
		return
	}
	arc := &model.WfArc{}
	arc, err = arc.FindBy(model.NewSqlCnd().Where("id =?", ardID))
	if arc.Direction == 0 {
		handlers.JSONResponse(c, 502, errors.New("only out direction arc can set guard"), "err")
	}
	guardInfo, err := guard.Create()
	if err != nil {
		logs.Errorw("create arc error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, guardInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 8:51 AM 2021/4/10
 * @Param
 * @return
 **/
func (guards *GuardsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var err error
	arcID := c.Param("arc_id")
	if arcID == "" {
		handlers.JSONResponse(c, 2001, "arc_id is required.", "err")
		return
	}

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfGuard := &model.WfGuard{}
	wfGuard, err = wfGuard.FindBy(model.NewSqlCnd().Where("id =? and arc_id =?", id, arcID))
	if err != nil {
		handlers.JSONResponse(c, 404, "guard not found", "err")
		return
	}
	err = wfGuard.Delete(wfGuard.ID)
	if err != nil {
		handlers.JSONResponse(c, 500, "internal error", "err")
	}
	handlers.JSONResponse(c, 200, wfGuard, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description 流程case 列表
 * @Date 12:00 AM 2021/4/5
 * @Param
 * @return
 **/
func (guards *GuardsController) List(c *gin.Context) {
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 8:51 AM 2021/4/10
 * @Param
 * @return
 **/
func (guards *GuardsController) Put(c *gin.Context) {
	var body struct {
		FieldableType string `json:"fieldable_type"`
		FieldableID   int    `json:"fieldable_id"`
		Op            string `json:"op"`
		Value         string `json:"value"`
		Exp           string `json:"exp"`
	}
	if err := c.ShouldBind(body); err != nil {
		handlers.ParameterValidateResponse(c, 2001, err.Error(), "参数验证错误")
		return
	}

	arcID := c.Param("arc_id")
	if arcID == "" {
		handlers.JSONResponse(c, 500, "arc_id is required.", "err")
		return
	}
	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 500, "id is required.", "err")
		return
	}
	wfGuard := &model.WfGuard{}
	wfGuard, _ = wfGuard.FindBy(model.NewSqlCnd().Where("id =? and arc_id = ?", id, arcID))
	err := wfGuard.Update()
	if err != nil {
		handlers.JSONResponse(c, 500, "更新失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, wfGuard, "ok")
}
