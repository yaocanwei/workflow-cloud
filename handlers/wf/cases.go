package wf

import (
	"flowpipe-server/core/model"
	"flowpipe-server/handlers"
	"flowpipe-server/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  创建流程实例
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type CasesController struct {
}

/**
 * @Author canweiyao
 * @Description 创建工作流case
 * @Date 12:54 AM 2021/4/5
 * @Param 
 * @return 
 **/
func (cases *CasesController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var arc model.WfArc
	if err := c.ShouldBind(&arc); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}

	workflowID, err := strconv.Atoi(c.Param("workflow_id"))
	if err != nil {
		handlers.JSONResponse(c, 404, err.Error(), "err")
	}

	arc.WorkflowID = workflowID
	currentUser := c.GetHeader("wf-usereid")
	arc.CreatedBy = currentUser
	arcInfo, err := arc.Create()
	if err != nil {
		logs.Errorw("create arc error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, arcInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description //获取单条case
 * @Date 1:49 PM 2021/4/5
 * @Param
 * @return
 **/
func (cases *CasesController) Get(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	workflowID := c.Param("workflow_id")
	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfCase := &model.WfCase{}
	var err error
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =? and workflow_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfCase, "ok")
	return
}

func (cases *CasesController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	workflowID := c.Param("workflow_id")

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfCase := &model.WfCase{}
	var err error
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =? and workflow_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfCase, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description 流程case 列表
 * @Date 12:00 AM 2021/4/5
 * @Param
 * @return
 **/
func (cases *CasesController) List(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	limit, skip, err := handlers.ParseLimitSkip(c)
	if err != nil {
		logs.Errorw("参数验证失败", err)
		handlers.JSONResponse(c, 2001, err.Error(), "err")
		return
	}

	workflowID := c.Param("workflow_id")

	wfCase := &model.WfCase{}
	wfCases, err := wfCase.Find(model.NewSqlCnd().Where("workflow_id =?", workflowID))
	if err != nil {
		logs.Errorw("list workflow case error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}
	count, err := wfCase.Count()
	if err != nil {
		logs.Errorw("list workflow case count error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}

	pagination := handlers.Pagination{Count: count, Skip: skip, Limit: limit}
	handlers.JSONResponseWithPagination(c, 200, wfCases, "ok", pagination)
	return
}
