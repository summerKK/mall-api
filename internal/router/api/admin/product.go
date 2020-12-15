package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/service"
	"github.com/summerKK/mall-api/pkg/convert"
	"github.com/summerKK/mall-api/pkg/error"
)

type pmsProductController struct {
	BaseAdmin
}

var PmsProductController = pmsProductController{
	BaseAdminController,
}

func (p pmsProductController) Create(c *gin.Context) {
	productRequest := &admin.ProductRequest{}
	ok, response := p.VerifyParams(c, productRequest)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	product, err := svc.Create(productRequest)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(product)
}

func (p pmsProductController) Update(c *gin.Context) {
	productRequest := &admin.ProductRequest{}
	ok, response := p.VerifyParams(c, productRequest)
	if !ok {
		return
	}

	productIdS := c.Param("id")
	productId := convert.StrTo(productIdS).MustInt()
	if productId == 0 {
		response.Fail(error.InvalidParams)
		return
	}

	svc := service.NewProductService(c)
	svc.Update(productRequest, productId)
}
