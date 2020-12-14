package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/service"
	"github.com/summerKK/mall-api/pkg/app"
	"github.com/summerKK/mall-api/pkg/error"
)

type pmsProductController struct{}

var PmsProductController = pmsProductController{}

func (_ pmsProductController) Create(c *gin.Context) {
	productRequest := &admin.ProductRequest{}
	response := app.NewResponse(c)
	ok, errors := app.BindAndValid(c, productRequest)
	if !ok {
		response.Fail(error.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.NewProductService(c)
	err := svc.Create(productRequest)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}
