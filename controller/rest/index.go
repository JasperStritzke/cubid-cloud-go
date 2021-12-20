package rest

import (
	"github.com/JasperStritzke/cubid-cloud/controller/rest/libs"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/middlewares"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Use(func(ctx *gin.Context) {
		ctx.Header("Server", "cubid-cloud-v"+constants.Version)
		ctx.Next()
	})

	//SPA Host
	engine.Use(middlewares.SinglePageApp("/", "./controller-frontend/", "/api"))

	engine.NoRoute(notFoundHandler(http.StatusNotFound))
	engine.NoMethod(notFoundHandler(http.StatusMethodNotAllowed))

	apiPath := engine.Group("/api")
	apiPath.GET("/status", libs.ResponseStaticMessage("Service is up and running."))

	run(engine)
}

func run(engine *gin.Engine) {
	_ = config.InitConfigIfNotExists("ui.json", func() interface{} {
		return Config{
			Address: "0.0.0.0:6080",
		}
	})

	var cfg Config
	_ = config.LoadConfig("ui.json", &cfg)
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

func notFoundHandler(statusCode int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(statusCode, gin.H{"message": "cubid-cloud version " + constants.Version})
	}
}
