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
 * @Description  //TODO
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type FieldsController struct {
}

/**
 * @Author canweiyao
 * @Description 创建工作流case
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (fields *FieldsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var field model.WfField
	if err := c.ShouldBind(&field); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}

	formID, err := strconv.Atoi(c.Param("form_id"))
	if err != nil {
		handlers.JSONResponse(c, 404, err.Error(), "err")
	}
	//currentUser := c.GetHeader("wf-usereid")
	field.FormID = formID

	fieldInfo, err := field.Create()

	if err != nil {
		logs.Errorw("create arc error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, fieldInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description //获取单条field
 * @Date 1:49 PM 2021/4/5
 * @Param
 * @return
 **/
func (fields *FieldsController) Get(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	formID := c.Param("form_id")
	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfField := &model.WfField{}
	var err error
	wfField, err = wfField.FindBy(model.NewSqlCnd().Where("id =? and form_id =?", id, formID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfField, "ok")
	return
}

func (fields *FieldsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	formID := c.Param("form_id")

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfField := &model.WfField{}
	var err error
	wfField, err = wfField.FindBy(model.NewSqlCnd().Where("id =? and form_id =?", id, formID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfField, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description 更新field
 * @Date 10:59 PM 2021/4/6
 * @Param
 * @return
 **/
func (fields FieldsController) Put(c *gin.Context) {
	var body struct {
		Name         string `json:"name"`
		FormID       int `json:"form_id"`
		FieldType    int `json:"field_type"`
		Position     int `json:"position"`
		DefaultValue int `json:"default_value"`
	}
	if err := c.ShouldBind(body); err != nil {
		handlers.ParameterValidateResponse(c, 2001, err.Error(), "参数验证错误")
		return
	}

	formID := c.Param("form_id")
	if formID == "" {
		handlers.JSONResponse(c, 500, "form_id is required.", "err")
		return
	}
	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 500, "id is required.", "err")
		return
	}
	wfField := &model.WfField{}
	wfField, _ = wfField.FindBy(model.NewSqlCnd().Where("id =? and form_id = ?", id, formID))
	wfField.Name = body.Name
	wfField.FormID = body.FormID
	wfField.FieldType = body.FieldType
	wfField.Position = body.Position
	err := wfField.Update()
	if err != nil {
		handlers.JSONResponse(c, 500, "更新失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, wfField, "ok")
}
