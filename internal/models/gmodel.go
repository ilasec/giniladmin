package models

import (
	"gorm.io/gorm"
	"time"
)

type GModel struct {
	ID        int            `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // 主键ID
	Gid       int            `gorm:"gid"`                                  //group id
	Uid       int            `gorm:"uid"`                                  //user id
	Pid       int            `gorm:"pid"`                                  //parament id
	CreatedAt time.Time      `gorm:"createdAt" json:"createdAt"`           // 创建时间
	UpdatedAt time.Time      `gorm:"updatedAt" json:"updatedAt"`           // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                       // 删除时间
}
