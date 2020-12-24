package dao

import (
	"github.com/summerKK/mall-api/internal/model"
	"gorm.io/gorm"
)

type Auth struct {
	*Dao
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		NewDao(db),
	}
}

func (a *Auth) Register(user *model.UmsAdmin) error {
	err := a.db.Create(user).Error
	return err
}

func (a *Auth) GetItemByName(username string) (*model.UmsAdmin, error) {
	user := &model.UmsAdmin{}
	_, err := a.GetItemByColumns(map[string]interface{}{"username": username}, user)

	return user, err
}

func (a *Auth) DeleteItemById(userId int) error {
	return a.db.Delete(&model.UmsAdmin{}, userId).Error
}
