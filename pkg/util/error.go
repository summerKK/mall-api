package util

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/global"
	businessError "github.com/summerKK/mall-api/pkg/error"
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

		if errList, ok := value.(*[]error); ok {
			stack := Stack(2)
			err = errors.New(err.Error() + "\n\n" + stack)
			*errList = append(*errList, err)
		}
	}
}

func GetCtxErrorKey() string {
	return global.ServerSetting.ProjectName + "_error_list"
}
