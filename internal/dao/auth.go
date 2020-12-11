package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/summerKK/mall-api/internal/model"
)

type AuthDao struct {
	*Dao
}

func NewAuth(db *gorm.DB) *AuthDao {
	return &AuthDao{
		NewDao(db),
	}
}

func (a *AuthDao) Register(user *model.UmsAdmin) error {
	err := a.db.Create(user).Error
	return err
}

func (a *AuthDao) GetItemByName(username string) (*model.UmsAdmin, error) {
	user := &model.UmsAdmin{}
	err := a.GetItemByColumns(map[string]interface{}{"username": username}, user)

	return user, err
}

func (a *AuthDao) DeleteItemById(userId int) error {
	return a.db.Delete(&model.UmsAdmin{}, userId).Error
}
