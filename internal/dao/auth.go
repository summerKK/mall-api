package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/model"
)

func (d *Dao) GetUserByName(username string) (*model.UmsAdmin, error) {
	var user model.UmsAdmin
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (d *Dao) Register(user *model.UmsAdmin) error {
	err := d.db.Create(user).Error
	return err
}
