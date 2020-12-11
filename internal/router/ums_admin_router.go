package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/router/api/admin"
)

func umsAdminRouter(r *gin.RouterGroup) {
	r1 := r.Group("/admin")
	r1.GET("/:id", admin.UmsAdminController.GetItem)
	r1.POST("/delete/:id", admin.UmsAdminController.DeleteItem)
}
