package dao

import "github.com/jinzhu/gorm"

type ProductFullReductionDao struct {
	*Dao
}

func NewProductFullReduction(db *gorm.DB) *ProductFullReductionDao {
	return &ProductFullReductionDao{
		NewDao(db),
	}
}
