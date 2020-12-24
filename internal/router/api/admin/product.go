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

// 批量修改删除状态
func (p pmsProductController) BatchSetDeleteStatus(c *gin.Context) {
	params := &admin.ProductBatchDeleteStatusRequest{}
	ok, response := p.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	err := svc.BatchSetDeleteStatus(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}

// 批量设为新品
func (p pmsProductController) BatchSetNewStatus(c *gin.Context) {
	params := &admin.ProductBatchSetNewStatusRequest{}
	ok, response := p.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	err := svc.BatchSetNewStatus(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}

func (p pmsProductController) BatchSetPublishStatus(c *gin.Context) {
	params := &admin.ProductBatchSetPublishStatusRequest{}
	ok, response := p.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	err := svc.BatchSetPublishStatus(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}

func (p pmsProductController) BatchSetRecommendStatus(c *gin.Context) {
	params := &admin.ProductBatchSetRecommendStatusRequest{}
	ok, response := p.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	err := svc.BatchSetRecommendStatus(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}

func (p pmsProductController) BatchSetVerifyStatus(c *gin.Context) {
	params := &admin.ProductBatchSetVerifyStatusRequest{}
	ok, response := p.VerifyParams(c, params)
	if !ok {
		return
	}

	svc := service.NewProductService(c)
	err := svc.BatchSetVerifyStatus(params)
	if err != nil {
		response.Fail(error.NewErrWithBusinessError(err))
		return
	}

	response.Success(nil)
}
