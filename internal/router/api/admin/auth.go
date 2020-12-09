package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/dto/admin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/service"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/app"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/error"
)

type umsAdminController struct{}

var UmsAdminController = umsAdminController{}

func (_ umsAdminController) Login(c *gin.Context) {
	params := &admin.UserLoginRequest{}
	response := app.NewResponse(c)
	ok, errors := app.BindAndValid(c, params)
	if !ok {
		response.ToErrorResponse(error.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.NewAdminService(c)
	err := svc.CheckAuth(params)
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
