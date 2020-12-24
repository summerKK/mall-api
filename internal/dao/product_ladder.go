package dao

import "gorm.io/gorm"

type ProductLadder struct {
	*Dao
}

func NewProductLadder(db *gorm.DB) *ProductLadder {
	return &ProductLadder{
		NewDao(db),
	}
}
