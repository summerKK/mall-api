package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/pkg/app"
	errorCode "github.com/summerKK/mall-api/pkg/error"
	"github.com/summerKK/mall-api/pkg/util"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		response := app.NewResponse(c)
		if bearToken == "" {
			response.Fail(errorCode.UnauthorizedTokenError)
			c.Abort()
			return
		}

		bearToken = strings.Replace(bearToken, "Bearer ", "", 1)
		_, err := app.ParseToken(bearToken)
		if err != nil {
			util.AddErrorToCtx(c, err)
			response.Fail(errorCode.UnauthorizedTokenError)
			c.Abort()
			return
		}

		c.Next()
	}
}
