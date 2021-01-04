package admin

import "github.com/summerKK/mall-api/internal/model"

type RoleAllocMenuRequest struct {
	MenuIds []uint `form:"menuIds" binding:"required"`
	RoleId  uint   `form:"roleId" binding:"required"`
}

type RoleCreateRequest struct {
	// 角色名称
	Name string `json:"name" binding:"required"`
	// 描述
	Description string `json:"description"`
	// 启用状态：0->禁用；1->启用
	Status uint8 `json:"status"`
}

func (r *RoleCreateRequest) Convert2Model() *model.UmsRole {
	return &model.UmsRole{
		Name:        r.Name,
		Description: r.Description,
		Status:      r.Status,
	}
}

type RoleAllocResourceRequest struct {
	ResourceIds []uint `json:"resourceIds" form:"resourceIds"`
	RoleId      uint   `json:"roleId" form:"roleId"`
}
