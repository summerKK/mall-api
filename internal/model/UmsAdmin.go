package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	UmsAdminStatusValid int8 = iota
	UmsAdminStatusInvalidValid
)

type UmsAdmin struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Icon       string    `json:"icon"`
	Email      string    `json:"email"`
	NickName   string    `json:"nick_name"`
	Note       string    `json:"note"`
	CreateTime time.Time `json:"create_time"`
	// 设置gorm tag.这样调用create() 方法的时候,如果当前字段为零值会自动过滤掉
	LoginTime time.Time `json:"login_time" gorm:"default:null"`
	Status    uint8     `json:"status" gorm:"default:1"`
}

func (a *UmsAdmin) TableName() string {
	return "ums_admin"
}

func (a *UmsAdmin) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreateTime = time.Now()

	return nil
}
