package repository

import (
	"context"
	"giniladmin/internal/configure"
	"giniladmin/pkg/utils"
	"gorm.io/gorm"
)

type Repository struct {
	Db     *gorm.DB
	dbtype string
}

type RepositoryModel interface {
	ToEntity() RepositoryModel
	FromEntity(entity RepositoryModel) interface{}
}

func NewRepository() *Repository {
	repo := Repository{}
	return &repo
}

func (p *Repository) Setup(ctx context.Context) (err error) {
	c := ctx.Value("value").(configure.DataBase)
	switch c.Type {
	case "mysql":
		p.Db = NewMysql(&c)
	case "sqlite":
		p.Db = NewSqlite(&c)
	}
	if p.Db == nil {
		//err = errors.New("invalid db configure")
	}
	utils.CheckAndExit(err)
	return
}

func (p *Repository) AutoMigrate(db *gorm.DB, dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}

func (p *Repository) MySqlDsn(username, password, path, port, dbname string) string {
	return username + ":" + password + "@tcp(" + path + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
}
