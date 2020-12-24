package dao

import "gorm.io/gorm"

type MemberPrice struct {
	*Dao
}

func NewMemberPrice(db *gorm.DB) *MemberPrice {
	return &MemberPrice{
		NewDao(db),
	}
}
