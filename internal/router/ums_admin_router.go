package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/router/api/admin"
)

func umsAdminRoute(r *gin.RouterGroup) {
	group := r.Group("/admin")

	group.GET("/:id", admin.UmsAdminController.GetItem)
	group.POST("/login", admin.UmsAdminController.Login)
	group.POST("/register", admin.UmsAdminController.Register)
}
