package model

// 商品参数及自定义规格属性
type PmsProductAttributeValue struct {
	ID
	ProductId          uint   `json:"productId"`
	ProductAttributeId uint   `json:"productAttributeId"`
	Value              string `json:"value"`
}

func (a *PmsProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}
