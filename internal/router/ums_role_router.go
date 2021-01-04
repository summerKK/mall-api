package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/router/api/admin"
)

func UmsRoleRouter(r *gin.RouterGroup) {
	r1 := r.Group("/role")
	// 给角色分配菜单
	r1.POST("allocMenu", admin.UmsRoleController.AllocMenu)
	// 添加角色
	r1.POST("create", admin.UmsRoleController.Create)
}
