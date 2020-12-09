package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}
