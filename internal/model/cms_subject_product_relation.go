package model

// 专题和商品关系
type CmsSubjectProductRelation struct {
	ID
	SubjectId uint `json:"subjectId"`
	ProductId uint `json:"productId"`
}

func (a *CmsSubjectProductRelation) TableName() string {
	return "cms_subject_product_relation"
}
