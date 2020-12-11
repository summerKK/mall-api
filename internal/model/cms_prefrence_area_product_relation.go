package model

type CmsPrefrenceAreaProductRelation struct {
	ID
	PrefrenceAreaId uint `json:"prefrenceAreaId"`
	ProductId       uint `json:"ProductId"`
}
