package model

type UmsRoleMenuRelation struct {
	ID
	RoleId uint `json:"roleId"`
	MenuId uint `json:"menuId"`
}

func (r *UmsRoleMenuRelation) TableName() string {
	return "ums_role_menu_relation"
}
