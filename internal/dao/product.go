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

func (p *ProductDao) List(product *model.PmsProduct, pageSize int, pageOffset int) (list []*model.PmsProduct, count int64, err error) {
	name := product.Name
	product.Name = ""
	db := p.db.Where(product)
	if name != "" {
		db.Where("name like ?", name)
	}

	err = db.Model(&product).Count(&count).Limit(pageSize).Offset(pageOffset).Find(&list).Error

	return
}

func (p *ProductDao) SimpleList(product *model.PmsProduct) (list []*model.PmsProduct, err error) {
	name := product.Name
	product.Name = ""
	db := p.db.Where(product)
	if name != "" {
		db.Where("delete_status = ? and name like ?", model.PmsProductDeleteStatusUnDel, name)
	}

	err = db.Find(&list).Error

	return
}

func (p *ProductDao) BatchSetDeleteStatus(ids []uint, deleteStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("delete_status", deleteStatus).Error
}

func (p ProductDao) BatchSetNewStatus(ids []uint, newStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("new_status", newStatus).Error
}

func (p ProductDao) BatchSetPublishStatus(ids []uint, publishStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("publish_status", publishStatus).Error
}

func (p ProductDao) BatchSetRecommendStatus(ids []uint, recommandStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("recommand_status", recommandStatus).Error
}

func (p ProductDao) BatchSetVerifyStatus(ids []uint, verifyStatus uint8) error {
	return p.db.Model(&model.PmsProduct{}).Where("id in (?)", ids).Update("verify_status", verifyStatus).Error
}
