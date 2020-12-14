package dao

import "gorm.io/gorm"

type ProductLadderDao struct {
	*Dao
}

func NewProductLadder(db *gorm.DB) *ProductLadderDao {
	return &ProductLadderDao{
		NewDao(db),
	}
}
