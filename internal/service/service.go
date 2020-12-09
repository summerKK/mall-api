package service

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/global"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/dao"
)

type service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func NewService(ctx *gin.Context) *service {
	return &service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
