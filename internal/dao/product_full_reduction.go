package dao

import "gorm.io/gorm"

type ProductFullReduction struct {
	*Dao
}

func NewProductFullReduction(db *gorm.DB) *ProductFullReduction {
	return &ProductFullReduction{
		NewDao(db),
	}
}
