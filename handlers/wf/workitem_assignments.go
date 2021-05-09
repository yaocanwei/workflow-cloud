package wf

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/core/model/command"
	"flowpipe-server/handlers"
	"flowpipe-server/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  创建place
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type WorkitemAssignmentsController struct {
}

/**
 * @Author canweiyao
 * @Description 创建指派任务
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (workitemAssignments *WorkitemAssignmentsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var body struct {
		PartyID int `json:"party_id"`
	}
	var err error
	staticAssgnment := model.WfWorkitemAssignment{}
	if err := c.ShouldBind(&body); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}
	workitemID := c.Param("workitem_id")
	if workitemID == "" {
		handlers.JSONResponse(c, 500, "workitem_id is required.", "err")
		return
	}

	staticAssgnment.PartyID = body.PartyID
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
func (workitemAssignments *WorkitemAssignmentsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()
	var err error
	workitemID := c.Param("workitem_id")
	if workitemID == "" {
		handlers.JSONResponse(c, 2001, "workitem_id is required.", "err")
		return
	}

	partyID := c.Query("party_id")
	intPartyID, err := strconv.Atoi(partyID)
	intWorkitemID, err := strconv.Atoi(workitemID)
	task := command.NewWorkitemAssignmentTask(intWorkitemID, intPartyID, true)
	err = command.HandleWorkitemAssignmentTask(context.Background(), task)
	if err != nil {
		logs.Errorw("delete workitem error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}
	return
}
