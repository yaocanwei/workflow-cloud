package wf

import (
	"flowpipe-server/core/model"
	"flowpipe-server/handlers"
	"flowpipe-server/util"

	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  流程变迁
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type TransitionsController struct {
}

/**
 * @Author canweiyao
 * @Description 创建流程变迁
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (transitions *TransitionsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	wfTransition := model.WfTranstion{}
	if err := c.ShouldBind(&wfTransition); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}

	transitionInfo, err := wfTransition.Create()
	if err != nil {
		logs.Errorw("create arc error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, transitionInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description //获取单条表单
 * @Date 1:49 PM 2021/4/5
 * @Param
 * @return
 **/
func (transitions *TransitionsController) Get(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	workflowID := c.Param("workflow_id")
	if workflowID == "" {
		handlers.JSONResponse(c, 2001, "workflow_id is required.", "err")
		return
	}

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfTransition := &model.WfTranstion{}
	var err error
	wfTransition, err = wfTransition.FindBy(model.NewSqlCnd().Where("id", id))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfTransition, "ok")
	return
}

func (transitions *TransitionsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	workflowID := c.Param("workflow_id")
	if workflowID == "" {
		handlers.JSONResponse(c, 2001, "workflow_id is required.", "err")
		return
	}

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfTransition := &model.WfTranstion{}
	var err error
	wfTransition, err = wfTransition.FindBy(model.NewSqlCnd().Where("id =? and workflow_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfTransition, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description transitions 列表
 * @Date 12:00 AM 2021/4/5
 * @Param
 * @return
 **/
func (transitions *TransitionsController) List(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	limit, skip, err := handlers.ParseLimitSkip(c)
	if err != nil {
		logs.Errorw("参数验证失败", err)
		handlers.JSONResponse(c, 2001, err.Error(), "err")
		return
	}

	wfTransition := &model.WfTranstion{}
	wfTransitionList, err := wfTransition.Find(model.NewSqlCnd())
	if err != nil {
		logs.Errorw("list form error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}
	count, err := wfTransition.Count()
	if err != nil {
		logs.Errorw("list transition case count error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}

	pagination := handlers.Pagination{Count: count, Skip: skip, Limit: limit}
	handlers.JSONResponseWithPagination(c, 200, wfTransitionList, "ok", pagination)
	return
}

func (transitions *TransitionsController) Put(c *gin.Context) {
	var body struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		TriggerLimit int    `json:"trigger_limit"`
		TriggerType  int    `json:"trigger_type"`
		SortOrder    int    `json:"sort_order"`
	}
	if err := c.ShouldBind(body); err != nil {
		handlers.ParameterValidateResponse(c, 2001, err.Error(), "参数验证错误")
		return
	}

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 500, "id is required.", "err")
		return
	}
	wfForm := &model.WfForm{}
	wfForm, _ = wfForm.FindBy(model.NewSqlCnd().Where("id =?", id))
	wfForm.Name = body.Name
	wfForm.Description = body.Description
	err := wfForm.Update()
	if err != nil {
		handlers.JSONResponse(c, 500, "更新失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, wfForm, "ok")
}
