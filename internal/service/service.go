package service

import (
	"github.com/gin-gonic/gin"
)

type service struct {
	ctx *gin.Context
}

func NewService(ctx *gin.Context) *service {
	return &service{
		ctx: ctx,
	}
}
