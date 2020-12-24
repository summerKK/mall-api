package dao

import "gorm.io/gorm"

type ProductVerifyRecordDao struct {
	*Dao
}

func NewProductVerifyRecord(db *gorm.DB) *ProductVerifyRecordDao {
	return &ProductVerifyRecordDao{
		NewDao(db),
	}
}
