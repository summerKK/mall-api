package util

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/model"
)

func GetUser(c *gin.Context) (*model.UmsAdmin, error) {
	value, exists := c.Get("userInfo")
	if exists {
		if admin, ok := value.(*model.UmsAdmin); ok {
			return admin, nil
		}
	}

	return nil, errors.New("user not exist")
}

func MustGetUser(c *gin.Context) *model.UmsAdmin {
	user, err := GetUser(c)
	if err != nil {
		return &model.UmsAdmin{}
	}

	return user
}
