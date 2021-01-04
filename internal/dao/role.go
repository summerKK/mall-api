package dao

import (
	"github.com/summerKK/mall-api/internal/model"
	"gorm.io/gorm"
)

type Role struct {
	*Dao
}

func NewRole(db *gorm.DB) *Role {
	return &Role{
		NewDao(db),
	}
}

func (r *Role) SyncRoleMenu(roleId uint, menuIds []uint) error {
	// 删除原有的关系
	err := r.db.Where("role_id = (?)", roleId).Delete(&model.UmsRoleMenuRelation{}).Error
	if err != nil {
		return err
	}

	relation := make([]model.UmsRoleMenuRelation, len(menuIds))
	for i, id := range menuIds {
		relation[i].RoleId = roleId
		relation[i].MenuId = id
	}

	return r.db.Create(relation).Error
}

func (r *Role) Create() {

}
