package dao

import (
	"gorm.io/gorm"
)

type ProductDao struct {
	*Dao
}

func NewProduct(db *gorm.DB) *ProductDao {
	return &ProductDao{
		NewDao(db),
	}
}
