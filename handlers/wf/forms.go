package wf

import (
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
type FormsController struct {
}

/**
 * @Author canweiyao
 * @Description 创建表单
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (forms *FormsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	form := model.WfForm{}
	if err := c.ShouldBind(&form); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}

	formInfo, err := form.Create()
	if err != nil {
		logs.Errorw("create arc error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, formInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description //获取单条表单
 * @Date 1:49 PM 2021/4/5
 * @Param
 * @return
 **/
func (forms *FormsController) Get(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfForm := &model.WfForm{}
	var err error
	wfForm, err = wfForm.FindBy(model.NewSqlCnd().Where("id", id))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfForm, "ok")
	return
}

func (forms *FormsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfForm := &model.WfForm{}
	var err error
	wfForm, err = wfForm.FindBy(model.NewSqlCnd().Where("id =?", id))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfForm, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description 流程case 列表
 * @Date 12:00 AM 2021/4/5
 * @Param
 * @return
 **/
func (forms *FormsController) List(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	limit, skip, err := handlers.ParseLimitSkip(c)
	if err != nil {
		logs.Errorw("参数验证失败", err)
		handlers.JSONResponse(c, 2001, err.Error(), "err")
		return
	}

	wfForm := &model.WfForm{}
	wfForms, err := wfForm.Find(model.NewSqlCnd())
	if err != nil {
		logs.Errorw("list form error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}
	count, err := wfForm.Count()
	if err != nil {
		logs.Errorw("list workflow case count error", err)
		handlers.JSONResponse(c, 5202, err.Error(), "err")
		return
	}

	pagination := handlers.Pagination{Count: count, Skip: skip, Limit: limit}
	handlers.JSONResponseWithPagination(c, 200, wfForms, "ok", pagination)
	return
}

func (forms *FormsController) Put(c *gin.Context) {
	var body struct {
		Name         string `json:"name"`
		Description string `json:"description"`
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