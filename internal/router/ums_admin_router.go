package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/middleware"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/router/api/admin"
)

func umsAdminRoute(r *gin.RouterGroup) {
	r1 := r.Group("/admin")

	r1.POST("/login", admin.UmsAdminController.Login)
	r1.POST("/register", admin.UmsAdminController.Register)

	r2 := r1.Use(middleware.AuthAdmin())
	r2.GET("/:id", admin.UmsAdminController.GetItem)
}
