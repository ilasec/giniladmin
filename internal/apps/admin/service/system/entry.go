package system

import (
	"context"
	"giniladmin/internal/apps/admin/service/system/application"
	"giniladmin/internal/apps/admin/service/system/groups"
	"giniladmin/internal/apps/admin/service/system/users"
)

func InitChilds(ctx context.Context) {
	users.Init(ctx)
	application.Init(ctx)
	groups.Init(ctx)
}

func Init(ctx context.Context) {
	//初始化子组件
	InitChilds(ctx)
}
