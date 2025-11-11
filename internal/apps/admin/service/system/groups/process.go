package groups

import (
	"context"
	"errors"
	"giniladmin/internal/apps/admin/service/system/application"
	"giniladmin/pkg/utils/structutils"
	"net/http"
	"time"
)

func DoCreateGroup(ctx context.Context, group GroupModel) (status int, message string, ret any, err error) {
	model := GroupModelDB{}
	appModel := application.AppModelDB{}

	app, err := appModel.FindByAppId(ctx, repo.Db, group.AppId)
	if err != nil || app == nil {
		status = http.StatusConflict
		err = errors.New("app不存在")
		return
	}

	u, err := model.FindByNames(ctx, repo.Db, group.Name, group.AppId)
	if err != nil {
		return
	}

	if u != nil {
		status = http.StatusConflict
		err = errors.New("Group已存在")
		return
	}

	err = model.Insert(ctx, repo.Db, group)
	return
}

func DoGetGroupList(ctx context.Context, page int, pageSize int, keyword string) (status int, message string, ret any, err error) {
	model := GroupModelDB{}

	groups, total, err := model.FindGroups(ctx, repo.Db, page, pageSize, keyword)
	if err != nil {
		ret = map[string]any{}
		return
	}

	//封装返回数据
	pageData := PageData{
		List:  groups,
		Total: int(total),
	}
	//
	ret = pageData
	return
}

func DoGetGroup(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := GroupModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("Group不存在")
		return
	}
	ret = u
	return
}

func DoUpdateGroup(ctx context.Context, group GroupModel) (status int, message string, ret any, err error) {
	model := GroupModelDB{}
	u, err := model.FindById(ctx, repo.Db, group.ID)
	if err != nil || u == nil {
		err = errors.New("Group不存在")
		return
	}

	u.UpdatedAt = time.Time{}
	ignored := []string{"GModel"}        // 忽略 CreatedAt 和 UpdatedAt 字段
	zeroChecks := []string{"Permission"} // 忽略 CreatedAt 和 UpdatedAt 字段

	updates, err := structutils.CompareStructs(u, &group, ignored, zeroChecks)
	if err != nil {
		err = errors.New("未知错误")
		return
	}
	if len(updates) > 0 {
		// updates 包含了不同的字段和值
		// 可以使用 updates 更新数据库或其他操作
		_, err = model.Update(ctx, repo.Db, u)

	}
	return
}

func DoDeleteGroup(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := GroupModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("Group不存在")
		return
	}
	err = model.Remove(ctx, repo.Db, *u)
	return
}
