package executor_controller

import (
	"github.com/JasperStritzke/cubid-cloud/controller/executor"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/libs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ListExecutors(e *executor.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"executors": e.ExecutorInfos()})
	}
}

type createExecutorRequest struct {
	Name        string `json:"name"`
	MaxCpuUsage int    `json:"maxCpuUsage"`
	MaxMemory   int    `json:"maxMemory"`
}

func CreateExecutor(e *executor.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createBody createExecutorRequest
		if !libs.BindBody(ctx, &createBody) {
			return
		}

		if len(createBody.Name) > 1 &&
			createBody.MaxCpuUsage >= 0 && createBody.MaxCpuUsage <= 100 &&
			!strings.ContainsAny(createBody.Name, "?") &&
			createBody.MaxMemory >= 128 && e.CreateExecutor(createBody.Name, createBody.MaxCpuUsage, createBody.MaxMemory) {
			libs.RespondWithMessage(200, "Executor created", ctx)
			return
		}

		libs.RespondWithMessage(http.StatusBadRequest, "Executor already exists", ctx)
	}
}

func DeleteExecutor(e *executor.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		if name == "" {
			libs.RespondWithMessage(http.StatusBadRequest, "Name not provided!", ctx)
			return
		}

		if e.DeleteExecutor(name) {
			libs.RespondWithMessage(http.StatusOK, "Executor deleted.", ctx)
			return
		}

		libs.RespondWithMessage(http.StatusBadRequest, "Executor doesn't exist!", ctx)
	}
}
