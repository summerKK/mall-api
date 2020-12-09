package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/dto/admin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/model"
	businessError "github.com/summerKK/go-code-snippet-library/koel-api/pkg/error"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/security"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/util"
)

type AdminService struct {
	*service
}

func NewAdminService(ctx *gin.Context) *AdminService {
	svc := NewService(ctx)
	return &AdminService{
		service: svc,
	}
}

func (s *AdminService) CheckAuth(param *admin.UserLoginRequest) error {
	var err error
	user, err := s.dao.GetUserByName(param.UserName)
	if err != nil {
		return err
	}

	if security.VerifyPassword(user.Password, param.Password) {
		return nil
	}

	return errors.New("check auth failed.")
}

func (s *AdminService) Register(param *admin.UserRegisterRequest) (user *model.UmsAdmin, err error) {
	defer func() {
		util.AddErrorToCtx(s.ctx, err)
	}()

	user = param.Convert2Model()
	// 查看用户是否已经存在
	existsUser, err := s.dao.GetUserByName(user.Username)
	if err != nil {
		return
	}

	if existsUser != nil {
		return nil, businessError.NewBusinessError("用户已存在")
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return
	}

	user.Password = hashedPassword
	err = s.dao.Register(user)
	if err != nil {
		return
	}

	return
}
