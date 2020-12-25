package model

import (
	"github.com/shopspring/decimal"
)

const (
	PmsProductDeleteStatusUnDel = 0
	PmsProductDeleteStatusDel   = 1
)

type PmsProductAdditional struct {
	// 商品描述
	Description string `json:"description"`
	DetailDesc  string `json:"detailDesc"`
	// 产品详情网页内容
	DetailHtml string `json:"detailHtml"`
	// 移动端网页详情
	DetailMobileHtml string `json:"detailMobileHtml"`
}

type PmsProductWithRelation struct {
	*PmsProduct
	ProductLadderList                []*PmsProductLadder                `json:"productLadderList" gorm:"foreignKey:ProductId"`
	ProductFullReductionList         []*PmsProductFullReduction         `json:"productFullReductionList" gorm:"foreignKey:ProductId"`
	MemberPriceList                  []*PmsMemberPrice                  `json:"memberPriceList" gorm:"foreignKey:ProductId"`
	SkuStockList                     []*PmsSkuStock                     `json:"skuStockList" gorm:"foreignKey:ProductId"`
	ProductAttributeValueList        []*PmsProductAttributeValue        `json:"productAttributeValueList" gorm:"foreignKey:ProductId"`
	SubjectProductRelationList       []*CmsSubjectProductRelation       `json:"subjectProductRelationList" gorm:"foreignKey:ProductId"`
	PrefrenceAreaProductRelationList []*CmsPrefrenceAreaProductRelation `json:"prefrenceAreaProductRelationList" gorm:"foreignKey:ProductId"`
}

type PmsProduct struct {
	ID
	BrandId                    uint   `form:"brandId" json:"brandId" binding:"required"`
	ProductCategoryId          uint   `form:"productCategoryId" json:"productCategoryId" binding:"required"`
	FeightTemplateId           uint   `json:"feightTemplateId"`
	ProductAttributeCategoryId uint   `json:"productAttributeCategoryId"`
	Name                       string `form:"name" json:"name" binding:"required"`
	Pic                        string `json:"pic"`
	// 货号
	ProductSn string `json:"productSn"`
	// 删除状态：0->未删除；1->已删除
	DeleteStatus uint8 `json:"deleteStatus"`
	// 上架状态：0->下架；1->上架"
	PublishStatus uint8 `json:"publishStatus"`
	// 新品状态:0->不是新品；1->新品
	NewStatus uint8 `json:"newStatus"`
	// 推荐状态；0->不推荐；1->推荐
	RecommandStatus uint8 `json:"recommandStatus"`
	// 审核状态：0->未审核；1->审核通过
	VerifyStatus uint8 `json:"verifyStatus"`
	// 排序
	Sort int `json:"sort"`
	// 销量
	Sale uint `json:"sale"`
	// 价格
	Price decimal.Decimal `json:"price"`
	// 促销价格
	PromotionPrice decimal.Decimal `json:"promotionPrice"`
	// 赠送的成长值
	GiftGrowth uint `json:"giftGrowth"`
	// 赠送的积分
	GiftPoint uint `json:"giftPoint"`
	// 限制使用的积分数
	UsePointLimit uint `json:"usePointLimit"`
	// 副标题
	SubTitle string `json:"subTitle"`
	// 市场价
	OriginalPrice decimal.Decimal `json:"originalPrice"`
	// 库存
	Stock uint `json:"stock"`
	// 库存预警值
	LowStock uint `json:"lowStock"`
	// 单位
	Unit string `json:"unit"`
	// 商品重量，默认为克
	Weight decimal.Decimal `json:"weight"`
	// 是否为预告商品：0->不是；1->是
	PreviewStatus uint8 `json:"previewStatus"`
	// 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	ServiceIds string `json:"serviceIds"`
	// 关键字
	Keywords string `json:"keywords"`
	Note     string `json:"note"`
	// 画册图片，连产品图片限制为5张，以逗号分割
	AlbumPics   string `json:"albumPics"`
	DetailTitle string `json:"detailTitle"`
	// 促销开始时间
	PromotionStartTime LocalTime `json:"promotionStartTime"`
	// 促销结束时间
	PromotionEndTime LocalTime `json:"promotionEndTime"`
	// 活动限购数量
	PromotionPerLimit uint `json:"promotionPerLimit"`
	// 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	PromotionType uint8
	// 品牌名称
	BrandName string
	// 商品分类名称
	ProductCategoryName string `json:"productCategoryName"`

	// 附加属性
	PmsProductAdditional
}

func (a *PmsProduct) TableName() string {
	return "pms_product"
}
