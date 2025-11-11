package health

import (
	"context"
	"giniladmin/internal/apps/admin/config"
	"giniladmin/internal/middleware/handler"
	"giniladmin/internal/repository"
	"giniladmin/internal/routers"
	"giniladmin/pkg/logging"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
)

var (
	repo   *repository.Repository
	conf   *config.Config
	logger *logging.Logger
)

func Init(ctx context.Context) {
	v := ctx.Value("value").(map[string]any)

	// 初始化repo
	repo = v["repo"].(*repository.Repository)

	// conf
	conf = v["conf"].(*config.Config)

	// logger
	logger = v["log"].(*logging.Logger)

	// 初始化数据库表
	err := repo.AutoMigrate(repo.Db, &HealthModelDB{})
	utils.CheckAndExit(err)

	routers.Register("healthCheck", func(router *gin.Engine) {
		v1Group := router.Group("api/v1/healthCheck")
		v1Group.GET("/healthz", handler.HandleFailure(Healthz))
	})
}
