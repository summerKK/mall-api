package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

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

// 设置默认 skuCode
func (a *PmsSkuStock) SetDefaultSkuCode(productId uint, i int) {
	date := time.Now().Format("2006-01-02 15:04:05")
	var buf strings.Builder
	buf.WriteString(date)
	// 四位商品id
	buf.WriteString(fmt.Sprintf("%04d", productId))
	// 三位索引id
	buf.WriteString(fmt.Sprintf("%03d", i))
	a.SkuCode = buf.String()
}
