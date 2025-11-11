package application

import (
	"context"
	"encoding/gob"
	"giniladmin/internal/models"
	gorm_generics "giniladmin/pkg/gorm-generics"
	"gorm.io/gorm"
)

type PageData struct {
	List  []AppModel `json:"list"`
	Total int        `json:"total"`
}

type AppModel struct {
	models.GModel
	AppId  string `json:"appId" gorm:"index;comment:应用ID"` // 应用ID
	Secret string `json:"secret" gorm:"comment:应用密钥"`      // 应用密钥
	Name   string `json:"name" gorm:"index;comment:应用名称"`  // 应用名称
}

type AppModelDB struct {
	models.GModel
	AppId  string `json:"app_id" gorm:"index"`
	Secret string `json:"secret"`
	Name   string `json:"name" gorm:"index"`
}

// 将 AppModelDB 转换为 AppModel
func (g AppModelDB) ToEntity() AppModel {
	return AppModel{
		GModel: g.GModel,
		AppId:  g.AppId,
		Secret: g.Secret,
		Name:   g.Name,
	}
}

// 从 AppModel 生成 AppModelDB
func (g AppModelDB) FromEntity(m AppModel) interface{} {
	return AppModelDB{
		GModel: m.GModel,
		AppId:  m.AppId,
		Secret: m.Secret,
		Name:   m.Name,
	}
}

// 插入应用
func (g AppModelDB) Insert(ctx context.Context, db *gorm.DB, m AppModel) error {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	return repository.Insert(ctx, &m)
}

// 删除应用
func (g AppModelDB) Remove(ctx context.Context, db *gorm.DB, m AppModel) error {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	return repository.DeleteById(ctx, m.ID)
}

// 查找所有应用
func (g AppModelDB) FindAll(ctx context.Context, db *gorm.DB) ([]AppModel, error) {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	return repository.FindAll(ctx)
}

// 通过应用名称查找应用
func (g AppModelDB) FindByName(ctx context.Context, db *gorm.DB, Name string) (*AppModel, error) {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	Apps, err := repository.Find(ctx, gorm_generics.Equal("name", Name))
	if err == nil && len(Apps) > 0 {
		return &Apps[0], err
	}
	return nil, err
}

// 通过应用ID查找应用
func (g AppModelDB) FindById(ctx context.Context, db *gorm.DB, id int) (*AppModel, error) {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	Apps, err := repository.Find(ctx, gorm_generics.Equal("id", id))
	if err == nil && len(Apps) > 0 {
		return &Apps[0], err
	}
	return nil, err
}

// 通过应用ID查找应用
func (g AppModelDB) FindByAppId(ctx context.Context, db *gorm.DB, id string) (*AppModel, error) {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	Apps, err := repository.Find(ctx, gorm_generics.Equal("app_id", id))
	if err == nil && len(Apps) > 0 {
		return &Apps[0], err
	}
	return nil, err
}

// 更新应用信息
func (g AppModelDB) Update(ctx context.Context, db *gorm.DB, m *AppModel) (*AppModel, error) {
	repository := gorm_generics.NewRepository[AppModelDB, AppModel](db)
	err := repository.Update(ctx, m)
	return m, err
}

func (g AppModelDB) FindApps(ctx context.Context, db *gorm.DB, page, pageSize int, keyword string) ([]AppModel, int64, error) {
	query := db.Model(&AppModelDB{})

	// 关键字搜索
	if keyword != "" {
		query = query.Where("Name LIKE ?", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)

	var Apps []AppModel
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&Apps).Error

	return Apps, total, err
}

// 指定表名
func (AppModelDB) TableName() string {
	return "applications"
}

// 注册 AppModel 以支持 gob 序列化
func init() {
	gob.Register(AppModel{})
}
