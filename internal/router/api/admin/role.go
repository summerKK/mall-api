package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/service"
	"github.com/summerKK/mall-api/pkg/error"
)

type umsRoleController struct {
	BaseAdmin
}

var UmsRoleController = umsRoleController{
	BaseAdminController,
}

func (u umsRoleController) AllocMenu(c *gin.Context) {
	params := &admin.RoleAllocMenuRequest{}
	ok, response := u.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewRoleService(c)
	err := svc.AllocMenu(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(gin.H{})
}

func (u umsRoleController) Create(c *gin.Context) {
	params := &admin.RoleCreateRequest{}
	ok, response := u.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewRoleService(c)
	err := svc.Create(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(gin.H{})
}
