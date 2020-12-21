package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/router/api/admin"
)

func pmsProductRouter(r *gin.RouterGroup) {
	r1 := r.Group("/product")
	r1.POST("/create", admin.PmsProductController.Create)
	r1.POST("/update/:id", admin.PmsProductController.Update)
	r1.GET("/list", admin.PmsProductController.List)
	r1.GET("/simpleList", admin.PmsProductController.SimpleList)
	r1.POST("/batchUpdate/deleteStatus", admin.PmsProductController.BatchDeleteStatus)
}
