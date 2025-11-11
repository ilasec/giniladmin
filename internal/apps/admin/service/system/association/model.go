package association

import (
	"context"
	"encoding/gob"
	"giniladmin/internal/models"
	gorm_generics "giniladmin/pkg/gorm-generics"
	"gorm.io/gorm"
)

type PageData struct {
	List  []AssociationModel `json:"list"`
	Total int                `json:"total"`
}

type AssociationModel struct {
	models.GModel
	UserID  int `json:"userId"`
	AppID   int `json:"appId"`
	GroupID int `json:"groupId"` // Added GroupID field
}

type AssociationModelDB struct {
	models.GModel
	UserID  int `json:"user_id"`
	AppID   int `json:"app_id"`
	GroupID int `json:"group_id"` // Added GroupID field
}

// 将 AssociationModelDB 转换为 AssociationModel
func (g AssociationModelDB) ToEntity() AssociationModel {
	return AssociationModel{
		GModel:  g.GModel,
		UserID:  g.UserID,
		AppID:   g.AppID,
		GroupID: g.GroupID,
	}
}

// 从 AssociationModel 生成 AssociationModelDB
func (g AssociationModelDB) FromEntity(m AssociationModel) interface{} {
	return AssociationModelDB{
		GModel:  m.GModel,
		UserID:  m.UserID,
		AppID:   m.AppID,
		GroupID: m.GroupID,
	}
}

// 插入应用
func (g AssociationModelDB) Insert(ctx context.Context, db *gorm.DB, m AssociationModel) error {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	return repository.Insert(ctx, &m)
}

// 删除应用
func (g AssociationModelDB) Remove(ctx context.Context, db *gorm.DB, m AssociationModel) error {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	return repository.DeleteById(ctx, m.ID)
}

// 查找所有应用
func (g AssociationModelDB) FindAll(ctx context.Context, db *gorm.DB) ([]AssociationModel, error) {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	return repository.FindAll(ctx)
}

// 通过应用名称查找应用
func (g AssociationModelDB) FindByName(ctx context.Context, db *gorm.DB, Name string) (*AssociationModel, error) {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	Associations, err := repository.Find(ctx, gorm_generics.Equal("name", Name))
	if err == nil && len(Associations) > 0 {
		return &Associations[0], err
	}
	return nil, err
}

// 通过应用ID查找应用
func (g AssociationModelDB) FindById(ctx context.Context, db *gorm.DB, id int) (*AssociationModel, error) {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	Associations, err := repository.Find(ctx, gorm_generics.Equal("id", id))
	if err == nil && len(Associations) > 0 {
		return &Associations[0], err
	}
	return nil, err
}

// 通过应用ID查找应用
func (g AssociationModelDB) FindByAssociationId(ctx context.Context, db *gorm.DB, id string) (*AssociationModel, error) {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	Associations, err := repository.Find(ctx, gorm_generics.Equal("Association_id", id))
	if err == nil && len(Associations) > 0 {
		return &Associations[0], err
	}
	return nil, err
}

// 更新应用信息
func (g AssociationModelDB) Update(ctx context.Context, db *gorm.DB, m *AssociationModel) (*AssociationModel, error) {
	repository := gorm_generics.NewRepository[AssociationModelDB, AssociationModel](db)
	err := repository.Update(ctx, m)
	return m, err
}

func (g AssociationModelDB) FindAssociations(ctx context.Context, db *gorm.DB, page, pageSize int, keyword string) ([]AssociationModel, int64, error) {
	query := db.Model(&AssociationModelDB{})

	// 关键字搜索
	if keyword != "" {
		query = query.Where("Name LIKE ?", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)

	var Associations []AssociationModel
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&Associations).Error

	return Associations, total, err
}

// 指定表名
func (AssociationModelDB) TableName() string {
	return "association"
}

// 注册 AssociationModel 以支持 gob 序列化
func init() {
	gob.Register(AssociationModel{})
}
