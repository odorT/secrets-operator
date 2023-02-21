package main

import (
	"github.com/gin-contrib/cors"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"secrets-operator/config"
	"secrets-operator/internal/adapters/handlers/findingHdl"
	"secrets-operator/internal/adapters/handlers/searchHdl"
	"secrets-operator/internal/adapters/repositories/notification"
	"secrets-operator/internal/adapters/repositories/storage"
	"secrets-operator/internal/core/services/findingsrv"
	"time"
)

func main() {

	var err error

	//setup logger
	// TODO: move logging to new interface and handle everything there
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugaredLogger := logger.Sugar()

	sugaredLogger.Infoln("Secrets Operator starting ...")

	// load configs
	cfg, err := config.LoadConfig("dev")
	if err != nil {
		sugaredLogger.Fatalln("Cannot load configuration variables.", err)
	}

	sugaredLogger.Infof("Active running profile: %s", cfg.ActiveEnvProfile)

	// setup handlers, services, ports and etc
	findingsRepository := storage.NewMongoDb(cfg, sugaredLogger)
	notifier := notification.NewSlackNotifier(cfg, sugaredLogger)
	findingService := findingsrv.NewFindingService(sugaredLogger, findingsRepository, notifier)
	findingsHandler := findingHdl.NewFindingsHandler(cfg, sugaredLogger, findingService)
	searchHandler := searchHdl.NewSearchHandler(cfg, sugaredLogger, findingService)

	// setup http router
	router := setupRouter(logger)

	router.StaticFile("/api/v1/config.toml", cfg.ConfigFilePath)
	router.StaticFile("/api/v1/pipelineScript.sh", cfg.ScriptFilePath)

	findingsGroup := router.Group("/api/v1/findings")
	findingsGroup.POST("/upload", findingsHandler.Create)
	findingsGroup.GET("/:id", findingsHandler.Get)

	searchGroup := router.Group("/api/v1/search")
	searchGroup.GET("/repos", searchHandler.SearchRepositories)

	sugaredLogger.Fatalln(router.Run(cfg.ServerAddr))
}

func setupRouter(logger *zap.Logger) *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(ginZap.Ginzap(logger, time.RFC3339, true))

	return router
}
