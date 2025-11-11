package application

import (
	"context"
	"errors"
	"giniladmin/pkg/utils"
	"net/http"
)

func CheckAppId(ctx context.Context, appId string) (s string, err error) {
	model := AppModelDB{}
	u, err := model.FindByAppId(ctx, repo.Db, appId)
	if err != nil || u == nil {
		err = errors.New("无效的AppId")
		return
	}
	s = u.Secret
	return
}

func DoCreateApp(ctx context.Context, App AppModel) (status int, message string, ret any, err error) {
	model := AppModelDB{}

	u, err := model.FindByName(ctx, repo.Db, App.Name)
	if err != nil {
		return
	}

	if u != nil {
		status = http.StatusConflict
		message = "APP已存在"
		return
	}

	App.AppId = utils.RandStringV2(8)
	App.Secret = utils.RandStringV2(16)

	err = model.Insert(ctx, repo.Db, App)
	return
}

func DoGetAppList(ctx context.Context, page int, pageSize int, keyword string) (status int, message string, ret any, err error) {
	model := AppModelDB{}

	Apps, total, err := model.FindApps(ctx, repo.Db, page, pageSize, keyword)
	if err != nil {
		ret = map[string]any{}
		return
	}

	//封装返回数据
	pageData := PageData{
		List:  Apps,
		Total: int(total),
	}
	//
	ret = pageData
	return
}

func DoGetApp(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := AppModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("APP不存在")
		return
	}
	ret = u
	return
}

func DoUpdateApp(ctx context.Context, App AppModel) (status int, message string, ret any, err error) {
	model := AppModelDB{}
	u, err := model.FindById(ctx, repo.Db, App.ID)
	if err != nil || u == nil {
		err = errors.New("APP不存在")
		return
	}
	_, err = model.Update(ctx, repo.Db, &App)
	return
}

func DoDeleteApp(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := AppModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("APP不存在")
		return
	}
	err = model.Remove(ctx, repo.Db, *u)
	return
}
