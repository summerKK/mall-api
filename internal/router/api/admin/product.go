package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/service"
	"github.com/summerKK/mall-api/pkg/app"
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
	product, err := svc.Update(productRequest, productId)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(product)
}

func (p pmsProductController) List(c *gin.Context) {
	productListRequest := &admin.ProductListRequest{}
	ok, response := p.VerifyParams(c, productListRequest)
	if !ok {
		return
	}

	pageSize := app.GetPageSize(c)
	pageNum := app.GetPage(c)
	pageOffset := app.GetPageOffset(pageNum, pageSize)

	svc := service.NewProductService(c)
	list, count, err := svc.List(productListRequest, pageSize, pageOffset)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.SuccessWithPage(list, int(count), pageNum, pageSize)
}

// 根据商品名称或货号模糊查询
func (p pmsProductController) SimpleList(c *gin.Context) {
	keyword := c.Query("keyword")
	response := app.NewResponse(c)

	svc := service.NewProductService(c)
	list, err := svc.SimpleList(keyword)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(list)
}

func (p pmsProductController) BatchDeleteStatus(c *gin.Context) {
	params := &admin.ProductBatchDeleteStatusRequest{}
	ok, response := p.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	err := svc.BatchDeleteStatus(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}
