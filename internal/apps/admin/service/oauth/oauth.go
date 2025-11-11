package oauth

import (
	"context"
	"giniladmin/internal/apps/admin/config"
	"giniladmin/internal/middleware"
	"giniladmin/internal/middleware/handler"
	"giniladmin/internal/repository"
	"giniladmin/internal/routers"
	"giniladmin/pkg/logging"
	cache "giniladmin/pkg/utils/cache"
	"github.com/gin-gonic/gin"
)

var (
	repo       *repository.Repository
	blackCache *cache.Cache
	conf       *config.Config
	logger     *logging.Logger
)

func Init(ctx context.Context) {
	v := ctx.Value("value").(map[string]any)

	// 初始化repo
	repo = v["repo"].(*repository.Repository)

	// conf
	conf = v["conf"].(*config.Config)

	// logger
	logger = v["log"].(*logging.Logger)

	// 初始化cache
	blackCache = v["bcache"].(*cache.Cache)

	// 初始化认证middleware
	middleware.JwtBlackCache = blackCache

	routers.Register("oauth", func(router *gin.Engine) {
		v1Group := router.Group("oauth/v1")
		v1Group.GET("/captcha", handler.HandleFailure(Captcha))
		v1Group.POST("/login", handler.HandleFailure(Login))
	})
}
