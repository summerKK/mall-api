package dao

import (
	"github.com/summerKK/mall-api/internal/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*Dao
}

func NewProduct(db *gorm.DB) *ProductDao {
	return &ProductDao{
		NewDao(db),
	}
}

func (p *ProductDao) DeleteAdditionalAttr(product *model.PmsProduct) {
	// 会员价格
	p.db.Where("product_id", product.Id).Delete(&model.PmsMemberPrice{})
	// 阶梯价格
	p.db.Where("product_id", product.Id).Delete(&model.PmsProductLadder{})
	// 满减价格
	p.db.Where("product_id", product.Id).Delete(&model.PmsProductFullReduction{})
	// 库存
	p.db.Where("product_id", product.Id).Delete(&model.PmsSkuStock{})
	// 商品属性
	p.db.Where("product_id", product.Id).Delete(&model.PmsProductAttributeValue{})
	// 关联专题
	p.db.Where("product_id", product.Id).Delete(&model.CmsSubjectProductRelation{})
	// 关联优选
	p.db.Where("product_id", product.Id).Delete(&model.CmsPrefrenceAreaProductRelation{})
}
