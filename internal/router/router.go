package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Translations(), middleware.CollectError())

	r1 := r.Group("/api")
	// 公共接口
	publicRouter(r1)

	{
		r1.Use(middleware.AuthAdmin())
		// 后台用户管理
		umsAdminRouter(r1)
		// 商品管理
		pmsProductRouter(r1)
	}

	return r
}
