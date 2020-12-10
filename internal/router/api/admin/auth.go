package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/dto/admin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/service"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/app"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/convert"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/error"
)

type umsAdminController struct{}

var UmsAdminController = umsAdminController{}

// 用户登录
func (_ umsAdminController) Login(c *gin.Context) {
	params := &admin.UserLoginRequest{}
	response := app.NewResponse(c)
	ok, errors := app.BindAndValid(c, params)
	if !ok {
		response.ToErrorResponse(error.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.NewAdminService(c)
	err := svc.Login(params)
	if err != nil {
		response.ToErrorResponse(error.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(params.UserName)
	if err != nil {
		response.ToErrorResponse(error.UnauthorizedTokenGenerate)
		return
	}

	response.Success(gin.H{
		"token": token,
	})
}

// 用户注册
func (_ umsAdminController) Register(c *gin.Context) {
	params := &admin.UserRegisterRequest{}
	response := app.NewResponse(c)
	ok, errors := app.BindAndValid(c, params)
	if !ok {
		response.ToErrorResponse(error.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.NewAdminService(c)
	user, err := svc.Register(params)
	if err != nil {
		response.ToErrorResponse(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(user)
}

// 获取指定用户信息
func (_ umsAdminController) GetItem(c *gin.Context) {
	p := c.Param("id")
	userId := convert.StrTo(p).MustInt()
	response := app.NewResponse(c)
	if userId == 0 {
		response.ToErrorResponse(error.InvalidParams)
		return
	}

	svc := service.NewAdminService(c)
	user, _ := svc.GetItem(userId)

	response.Success(user)
}
