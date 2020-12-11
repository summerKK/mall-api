package model

import "math/big"

// 商品会员价格设置
type PmsMemberPrice struct {
	ID
	ProductId       uint      `json:"productId"`
	MemberLevelId   uint      `json:"memberLevelId"`
	MemberPrice     big.Float `json:"memberPrice"`
	MemberLevelName string    `json:"memberLevelName"`
}

func (a *PmsMemberPrice) TableName() string {
	return "pms_member_price"
}
