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
type StaticAssignmentsController struct {
}

/**
 * @Author canweiyao
 * @Description 创建指派任务
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (staticAssignments *StaticAssignmentsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var err error
	staticAssgnment := model.WfTransitionStaticAssignment{}
	if err := c.ShouldBind(&staticAssgnment); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}
	transitionID := c.Param("transition_id")
	if transitionID == "" {
		handlers.JSONResponse(c, 500, "transition_id is required.", "err")
		return
	}

	staticAssgnmentInfo, err := staticAssgnment.Create()
	if err != nil {
		logs.Errorw("create static assgnment error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	handlers.JSONResponse(c, 200, staticAssgnmentInfo, "ok")
}

/**
 * @Author canweiyao
 * @Description 删除place
 * @Date 8:51 AM 2021/4/10
 * @Param
 * @return
 **/
func (staticAssignments *StaticAssignmentsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var err error
	transitionID := c.Param("transition_id")
	if transitionID == "" {
		handlers.JSONResponse(c, 2001, "transition_id is required.", "err")
		return
	}

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	staticAssgnment := &model.WfTransitionStaticAssignment{}
	staticAssgnment, err = staticAssgnment.FindBy(model.NewSqlCnd().Where("id =? and transition_id =?", id, transitionID))
	if err != nil {
		handlers.JSONResponse(c, 404, "places not found", "err")
		return
	}
	err = staticAssgnment.Delete(staticAssgnment.ID)
	if err != nil {
		handlers.JSONResponse(c, 500, "internal error", "err")
	}
	handlers.JSONResponse(c, 200, staticAssgnment, "ok")
	return
}
