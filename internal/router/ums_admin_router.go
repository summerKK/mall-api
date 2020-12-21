package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/router/api/admin"
)

func umsAdminRouter(r *gin.RouterGroup) {
	r1 := r.Group("/admin")
	// 获取指定用户信息
	r1.GET("/:id", admin.UmsAdminController.GetItem)
	// 删除指定用户信息
	r1.POST("/delete/:id", admin.UmsAdminController.DeleteItem)
}
