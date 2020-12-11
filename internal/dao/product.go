package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/summerKK/mall-api/internal/model"
)

type ProductDao struct {
	*Dao
}

func NewProduct(db *gorm.DB) *ProductDao {
	return &ProductDao{
		NewDao(db),
	}
}

func (p *ProductDao) Insert(product *model.PmsProduct) error {
	return p.db.Create(product).Error
}
