package model

import (
	"time"

	"gorm.io/gorm"
)

type UmsPermission struct {
	ID
	// 父级权限id
	Pid uint `json:"pid"`
	// 名称
	Name string `json:"name"`
	// 权限值
	Value string `json:"value"`
	//  图标
	Icon string `json:"icon"`
	// 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
	Type uint8 `json:"type"`
	// 前端资源路径
	Uri string `json:"uri"`
	// 启用状态；0->禁用；1->启用
	Status uint8 `json:"status"`
	// 排序
	Sort uint `json:"sort"`
	// 创建时间
	CreateTime LocalTime `json:"createTime"`
}

func (r *UmsPermission) TableName() string {
	return "ums_permission"
}

func (r *UmsPermission) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreateTime = LocalTime{time.Now()}

	return nil
}
