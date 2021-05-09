package wf

import (
	"flowpipe-server/core/model"
	"flowpipe-server/handlers"
	"flowpipe-server/util"

	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  创建流程
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type WorkflowsController struct {
}

type Body struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

/**
 * @Author canweiyao
 * @Description 创建工作流case
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (workflows *WorkflowsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	currentUser := c.GetHeader("wf-usereid")
	var body *Body
	if err := c.ShouldBind(&body); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}

	workflow := model.WfWorkflow{}
	workflow.Name = body.Name
	workflow.Description = body.Description
	workflow.CreatedBy = currentUser
	workflowInfo, err := workflow.Create()
	if err != nil {
		logs.Errorw("create workflow error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, workflowInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description //获取单条流程
 * @Date 1:49 PM 2021/4/5
 * @Param
 * @return
 **/
func (workflows *WorkflowsController) Get(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfWorkflow := &model.WfWorkflow{}
	var err error
	wfWorkflow, err = wfWorkflow.FindBy(model.NewSqlCnd().Where("id =?", id))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfWorkflow, "ok")
	return
}

func (workflows *WorkflowsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfWorkflow := &model.WfWorkflow{}
	var err error
	wfWorkflow, err = wfWorkflow.FindBy(model.NewSqlCnd().Where("id =?", id))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfWorkflow, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description 流程列表
 * @Date 12:00 AM 2021/4/5
 * @Param
 * @return
 **/
func (workflows *WorkflowsController) List(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	limit, skip, err := handlers.ParseLimitSkip(c)
	if err != nil {
		logs.Errorw("参数验证失败", err)
		handlers.JSONResponse(c, 2001, err.Error(), "err")
		return
	}

	wfWorkflow := &model.WfWorkflow{}
	wfWorkflowList, err := wfWorkflow.Find(model.NewSqlCnd().Order("id", false))
	if err != nil {
		logs.Errorw("list workflow error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}
	count, err := wfWorkflow.Count(model.NewSqlCnd())
	if err != nil {
		logs.Errorw("list workflow case count error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}

	pagination := handlers.Pagination{Count: count, Skip: skip, Limit: limit}
	handlers.JSONResponseWithPagination(c, 200, wfWorkflowList, "ok", pagination)
	return
}

func (workflows *WorkflowsController) Put(c *gin.Context) {
	var body *Body
	if err := c.ShouldBind(&body); err != nil {
		handlers.ParameterValidateResponse(c, 2001, err.Error(), "参数验证错误")
		return
	}

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 500, "id is required.", "err")
		return
	}
	wfWorkflow := &model.WfWorkflow{}
	wfWorkflow, _ = wfWorkflow.FindBy(model.NewSqlCnd().Where("id =?", id))
	wfWorkflow.Name = body.Name
	wfWorkflow.Description = body.Description
	err := wfWorkflow.Update()
	if err != nil {
		handlers.JSONResponse(c, 500, "更新失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, wfWorkflow, "ok")
}
