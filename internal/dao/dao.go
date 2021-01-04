package dao

import (
	"fmt"
	"strings"

	"github.com/summerKK/mall-api/pkg/util"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) GetItemById(id uint, model interface{}) (exists bool, err error) {
	return d.GetItemByColumns(map[string]interface{}{"id": id}, model)
}

func (d *Dao) GetItemByColumns(columns map[string]interface{}, model interface{}) (exists bool, err error) {
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

	err = d.db.Where(queryStr, values...).First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}

		return
	}
	exists = true

	return
}

func (d *Dao) Insert(model interface{}) error {
	// 判断model 是否是指针类型 / 或者 []*ptr
	if !util.IsStructPtr(model) && !util.IsSliceElemPtr(model) {
		panic("model 只能为指针类型的结构体")
	}

	return d.db.Create(model).Error
}

func (d *Dao) Save(model interface{}) error {
	if !util.IsStructPtr(model) {
		panic("model 只能为指针类型的结构体")
	}

	return d.db.Save(model).Error
}

func (d *Dao) GetDb() *gorm.DB {
	return d.db
}

func (d *Dao) SetDb(db *gorm.DB) {
	d.db = db
}
