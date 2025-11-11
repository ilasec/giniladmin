package association

import (
	"context"
	"errors"
)

func DoCreateAssociation(ctx context.Context, Association AssociationModel) (status int, message string, ret any, err error) {
	model := AssociationModelDB{}
	err = model.Insert(ctx, repo.Db, Association)
	return
}

func DoGetAssociationList(ctx context.Context, page int, pageSize int, keyword string) (status int, message string, ret any, err error) {
	model := AssociationModelDB{}

	Associations, total, err := model.FindAssociations(ctx, repo.Db, page, pageSize, keyword)
	if err != nil {
		ret = map[string]any{}
		return
	}

	//封装返回数据
	pageData := PageData{
		List:  Associations,
		Total: int(total),
	}
	//
	ret = pageData
	return
}

func DoGetAssociation(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := AssociationModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("Association不存在")
		return
	}
	ret = u
	return
}

func DoUpdateAssociation(ctx context.Context, Association AssociationModel) (status int, message string, ret any, err error) {
	model := AssociationModelDB{}
	u, err := model.FindById(ctx, repo.Db, Association.ID)
	if err != nil || u == nil {
		err = errors.New("Association不存在")
		return
	}
	_, err = model.Update(ctx, repo.Db, &Association)
	return
}

func DoDeleteAssociation(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := AssociationModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("Association不存在")
		return
	}
	err = model.Remove(ctx, repo.Db, *u)
	return
}
