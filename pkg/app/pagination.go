package app

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/global"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}

	return pageSize
}

// 计算分页偏移量
func GetPageOffset(page, pageSize int) int {
	if page > 0 {
		return (page - 1) * pageSize
	}

	return 0
}
