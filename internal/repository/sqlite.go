package repository

import (
	"giniladmin/internal/configure"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

type Sqlite struct {
	*configure.DataBase
}

func (m *Sqlite) Dsn() string {
	return filepath.Join(m.Path, m.Dbname+".db")
}

func (m *Sqlite) GetLogMode() string {
	return m.LogMode
}

func NewSqlite(c *configure.DataBase) *gorm.DB {
	m := Sqlite{c}
	if m.Dbname == "" {

		return nil
	}
	if db, err := gorm.Open(sqlite.Open(m.Dsn())); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		if m.GetLogMode() == "debug" {
			db = db.Debug()
		}
		return db
	}
}
