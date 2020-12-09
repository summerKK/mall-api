package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Translations(), middleware.CollectError())

	r1 := r.Group("/api")
	umsAdminRoute(r1)

	return r
}
