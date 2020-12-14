package dao

import "gorm.io/gorm"

type PrefrenceAreaProductRelation struct {
	*Dao
}

func NewPrefrenceAreaProductRelation(db *gorm.DB) *PrefrenceAreaProductRelation {
	return &PrefrenceAreaProductRelation{
		NewDao(db),
	}
}
