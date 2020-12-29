package model

import (
	"time"

	"gorm.io/gorm"
)

type PmsProductVerifyRecord struct {
	ID
	ProductId  uint      `json:"productId"`
	CreateTime LocalTime `json:"createTime"`
	VertifyMan string    `json:"vertifyMan"`
	Status     uint8     `json:"status"`
	Detail     string    `json:"detail"`
}

func (p *PmsProductVerifyRecord) TableName() string {
	return "pms_product_vertify_record"
}

func (p *PmsProductVerifyRecord) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreateTime = LocalTime{time.Now()}

	return nil
}
