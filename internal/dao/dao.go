package dao

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/summerKK/mall-api/pkg/util"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) GetItemById(userId int, model interface{}) error {
	return d.GetItemByColumns(map[string]interface{}{"id": userId}, model)
}

func (d *Dao) GetItemByColumns(columns map[string]interface{}, model interface{}) error {
	// 判断model 是否是指针类型
	if !util.IsStructPtr(model) {
		panic("model 只能为指针类型的结构体")
	}

	var values []interface{}
	queryStr := " "
	for c, v := range columns {
		queryStr += fmt.Sprintf(" %v = ? and", c)
		values = append(values, v)
	}
	queryStr = strings.TrimRight(queryStr, "and")

	err := d.db.Where(queryStr, values...).First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return err
	}

	return nil
}
