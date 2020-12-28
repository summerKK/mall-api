package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/service"
	"github.com/summerKK/mall-api/pkg/app"
	errorCode "github.com/summerKK/mall-api/pkg/error"
	"github.com/summerKK/mall-api/pkg/util"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		response := app.NewResponse(c)
		var err error

		responseHandler := func() {
			util.AddErrorToCtx(c, err)
			response.Fail(errorCode.UnauthorizedTokenError)
			c.Abort()
			return
		}

		if bearToken == "" {
			responseHandler()
			return
		}

		bearToken = strings.Replace(bearToken, "Bearer ", "", 1)
		claims, err := app.ParseToken(bearToken)
		if err != nil {
			responseHandler()
			return
		}

		// 获取用户信息
		user, err := service.NewAdminService(c).GetItem(claims.UserId)
		if err != nil {
			responseHandler()
			return
		}

		c.Set("userInfo", user)
		c.Next()
	}
}
