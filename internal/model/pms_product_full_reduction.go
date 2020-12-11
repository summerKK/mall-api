package model

import "math/big"

// 商品满减价格设置
type PmsProductFullReduction struct {
	ID
	ProductId   uint      `json:"productId"`
	FullPrice   big.Float `json:"fullPrice"`
	ReducePrice big.Float `json:"reducePrice"`
}

func (a *PmsProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}
