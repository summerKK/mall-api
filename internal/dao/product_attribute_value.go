package dao

import "gorm.io/gorm"

type ProductAttributeValueDao struct {
	*Dao
}

func NewProductAttributeValue(db *gorm.DB) *ProductAttributeValueDao {
	return &ProductAttributeValueDao{
		NewDao(db),
	}
}
