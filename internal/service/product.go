package service

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/global"
	"github.com/summerKK/mall-api/internal/dao"
)

type ProductService struct {
	service *service
	dao     *dao.ProductDao
}

func NewProductService(ctx *gin.Context) *ProductService {
	service := NewService(ctx)
	return &ProductService{
		service: service,
		dao:     dao.NewProduct(global.DBEngine),
	}
}
