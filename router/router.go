package router

import (
	"flowpipe-server/handlers/wf"
	"flowpipe-server/router/middleware/header"

	"github.com/gin-gonic/gin"
)

// SetupRouter 启动路由
func SetupRouter(mids ...gin.HandlerFunc) (e *gin.Engine, err error) {
	e = gin.New()

	e.Use(gin.Recovery())
	e.Use(mids...)
	e.Use(
		header.NoCache,
		header.Secure,
		header.Options,
		header.Validator,
	)

	root := e.Group("/api")

	workflowGroup := root.Group(("workflows"))
	{
		workflowGroup.POST("", wf.CreateWorkflow)
	}

	caseGroup := root.Group("workflows/:workflow_id/cases")
	{
		casesController := new(wf.CasesController)
		caseGroup.POST("", casesController.Create)
		caseGroup.GET("", casesController.List)
		caseGroup.GET("/:id", casesController.Get)
	}

	placeGroup := root.Group("workflows/:workflow_id/places")
	{
		placeGroup.POST("", wf.CreatePlace)
		placeGroup.DELETE("/:id", wf.DestroyPlace)
		placeGroup.GET("/:id", wf.DestroyPlace)
		placeGroup.PUT("/:id", wf.UpdatePlace)
	}

	return
}
