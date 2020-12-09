package error

import "net/http"

var (
	Success                   = NewError(200, "成功")
	UnauthorizedAuthNotExist  = NewError(401, "鉴权失败,找不到对应的账号或者密码错误").WithHttpCode(http.StatusUnauthorized)
	UnauthorizedTokenError    = NewError(401, "鉴权失败,Token错误").WithHttpCode(http.StatusUnauthorized)
	UnauthorizedTokenTimeout  = NewError(401, "鉴权失败,Token超时").WithHttpCode(http.StatusUnauthorized)
	UnauthorizedTokenGenerate = NewError(401, "鉴权失败,Token生成失败").WithHttpCode(http.StatusUnauthorized)
	InvalidParams             = NewError(400, "入参错误")
)
