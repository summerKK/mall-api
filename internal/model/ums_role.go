package model

import (
	"time"

	"gorm.io/gorm"
)

type UmsRole struct {
	ID
	// 角色名称
	Name string `json:"name"`
	// 描述
	Description string `json:"description"`
	// 后台用户数量
	AdminCount uint `json:"adminCount"`
	// 启用状态：0->禁用；1->启用
	Status uint8 `json:"status" gorm:"default:1"`
	Sort   uint  `json:"sort" gorm:"default:0"`
	// 创建时间
	CreateTime LocalTime `json:"createTime"`

	Menus []*UmsMenu `gorm:"many2many:ums_role_menu_relation;joinforeignKey:role_id;joinReferences:menu_id"`
}

func (r *UmsRole) TableName() string {
	return "ums_role"
}

func (r *UmsRole) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreateTime = LocalTime{time.Now()}

	return nil
}
