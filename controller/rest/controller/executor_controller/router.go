package executor_controller

import (
	"github.com/JasperStritzke/cubid-cloud/controller/executor"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func Mount(r *gin.RouterGroup, e *executor.Manager) {
	r.GET("/", middlewares.RequirePermission(auth.Permissions.ManageExecutors), ListExecutors(e))
	r.PUT("/create", middlewares.RequirePermission(auth.Permissions.ManageExecutors), CreateExecutor(e))
	r.DELETE("/:name", middlewares.RequirePermission(auth.Permissions.ManageExecutors), DeleteExecutor(e))
}
