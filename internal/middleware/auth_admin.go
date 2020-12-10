package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/app"
	errorCode "github.com/summerKK/go-code-snippet-library/koel-api/pkg/error"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/util"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		response := app.NewResponse(c)
		if bearToken == "" {
			response.ToErrorResponse(errorCode.UnauthorizedTokenError)
			c.Abort()
			return
		}

		bearToken = strings.Replace(bearToken, "Bearer ", "", 1)
		_, err := app.ParseToken(bearToken)
		if err != nil {
			util.AddErrorToCtx(c, err)
			response.ToErrorResponse(errorCode.UnauthorizedTokenError)
			c.Abort()
			return
		}

		c.Next()
	}
}
