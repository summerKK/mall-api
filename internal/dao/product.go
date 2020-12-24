package dao

import (
	"github.com/summerKK/mall-api/internal/model"
	"gorm.io/gorm"
)

type Product struct {
	*Dao
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		NewDao(db),
	}
}

func (p *Product) DeleteAdditionalAttr(product *model.PmsProduct) {
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

func (p *Product) List(product *model.PmsProduct, pageSize int, pageOffset int) (list []*model.PmsProduct, count int64, err error) {
	name := product.Name
	product.Name = ""
	db := p.db.Where(product)
	if name != "" {
		db.Where("name like ?", name)
	}

	err = db.Model(&product).Count(&count).Limit(pageSize).Offset(pageOffset).Find(&list).Error

	return
}

func (p *Product) SimpleList(product *model.PmsProduct) (list []*model.PmsProduct, err error) {
	name := product.Name
	product.Name = ""
	db := p.db.Where(product)
	if name != "" {
		db.Where("delete_status = ? and name like ?", model.PmsProductDeleteStatusUnDel, name)
	}

	err = db.Find(&list).Error

	return
}

func (p *Product) BatchSetDeleteStatus(ids []uint, deleteStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("delete_status", deleteStatus).Error
}

func (p Product) BatchSetNewStatus(ids []uint, newStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("new_status", newStatus).Error
}

func (p Product) BatchSetPublishStatus(ids []uint, publishStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("publish_status", publishStatus).Error
}

func (p Product) BatchSetRecommendStatus(ids []uint, recommandStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("recommand_status", recommandStatus).Error
}

func (p Product) BatchSetVerifyStatus(ids []uint, verifyStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("verify_status", verifyStatus).Error
}
