package dao

import "github.com/jinzhu/gorm"

type ProductLadderDao struct {
	*Dao
}

func NewProductLadder(db *gorm.DB) *ProductLadderDao {
	return &ProductLadderDao{
		NewDao(db),
	}
}
