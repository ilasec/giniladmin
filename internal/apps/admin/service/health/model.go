package health

import (
	"context"
	"encoding/gob"
	"giniladmin/internal/models"
	gorm_generics "giniladmin/pkg/gorm-generics"
	"gorm.io/gorm"
)

type HealthModel struct {
	models.GModel
	Name string `json:"name" form:"name" db:"name" gorm:"unique"`
}

type HealthModelDB struct {
	models.GModel
	Name string `json:"name" form:"name" db:"name"`
}

func (g HealthModelDB) ToEntity() HealthModel {
	return HealthModel{
		// some fields
		GModel: models.GModel{ID: g.ID, Gid: g.Gid, Uid: g.Uid, Pid: g.Pid, CreatedAt: g.CreatedAt, UpdatedAt: g.UpdatedAt},
		Name:   g.Name,
	}
}

func (g HealthModelDB) FromEntity(m HealthModel) interface{} {
	return HealthModelDB{
		// some fields
		GModel: models.GModel{ID: m.ID, Gid: m.Gid, Uid: m.Uid, Pid: m.Pid, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt},
		Name:   m.Name,
	}
}

func (g HealthModelDB) Insert(ctx context.Context, db *gorm.DB, m HealthModel) (err error) {
	repository := gorm_generics.NewRepository[HealthModelDB, HealthModel](db)
	err = repository.Insert(ctx, &m)

	return
}

func (g HealthModelDB) Remove(ctx context.Context, db *gorm.DB, m HealthModel) (err error) {
	repository := gorm_generics.NewRepository[HealthModelDB, HealthModel](db)
	err = repository.DeleteById(ctx, m.ID)
	return
}

func (g HealthModelDB) FindAll(ctx context.Context, db *gorm.DB) ([]HealthModel, error) {
	repository := gorm_generics.NewRepository[HealthModelDB, HealthModel](db)
	all, err := repository.FindAll(ctx)
	return all, err
}

func (g HealthModelDB) FindByName(ctx context.Context, db *gorm.DB, name string) (*HealthModel, error) {
	repository := gorm_generics.NewRepository[HealthModelDB, HealthModel](db)
	all, err := repository.Find(ctx, gorm_generics.Equal("name", name))
	if err == nil && len(all) > 0 {
		return &all[0], err
	}
	return nil, err
}

func (g HealthModelDB) Update(ctx context.Context, db *gorm.DB, m *HealthModel) (*HealthModel, error) {
	repository := gorm_generics.NewRepository[HealthModelDB, HealthModel](db)
	err := repository.Update(ctx, m)
	return m, err
}

func (HealthModelDB) TableName() string {
	return "health"
}

func init() {
	gob.Register(HealthModel{})
}
