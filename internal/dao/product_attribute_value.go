package dao

import "gorm.io/gorm"

type ProductAttributeValue struct {
	*Dao
}

func NewProductAttributeValue(db *gorm.DB) *ProductAttributeValue {
	return &ProductAttributeValue{
		NewDao(db),
	}
}
