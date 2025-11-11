package users

import (
	"context"
	"giniladmin/internal/apps/admin/config"
	"giniladmin/internal/middleware"
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
	err := repo.AutoMigrate(repo.Db, &UserModelDB{})
	utils.CheckAndExit(err)

	routers.Register("user", func(router *gin.Engine) {
		v1Group := router.Group("api/v1/system/user")
		v1Group.Use(middleware.JWTAuth())
		//  POST /api/v1/system/user  (创建用户)
		v1Group.POST("", handler.HandleFailure(CreateUser))
		// GET /api/v1/system/user (获取用户列表)
		v1Group.GET("", handler.HandleFailure(GetUserList))
		// PUT /api/v1/system/user/{id}/password (修改密码)
		v1Group.PUT("/:id/password", handler.HandleFailure(ChangePassword))
		// GET /api/v1/system/user/{id} 获取指定ID用户
		v1Group.GET("/:id", handler.HandleFailure(GetUser))
		// PUT /api/v1/system/user/{id} 更新用户信息
		v1Group.PUT("/:id", handler.HandleFailure(UpdateUser))
		// DELETE /api/v1/system/user/{id} 删除用户
		v1Group.DELETE("/:id", handler.HandleFailure(DeleteUser))
	})
}
