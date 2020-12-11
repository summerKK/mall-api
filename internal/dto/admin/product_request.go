package admin

import "github.com/summerKK/mall-api/internal/model"

type ProductRequest struct {
	// 商品详情
	model.PmsProduct
	// 商品阶梯价格设置
	ProductLadderList []model.PmsProductLadder `form:"productLadderList" json:"productLadderList" binding:"required"`
	// 商品满减价格设置
	ProductFullReductionList []model.PmsProductFullReduction `form:"productFullReductionList" json:"productFullReductionList" binding:"required"`
	// 商品会员价格设置
	MemberPriceList []model.PmsMemberPrice `form:"memberPriceList" json:"memberPriceList" binding:"required"`
	// 商品的sku库存信息
	SkuStockList []model.PmsSkuStock `form:"skuStockList" json:"skuStockList" binding:"required"`
	// 商品参数及自定义规格属性
	ProductAttributeValueList []model.PmsProductAttributeValue `form:"productAttributeValueList" json:"productAttributeValueList" binding:"required"`
	// 专题和商品关系
	SubjectProductRelationList []model.CmsSubjectProductRelation `form:"subjectProductRelationList" json:"subjectProductRelationList" binding:"required"`
	// 优选专区和商品的关系
	PrefrenceAreaProductRelationList []model.CmsPrefrenceAreaProductRelation `form:"prefrenceAreaProductRelationList" json:"prefrenceAreaProductRelationList" binding:"required"`
}
