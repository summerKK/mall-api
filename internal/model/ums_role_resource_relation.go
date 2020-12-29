package model

type UmsRoleResourceRelation struct {
	ID
	RoleId     uint `json:"roleId"`
	ResourceId uint `json:"resourceId"`
}

func (r *UmsRoleResourceRelation) TableName() string {
	return "ums_role_resource_relation"
}
