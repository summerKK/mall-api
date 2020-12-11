package model

import "github.com/shopspring/decimal"

// 商品阶梯价格设置
type PmsProductLadder struct {
	ID
	ProductId uint            `json:"productId"`
	Count     uint            `json:"count"`
	Discount  decimal.Decimal `json:"discount"`
	Price     decimal.Decimal `json:"price"`
}

func (a *PmsProductLadder) TableName() string {
	return "pms_product_ladder"
}
