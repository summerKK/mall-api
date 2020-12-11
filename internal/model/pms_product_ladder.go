package model

import "math/big"

// 商品阶梯价格设置
type PmsProductLadder struct {
	ID
	ProductId uint      `json:"productId"`
	Count     uint      `json:"count"`
	Discount  big.Float `json:"discount"`
	Price     big.Float `json:"price"`
}

func (a *PmsProductLadder) TableName() string {
	return "pms_product_ladder"
}
