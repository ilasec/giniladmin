package users

import (
	"context"
	"encoding/gob"
	"giniladmin/internal/models"
	gorm_generics "giniladmin/pkg/gorm-generics"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PageData struct {
	List  []UserModel `json:"list"`
	Total int         `json:"total"`
}

type User struct {
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	Password string `json:"password"  gorm:"comment:用户登录密码"`
}

type UserModel struct {
	models.GModel
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	AuthorityId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Enable      int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                      // 用户是否被冻结 1正常 2冻结
	IsSystem    int       `json:"system" gorm:"default:0;comment:是否系统用户 1是 2否"`
}

type UserModelDB struct {
	models.GModel
	UUID        uuid.UUID `json:"uuid" gorm:"index"`
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	AuthorityId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Enable      int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	IsSystem    int       `json:"isSystem" gorm:"default:0;comment:是否系统用户 1是 2否"`
}

// 将 UserModelDB 转换为 UserModel
func (g UserModelDB) ToEntity() UserModel {
	return UserModel{
		GModel:      g.GModel,
		UUID:        g.UUID,
		Username:    g.Username,
		Password:    g.Password,
		NickName:    g.NickName,
		HeaderImg:   g.HeaderImg,
		AuthorityId: g.AuthorityId,
		Phone:       g.Phone,
		Email:       g.Email,
		Enable:      g.Enable,
		IsSystem:    g.IsSystem,
	}
}

// 从 UserModel 生成 UserModelDB
func (g UserModelDB) FromEntity(m UserModel) interface{} {
	return UserModelDB{
		GModel:      m.GModel,
		UUID:        m.UUID,
		Username:    m.Username,
		Password:    m.Password,
		NickName:    m.NickName,
		HeaderImg:   m.HeaderImg,
		AuthorityId: m.AuthorityId,
		Phone:       m.Phone,
		Email:       m.Email,
		Enable:      m.Enable,
		IsSystem:    m.IsSystem,
	}
}

// 插入用户
func (g UserModelDB) Insert(ctx context.Context, db *gorm.DB, m UserModel) error {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	return repository.Insert(ctx, &m)
}

// 删除用户
func (g UserModelDB) Remove(ctx context.Context, db *gorm.DB, m UserModel) error {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	return repository.DeleteById(ctx, m.ID)
}

// 查找所有用户
func (g UserModelDB) FindAll(ctx context.Context, db *gorm.DB) ([]UserModel, error) {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	return repository.FindAll(ctx)
}

// 通过用户名查找用户
func (g UserModelDB) FindByUsername(ctx context.Context, db *gorm.DB, username string) (*UserModel, error) {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	users, err := repository.Find(ctx, gorm_generics.Equal("username", username))
	if err == nil && len(users) > 0 {
		return &users[0], err
	}
	return nil, err
}

// 通过用户id找用户
func (g UserModelDB) FindById(ctx context.Context, db *gorm.DB, id int) (*UserModel, error) {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	users, err := repository.Find(ctx, gorm_generics.Equal("id", id))
	if err == nil && len(users) > 0 {
		return &users[0], err
	}
	return nil, err
}

// 更新用户信息
func (g UserModelDB) Update(ctx context.Context, db *gorm.DB, m *UserModel) (*UserModel, error) {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	err := repository.Update(ctx, m)
	return m, err
}

func (g UserModelDB) UpdateEx(ctx context.Context, db *gorm.DB, m *UserModel) (*UserModel, error) {
	repository := gorm_generics.NewRepository[UserModelDB, UserModel](db)
	err := repository.Update(ctx, m)
	return m, err
}

func (g UserModelDB) FindUsers(ctx context.Context, db *gorm.DB, page, pageSize int, keyword string) ([]UserModel, int64, error) {
	query := db.Model(&UserModelDB{})

	// 关键字搜索
	if keyword != "" {
		query = query.Where("username LIKE ? OR nick_name LIKE ? OR phone LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)

	var users []UserModel
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error

	return users, total, err
}

// 指定表名
func (UserModelDB) TableName() string {
	return "users"
}

// 注册 UserModel 以支持 gob 序列化
func init() {
	gob.Register(UserModel{})
}
