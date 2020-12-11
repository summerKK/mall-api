package dao

import "github.com/jinzhu/gorm"

type SkuStockDao struct {
	*Dao
}

func NewSkuStock(db *gorm.DB) *SkuStockDao {
	return &SkuStockDao{
		NewDao(db),
	}
}
