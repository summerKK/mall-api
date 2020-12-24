package admin

import "github.com/summerKK/mall-api/internal/model"

type ProductRequest struct {
	// 商品详情
	*model.PmsProduct
	// 商品阶梯价格设置
	ProductLadderList []*model.PmsProductLadder `form:"productLadderList" json:"productLadderList" binding:"required"`
	// 商品满减价格设置
	ProductFullReductionList []*model.PmsProductFullReduction `form:"productFullReductionList" json:"productFullReductionList" binding:"required"`
	// 商品会员价格设置
	MemberPriceList []*model.PmsMemberPrice `form:"memberPriceList" json:"memberPriceList" binding:"required"`
	// 商品的sku库存信息
	SkuStockList []*model.PmsSkuStock `form:"skuStockList" json:"skuStockList" binding:"required"`
	// 商品参数及自定义规格属性
	ProductAttributeValueList []*model.PmsProductAttributeValue `form:"productAttributeValueList" json:"productAttributeValueList" binding:"required"`
	// 专题和商品关系
	SubjectProductRelationList []*model.CmsSubjectProductRelation `form:"subjectProductRelationList" json:"subjectProductRelationList" binding:"required"`
	// 优选专区和商品的关系
	PrefrenceAreaProductRelationList []*model.CmsPrefrenceAreaProductRelation `form:"prefrenceAreaProductRelationList" json:"prefrenceAreaProductRelationList" binding:"required"`
}

type ProductListRequest struct {
	// 上架状态
	PublishStatus uint8 `form:"publishStatus" json:"publishStatus"`
	// 审核状态
	VerifyStatus uint8 `form:"verifyStatus" json:"verifyStatus"`
	// 商品名称模糊关键字
	Keyword string `form:"keyword" json:"keyword"`
	// 商品货号
	ProductSn string `form:"productSn" json:"productSn"`
	// 商品分类编号
	ProductCategoryId uint `form:"productCategoryId" json:"productCategoryId"`
	// 商品品牌编号
	BrandId uint `form:"brandId" json:"brandId"`
}

type ProductBatchDeleteStatusRequest struct {
	Ids          []uint `form:"ids" binding:"required"`
	DeleteStatus uint8  `form:"deleteStatus" binding:"required,oneof=0 1"`
}

type ProductBatchSetNewStatusRequest struct {
	Ids       []uint `form:"ids" binding:"required"`
	NewStatus uint8  `form:"newStatus" binding:"required,oneof=0 1"`
}

type ProductBatchSetPublishStatusRequest struct {
	Ids           []uint `form:"ids" binding:"required"`
	PublishStatus uint8  `form:"publishStatus" binding:"required,oneof=0 1"`
}

type ProductBatchSetRecommendStatusRequest struct {
	Ids             []uint `form:"ids" binding:"required"`
	RecommendStatus uint8  `form:"recommendStatus" binding:"required,oneof=0 1"`
}

type ProductBatchSetVerifyStatusRequest struct {
	Ids          []uint `form:"ids" binding:"required"`
	VerifyStatus uint8  `form:"verifyStatus" binding:"required,oneof=0 1"`
}
