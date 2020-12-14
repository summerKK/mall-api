package dao

import "gorm.io/gorm"

type MemberPriceDao struct {
	*Dao
}

func NewMemberPrice(db *gorm.DB) *MemberPriceDao {
	return &MemberPriceDao{
		NewDao(db),
	}
}
