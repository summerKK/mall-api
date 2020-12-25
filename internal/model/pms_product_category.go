package model

type PmsProductCategory struct {
	ID
	// 上机分类的编号：0表示一级分类
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
	// 分类级别：0->1级；1->2级
	Level        uint   `json:"level"`
	ProductCount uint   `json:"productCount"`
	ProductUnit  string `json:"productUnit"`
	// 是否显示在导航栏：0->不显示；1->显示
	NavStatus uint8 `json:"navStatus"`
	// 显示状态：0->不显示；1->显示
	ShowStatus uint8 `json:"showStatus"`
	// 排序
	Sort uint `json:"sort"`
	// 图标
	Icon     string `json:"icon"`
	Keywords string `json:"keywords"`
	// 描述
	Description string `json:"description"`
}

func (a *PmsProductCategory) TableName() string {
	return "pms_product_category"
}
