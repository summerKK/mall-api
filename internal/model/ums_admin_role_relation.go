package model

type UmsAdminRoleRelation struct {
	ID
	AdminId uint `json:"adminId"`
	RoleId  uint `json:"roleId"`
}

func (a *UmsAdminRoleRelation) TableName() string {
	return "ums_admin_role_relation"
}
