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

type WorkitemsController struct {
}

/**
 * @Author canweiyao
 * @Description  创建任务
 * @Date 9:34 PM 2021/4/11
 * @Param
 * @return
 **/
func (workitems *WorkitemsController) Create(c *gin.Context) {
}

/**
 * @Author canweiyao
 * @Description  获取任务详情
 * @Date 9:37 PM 2021/4/11
 * @Param
 * @return
 **/
func (workitems *WorkitemsController) Get(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	wfWorkitem := &model.WfWorkitem{}
	var err error
	wfWorkitem, err = wfWorkitem.FindBy(model.NewSqlCnd().Where("id =?", id))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfWorkitem, "ok")
	return
}

func (workitems *WorkitemsController) List(c *gin.Context) {
	currentUser := c.GetHeader("wf-usereid")
	userID, _ := strconv.Atoi(currentUser)
	wfWorkitem := &model.WfWorkitem{}
	wfWrokitems, err := wfWorkitem.Todo(userID, c.Query("state"))
	if err != nil {
		handlers.JSONResponse(c, 404, "not todo workitems return", "ok")
		return
	}
	handlers.JSONResponse(c, 200, wfWrokitems, "ok")
	return
}

func (workitems *WorkitemsController) Start(c *gin.Context) {
	var err error
	currentUser, _ := strconv.Atoi(c.GetHeader("wf-usereid"))
	workitemID, _ := strconv.Atoi(c.Param("id"))
	task := command.NewWorkitemStartTask(workitemID, currentUser)
	err = command.HandleWorkitemStartTask(context.Background(), task)
	if err != nil {
		handlers.JSONResponse(c, 500, "cannot start workitem", "ok")
		return
	}

	handlers.JSONResponse(c, 200, "start workitem successfully", "ok")
}

func (workitems *WorkitemsController) Finish(c *gin.Context) {

}
