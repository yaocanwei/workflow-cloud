package wf

import (
	"flowpipe-server/core/model"
	"flowpipe-server/handlers"
	"flowpipe-server/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArcsController struct {
}

/**
 * @Author canweiyao
 * @Description  创建有向弧
 * @Date 5:39 PM 2021/4/5
 * @Param
 * @return
 **/
func (arcs ArcsController) Create(c *gin.Context) {
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

	currentUser := c.GetHeader("wf-usereid")
	arc.WorkflowID = workflowID
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
 * @Description  删除arc
 * @Date 10:59 PM 2021/4/6
 * @Param
 * @return
 **/
func (arcs ArcsController) delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	workflowID := c.Param("workflow_id")

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfArc := &model.WfArc{}
	var err error
	wfArc, err = wfArc.FindBy(model.NewSqlCnd().Where("id =? and workflow_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfArc, "ok")
	return

}

/**
 * @Author canweiyao
 * @Description 更新arc
 * @Date 10:59 PM 2021/4/6
 * @Param
 * @return
 **/
func (arcs ArcsController) Put(c *gin.Context) {
	var body struct {
		direction     int `json:"direction"`
		transitionID int `json:"transition_id"`
		placeID int    `json:"place_id"`
	}
	if err := c.ShouldBind(body); err != nil {
		handlers.ParameterValidateResponse(c, 2001, err.Error(), "参数验证错误")
		return
	}
	//currentUser := c.GetHeader("wf-usereid")
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
	arc := &model.WfArc{}
	arc, _ = arc.FindBy(model.NewSqlCnd().Where("id =? and workflow_id = ?", id, workflowID))
	arc.Direction = body.direction
	arc.TransitionID = body.transitionID
	arc.PlaceID = body.placeID
	err := arc.Update()
	if err != nil {
		handlers.JSONResponse(c, 500, "更新失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, arc, "ok")
}

/**
 * @Author canweiyao
 * @Description  获取arc列表
 * @Date 10:59 PM 2021/4/6
 * @Param
 * @return
 **/
func (arcs ArcsController)  List(c *gin.Context) {
	var err error
	workflowID := c.Param("workflow_id")
	if workflowID == "" {
		handlers.JSONResponse(c, 500, "workflow_id is required.", "err")
		return
	}
	arc := &model.WfArc{}
	arcList, err := arc.Find(model.NewSqlCnd().Where("workflow_id =?", workflowID))
	if err != nil {
		handlers.JSONResponse(c, 500, "查找失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, arcList, "ok")
}

/**
 * @Author canweiyao
 * @Description 获取单条arc
 * @Date 10:58 PM 2021/4/6
 * @Param
 * @return
 **/
func (arcs ArcsController) Get(c *gin.Context) {
	var err error
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
	arc := &model.WfArc{}
	arc, err = arc.FindBy(model.NewSqlCnd().Where("id =? and workflow_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 500, "查找失败.", "err")
		return
	}
	handlers.JSONResponse(c, 200, arc, "ok")
}
