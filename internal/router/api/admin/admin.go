package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/pkg/app"
	errorCode "github.com/summerKK/mall-api/pkg/error"
)

type BaseAdmin struct{}

var BaseAdminController = BaseAdmin{}

// 验证参数
func (b BaseAdmin) VerifyParams(c *gin.Context, requestStruct interface{}) (ok bool, response *app.Response) {
	response = app.NewResponse(c)
	ok, errors := app.BindAndValid(c, requestStruct)
	if !ok {
		response.Fail(errorCode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	return true, response
}
