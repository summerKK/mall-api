package model

import "github.com/shopspring/decimal"

// 商品的sku库存信息
type PmsSkuStock struct {
	ID
	ProductId      uint            `json:"productId"`
	SkuCode        string          `json:"skuCode"`
	Price          decimal.Decimal `json:"price"`
	Stock          uint            `json:"stock"`
	LowStock       uint            `json:"lowStock"`
	Pic            string          `json:"pic"`
	Sale           uint            `json:"sale"`
	PromotionPrice decimal.Decimal `json:"promotionPrice"`
	LockStock      uint            `json:"LockStock"`
	SpData         string          `json:"spData"`
}

func (a *PmsSkuStock) TableName() string {
	return "pms_sku_stock"
}
