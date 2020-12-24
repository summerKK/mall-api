package model

type PmsProductVerifyRecord struct {
	ID
	ProductId  uint      `json:"productId"`
	CreateTime LocalTime `json:"createTime"`
	VertifyMan string    `json:"vertifyMan"`
	Status     uint8     `json:"status"`
	Detail     string    `json:"detail"`
}

func (a *PmsProductVerifyRecord) TableName() string {
	return "pms_product_vertify_record"
}
