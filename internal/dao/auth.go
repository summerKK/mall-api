package dao

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/model"
)

func (d *Dao) GetItemByName(username string) (*model.UmsAdmin, error) {
	return d.getItemByColumns(map[string]interface{}{"username": username})
}

func (d *Dao) GetItemById(userId int) (*model.UmsAdmin, error) {
	return d.getItemByColumns(map[string]interface{}{"id": userId})
}

func (d *Dao) getItemByColumns(columns map[string]interface{}) (*model.UmsAdmin, error) {
	var user model.UmsAdmin
	var values []interface{}
	queryStr := " "
	for c, v := range columns {
		queryStr += fmt.Sprintf(" %v = ? and", c)
		values = append(values, v)
	}
	queryStr = strings.TrimRight(queryStr, "and")

	err := d.db.Where(queryStr, values...).First(&user).Error
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
