package model

import (
	"github.com/shopspring/decimal"
)

// 商品会员价格设置
type PmsMemberPrice struct {
	ID
	ProductId       uint            `json:"productId"`
	MemberLevelId   uint            `json:"memberLevelId"`
	MemberPrice     decimal.Decimal `json:"memberPrice"`
	MemberLevelName string          `json:"memberLevelName"`
}

func (a *PmsMemberPrice) TableName() string {
	return "pms_member_price"
}
