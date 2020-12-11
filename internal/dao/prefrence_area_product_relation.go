package dao

import "github.com/jinzhu/gorm"

type PrefrenceAreaProductRelation struct {
	*Dao
}

func NewPrefrenceAreaProductRelation(db *gorm.DB) *PrefrenceAreaProductRelation {
	return &PrefrenceAreaProductRelation{
		NewDao(db),
	}
}
