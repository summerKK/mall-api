package model

type UmsRolePermissionRelation struct {
	ID
	RoleId       uint `json:"roleId"`
	PermissionId uint `json:"permissionId"`
}

func (r *UmsRolePermissionRelation) TableName() string {
	return "ums_role_permission_relation"
}
