package service

import (
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

// 用户登录
func (s *AdminService) Login(param *admin.UserLoginRequest) error {
	var err error
	user, err := s.dao.GetItemByName(param.UserName)
	if err != nil {
		return err
	}

	if security.VerifyPassword(user.Password, param.Password) {
		return nil
	}

	return businessError.NewBusinessError("check auth failed")
}

// 用户注册
func (s *AdminService) Register(param *admin.UserRegisterRequest) (user *model.UmsAdmin, err error) {
	defer func() {
		util.AddErrorToCtx(s.ctx, err)
	}()

	user = param.Convert2Model()
	// 查看用户是否已经存在
	existsUser, err := s.dao.GetItemByName(user.Username)
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

// 获取指定用户信息
func (s *AdminService) GetItem(userId int) (user *model.UmsAdmin, err error) {
	return s.dao.GetItemById(userId)
}
