package util

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/global"
	businessError "github.com/summerKK/go-code-snippet-library/koel-api/pkg/error"
)

func AddErrorToCtx(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	value, exists := ctx.Get(GetCtxErrorKey())
	if exists {
		// 业务的错误不用加到里面去
		if _, ok := err.(*businessError.BusinessError); ok {
			return
		}

		if errors, ok := value.(*[]error); ok {
			*errors = append(*errors, err)
		}
	}
}

func GetCtxErrorKey() string {
	return global.ServerSetting.ProjectName + "_error_list"
}
