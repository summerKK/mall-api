package dao

import "github.com/jinzhu/gorm"

type MemberPriceDao struct {
	*Dao
}

func NewMemberPrice(db *gorm.DB) *MemberPriceDao {
	return &MemberPriceDao{
		NewDao(db),
	}
}
