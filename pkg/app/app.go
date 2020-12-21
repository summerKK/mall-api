package app

import (
	"math"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/pkg/error"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}, err *error.Error) {
	h := gin.H{
		"code":    err.Code(),
		"message": err.Msg(),
		"data":    data,
	}
	if len(err.Details()) > 0 {
		h["error_details"] = err.Details()
	}

	r.Ctx.JSON(err.HttpCode(), h)
}

// 列表返回
func (r *Response) Success(data interface{}) {
	r.ToResponse(data, error.Success)
}

func (r *Response) Fail(err *error.Error) {
	r.ToResponse(nil, err)
}

func (r *Response) SuccessWithPage(data interface{}, count, pageNum, pageSize int) {
	h := gin.H{
		"list":      data,
		"pageNum":   pageNum,
		"pageSize":  pageSize,
		"total":     count,
		"totalPage": math.Ceil(float64(count / pageSize)),
	}

	r.Success(h)
}
