package middleware

import (
	"bytes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/util"
)

//  收集运行时的错误日志并打印
func CollectError() gin.HandlerFunc {
	key := util.GetCtxErrorKey()
	return func(c *gin.Context) {
		// 设置收集错误的collection
		errorList := make([]error, 0, 4)
		c.Set(key, &errorList)

		c.Next()

		value, exists := c.Get(key)
		if exists {
			if errors, ok := value.(*[]error); ok && len(*errors) > 0 {

				buf := bytes.Buffer{}
				for _, err := range *errors {
					buf.WriteString(err.Error())
					buf.WriteByte('\n')
				}

				log.Printf("-------------\n%s-------------", buf.String())
			}
		}
	}
}
