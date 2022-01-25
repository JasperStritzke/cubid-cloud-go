package rest

import (
	"github.com/JasperStritzke/cubid-cloud/controller/executor"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/controller/auth_controller"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/controller/executor_controller"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/libs"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/middlewares"
	"github.com/JasperStritzke/cubid-cloud/util/cache"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/constants"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authManager *auth.Manager

func InitRouter(executorManager *executor.Manager) {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(
		middlewares.CORS, gin.Recovery(), cache.RateLimit,
		//SPA Host
		middlewares.SinglePageApp("/", "./web", "/api"),
	)

	engine.NoRoute(notFoundHandler(http.StatusNotFound))
	engine.NoMethod(notFoundHandler(http.StatusMethodNotAllowed))

	authManager = auth.NewManager()

	apiPath := engine.Group("/api")
	apiPath.GET("/status", libs.ResponseStaticMessage("Service is up and running."))

	//Mount routers
	auth_controller.Mount(apiPath.Group("/auth"), authManager)
	executor_controller.Mount(apiPath.Group("/executor"), executorManager)

	run(engine)
}

func run(engine *gin.Engine) {
	_ = config.InitConfigIfNotExists("ui.json", fileutil.CodingJSON, Config{Address: "0.0.0.0:6080"})

	var cfg Config
	_ = config.LoadConfig("ui.json", fileutil.CodingJSON, &cfg)
	url := cfg.Address

	var err error
	url = "http://" + url

	go func() {
		err = engine.Run(cfg.Address)
		if err != nil {
			logger.Warn("Unable to start rest-service")
		}
	}()

	if err == nil {
		logger.Info("Webinterface is running on: " + url)
	}
}

func notFoundHandler(statusCode int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(statusCode, gin.H{"message": "cubid-cloud version " + constants.Version})
	}
}
