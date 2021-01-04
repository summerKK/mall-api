package service

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/global"
	"github.com/summerKK/mall-api/internal/dao"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/model"
	businessError "github.com/summerKK/mall-api/pkg/error"
	"github.com/summerKK/mall-api/pkg/util"
)

type RoleService struct {
	service *service
	dao     *dao.Role
}

func NewRoleService(ctx *gin.Context) *RoleService {
	svc := NewService(ctx)
	return &RoleService{
		service: svc,
		dao:     dao.NewRole(global.DBEngine),
	}
}

func (r *RoleService) AllocMenu(params *admin.RoleAllocMenuRequest) (err error) {
	defer func() {
		util.AddErrorToCtx(r.service.ctx, err)
	}()

	role := &model.UmsRole{}
	ok, err := r.dao.GetItemById(params.RoleId, role)
	if err != nil {
		return
	}

	if !ok {
		err = businessError.NewBusinessError("角色不存在")
		return
	}

	err = r.dao.SyncRoleMenu(params.RoleId, params.MenuIds)

	return
}

func (r *RoleService) Create(params *admin.RoleCreateRequest) (err error) {
	defer func() {
		util.AddErrorToCtx(r.service.ctx, err)
	}()

	role := params.Convert2Model()
	err = r.dao.Insert(role)

	return
}
