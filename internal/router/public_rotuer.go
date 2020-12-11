package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/router/api/admin"
)

func publicRouter(r *gin.RouterGroup) {
	{
		r1 := r.Group("/admin")
		r1.POST("/login", admin.UmsAdminController.Login)
		r1.POST("/register", admin.UmsAdminController.Register)
	}

}
