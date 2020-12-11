package model

import "github.com/shopspring/decimal"

// 商品满减价格设置
type PmsProductFullReduction struct {
	ID
	ProductId   uint            `json:"productId"`
	FullPrice   decimal.Decimal `json:"fullPrice"`
	ReducePrice decimal.Decimal `json:"reducePrice"`
}

func (a *PmsProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}
