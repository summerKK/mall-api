package model

type CmsPrefrenceAreaProductRelation struct {
	ID
	PrefrenceAreaId uint `json:"prefrenceAreaId"`
	ProductId       uint `json:"ProductId"`
}

func (a *CmsPrefrenceAreaProductRelation) TableName() string {
	return "cms_prefrence_area_product_relation"
}
