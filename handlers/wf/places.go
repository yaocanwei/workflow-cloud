package wf

import (
	"flowpipe-server/core/model"
	"flowpipe-server/handlers"
	"flowpipe-server/util"
	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  创建place
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type PlacesController struct {
}

/**
 * @Author canweiyao
 * @Description 创建表单
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (places *PlacesController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var err error
	place := model.WfPlace{}
	if err := c.ShouldBind(&place); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}
	workflowID := c.Param("workflow_id")
	if workflowID == "" {
		handlers.JSONResponse(c, 500, "arc_id is required.", "err")
		return
	}

	placeInfo, err := place.Create()
	if err != nil {
		logs.Errorw("create place error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, placeInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description 删除place
 * @Date 8:51 AM 2021/4/10
 * @Param
 * @return
 **/
func (places *PlacesController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var err error
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
	wfPlace := &model.WfPlace{}
	wfPlace, err = wfPlace.FindBy(model.NewSqlCnd().Where("id =? and arc_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 404, "places not found", "err")
		return
	}
	err = wfPlace.Delete(wfPlace.ID)
	if err != nil {
		handlers.JSONResponse(c, 500, "internal error", "err")
	}
	handlers.JSONResponse(c, 200, wfPlace, "ok")
	return
}

/**
 * @Author canweiyao
 * @Description places 列表
 * @Date 12:00 AM 2021/4/5
 * @Param
 * @return
 **/
func (places *PlacesController) List(c *gin.Context) {
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 8:51 AM 2021/4/10
 * @Param
 * @return
 **/
func (places *PlacesController) Put(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PlaceType   string `json:"place_type"`
		SortOrder   string `json:"sort_order"`
	}
	if err := c.ShouldBind(body); err != nil {
		handlers.ParameterValidateResponse(c, 2001, err.Error(), "参数验证错误")
		return
	}

	workflowID := c.Param("workflow_id")
	if workflowID == "" {
		handlers.JSONResponse(c, 500, "workflow_id is required.", "err")
		return
	}
	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 500, "id is required.", "err")
		return
	}
	wfPlace := &model.WfPlace{}
	wfPlace, _ = wfPlace.FindBy(model.NewSqlCnd().Where("id =? and workflow_id = ?", id, workflowID))
	err := wfPlace.Update()
	if err != nil {
		handlers.JSONResponse(c, 500, "更新失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, wfPlace, "ok")
}
