package dao

import "gorm.io/gorm"

type ProductFullReductionDao struct {
	*Dao
}

func NewProductFullReduction(db *gorm.DB) *ProductFullReductionDao {
	return &ProductFullReductionDao{
		NewDao(db),
	}
}
