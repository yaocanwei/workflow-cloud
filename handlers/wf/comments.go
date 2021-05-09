package wf

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/core/model/command"
	"flowpipe-server/handlers"
	"flowpipe-server/util"
	"github.com/gin-gonic/gin"
)

/**
 * @Author canweiyao
 * @Description  创建流程实例
 * @Date 10:16 PM 2021/4/4
 * @Param
 * @return
 **/
type CommentsController struct {
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 12:54 AM 2021/4/5
 * @Param
 * @return
 **/
func (comments *CommentsController) Create(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	var comment model.WfComment
	if err := c.ShouldBind(&comment); err != nil {
		handlers.ParameterValidateResponse(c, 401, err.Error(), "参数错误")
		return
	}

	currentUser := c.GetHeader("wf-usereid")
	user := &model.WfUser{}
	var err error
	user, err = user.FindBy(model.NewSqlCnd().Where("account =?", currentUser))
	task := command.NewCommentAddTask(comment.WorkitemID, user.ID, comment.Body)
	err = command.HandleCommentAddTask(context.Background(), task)
	if err != nil {
		logs.Errorw("create arc error", err)
		handlers.JSONResponse(c, 500, err.Error(), "err")
		return
	}

	workitem := &model.WfWorkitem{}
	workitem, err = workitem.FindBy(model.NewSqlCnd().Where("id =?", comment.ID))
	handlers.JSONResponse(c, 200, workitem, "ok")
}

func (comments *CommentsController) Delete(c *gin.Context) {
	logs := util.Logs()
	defer logs.Sync()

	workflowID := c.Param("workitem_id")

	id := c.Param("id")
	if id == "" {
		handlers.JSONResponse(c, 2001, "id is required.", "err")
		return
	}
	comment := &model.WfComment{}
	var err error
	comment, err = comment.FindBy(model.NewSqlCnd().Where("id =? and workitem_id =?", id, workflowID))
	if err != nil {
		handlers.JSONResponse(c, 404, "", "ok")
		return
	}
	handlers.JSONResponse(c, 200, comment, "ok")
	return
}

