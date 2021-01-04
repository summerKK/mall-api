package admin

type RoleAllocMenuRequest struct {
	MenuIds []uint `form:"menuIds" binding:"required"`
	RoleId  uint   `form:"roleId" binding:"required"`
}
