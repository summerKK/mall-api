package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	UmsAdminStatusValid int8 = iota
	UmsAdminStatusInvalidValid
)

type UmsAdmin struct {
	ID
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Icon       string    `json:"icon"`
	Email      string    `json:"email"`
	NickName   string    `json:"nickName"`
	Note       string    `json:"note"`
	CreateTime LocalTime `json:"createTime"`
	// 设置gorm tag.这样调用create() 方法的时候,如果当前字段为零值会自动过滤掉
	LoginTime LocalTime `json:"loginTime" gorm:"default:null"`
	Status    uint8     `json:"status" gorm:"default:1"`
}

func (a *UmsAdmin) TableName() string {
	return "ums_admin"
}

func (a *UmsAdmin) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreateTime = LocalTime{time.Now()}

	return nil
}
