package dao

import "gorm.io/gorm"

type SkuStockDao struct {
	*Dao
}

func NewSkuStock(db *gorm.DB) *SkuStockDao {
	return &SkuStockDao{
		NewDao(db),
	}
}
