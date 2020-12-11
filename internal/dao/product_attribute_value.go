package dao

import "github.com/jinzhu/gorm"

type ProductAttributeValueDao struct {
	*Dao
}

func NewProductAttributeValue(db *gorm.DB) *ProductAttributeValueDao {
	return &ProductAttributeValueDao{
		NewDao(db),
	}
}
