package dao

import "gorm.io/gorm"

type SkuStock struct {
	*Dao
}

func NewSkuStock(db *gorm.DB) *SkuStock {
	return &SkuStock{
		NewDao(db),
	}
}
