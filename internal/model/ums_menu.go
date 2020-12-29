package model

import (
	"time"

	"gorm.io/gorm"
)

type UmsMenu struct {
	ID
	// 父级Id
	ParentId uint `json:"parentId"`
	// 菜单名称
	Title string `json:"title"`
	// 菜单级数
	Level uint8 `json:"level"`
	// 菜单排序
	Sort int `json:"sort" gorm:"default:0"`
	// 前端名称
	Name string `json:"name"`
	// 前端图标
	Icon string `json:"icon"`
	// 前端隐藏
	Hidden uint8 `json:"hidden"`
	// 创建时间
	CreateTime LocalTime `json:"createTime"`
}

func (m *UmsMenu) TableName() string {
	return "ums_menu"
}

func (m *UmsMenu) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateTime = LocalTime{time.Now()}

	return nil
}
