package groups

import (
	"context"
	"database/sql/driver"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"giniladmin/internal/models"
	gorm_generics "giniladmin/pkg/gorm-generics"
	"gorm.io/gorm"
)

type PageData struct {
	List  []GroupModel `json:"list"`
	Total int          `json:"total"`
}

type GroupModel struct {
	models.GModel
	AppId      string     `json:"appId" gorm:"index;comment:所属应用ID"`
	Name       string     `json:"name" gorm:"comment:用户组名称"`
	Permission Permission `json:"permission" gorm:"type:json;comment:用户组权限"`
}

type GroupModelDB struct {
	models.GModel
	AppId      string     `json:"app_id" gorm:"index;foreignKey:app_id"`
	Name       string     `json:"name"`
	Permission Permission `json:"permission" gorm:"type:json"`
}

// 将 UserGroupDB 转换为 UserGroup
func (g GroupModelDB) ToEntity() GroupModel {
	return GroupModel{
		GModel:     g.GModel,
		AppId:      g.AppId,
		Name:       g.Name,
		Permission: g.Permission,
	}
}

// 从 UserGroup 生成 UserGroupDB
func (g GroupModelDB) FromEntity(m GroupModel) interface{} {
	return GroupModelDB{
		GModel:     m.GModel,
		AppId:      m.AppId,
		Name:       m.Name,
		Permission: m.Permission,
	}
}

// 插入
func (g GroupModelDB) Insert(ctx context.Context, db *gorm.DB, m GroupModel) error {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	return repository.Insert(ctx, &m)
}

// 删除
func (g GroupModelDB) Remove(ctx context.Context, db *gorm.DB, m GroupModel) error {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	return repository.DeleteById(ctx, m.ID)
}

// 查找
func (g GroupModelDB) FindAll(ctx context.Context, db *gorm.DB) ([]GroupModel, error) {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	return repository.FindAll(ctx)
}

// 通过名称查找
func (g GroupModelDB) FindByName(ctx context.Context, db *gorm.DB, Name string) (*GroupModel, error) {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	groups, err := repository.Find(ctx, gorm_generics.Equal("name", Name))
	if err == nil && len(groups) > 0 {
		return &groups[0], err
	}
	return nil, err
}

// 通过名称查找
func (g GroupModelDB) FindByNames(ctx context.Context, db *gorm.DB, name string, appId string) (*GroupModel, error) {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	groups, err := repository.Find(ctx, gorm_generics.And(gorm_generics.Equal("app_id", appId), gorm_generics.Equal("name", name)))
	if err == nil && len(groups) > 0 {
		return &groups[0], err
	}
	return nil, err
}

// 通过ID查找
func (g GroupModelDB) FindById(ctx context.Context, db *gorm.DB, id int) (*GroupModel, error) {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	groups, err := repository.Find(ctx, gorm_generics.Equal("id", id))
	if err == nil && len(groups) > 0 {
		return &groups[0], err
	}
	return nil, err
}

// 更新信息
func (g GroupModelDB) Update(ctx context.Context, db *gorm.DB, m *GroupModel) (*GroupModel, error) {
	repository := gorm_generics.NewRepository[GroupModelDB, GroupModel](db)
	err := repository.Update(ctx, m)
	return m, err
}

func (g GroupModelDB) FindGroups(ctx context.Context, db *gorm.DB, page, pageSize int, keyword string) ([]GroupModel, int64, error) {
	query := db.Model(&GroupModelDB{})

	// 关键字搜索
	if keyword != "" {
		query = query.Where("Name LIKE ?", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)

	var groups []GroupModel
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&groups).Error

	// 子查询关联 AppModelDB 获取应用名称
	//err := query.
	//	Preload("applications"). // 预加载关联的 App 数据
	//	Offset((page - 1) * pageSize).
	//	Limit(pageSize).
	//	Find(&groups).Error

	return groups, total, err
}

// 指定表名
func (GroupModelDB) TableName() string {
	return "groups" // 建议表名使用复数形式
}

// 注册 GroupModel 以支持 gob 序列化
func init() {
	gob.Register(GroupModel{})
}

// 自定义 Permission 类型，用于存储 JSON 数据
type Permission map[string][]string

// 实现 gorm.Valuer 和 gorm.Scanner 接口，用于 JSON 字段的存储和读取
func (p Permission) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Permission) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unexpected type %T for Permission", value)
	}
	return json.Unmarshal(bytes, p)
}
