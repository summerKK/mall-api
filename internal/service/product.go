package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/global"
	"github.com/summerKK/mall-api/internal/dao"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/model"
	businessError "github.com/summerKK/mall-api/pkg/error"
	"github.com/summerKK/mall-api/pkg/util"
	"gorm.io/gorm"
)

type ProductService struct {
	service   *service
	dao       *dao.ProductDao
	productId uint
	// 操作事物对应的tx
	tx *gorm.DB
}

func NewProductService(ctx *gin.Context) *ProductService {
	service := NewService(ctx)
	return &ProductService{
		service: service,
		dao:     dao.NewProduct(global.DBEngine),
	}
}

func (s *ProductService) errorHandler(err error) (*model.PmsProduct, error) {
	// 事物回滚
	s.tx.Rollback()

	util.AddErrorToCtx(s.service.ctx, err)
	return nil, businessError.NewBusinessError("操作失败")
}

// 通过事务提交
func (s *ProductService) beginTransaction() *gorm.DB {
	// 开启事物
	db := s.dao.GetDb()
	// 在一个事物里面要使用同一个tx
	tx := db.Begin()
	// 使用事物创建商品
	s.dao.SetDb(tx)
	s.tx = tx

	return tx
}

func (s *ProductService) Create(params *admin.ProductRequest) (product *model.PmsProduct, err error) {

	tx := s.beginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 商品创建
	product = params.PmsProduct
	err = s.dao.Insert(product)
	if err != nil {
		return s.errorHandler(err)
	}

	// 商品Id
	s.productId = product.Id
	// 添加商品额外属性
	err = s.saveAdditionalAttr(params)
	if err != nil {
		return s.errorHandler(err)
	}

	// 提交事物
	tx.Commit()

	return
}

func (s *ProductService) Update(params *admin.ProductRequest, productId int) (product *model.PmsProduct, err error) {

	// 查看商品是否存在
	exists, err := s.dao.GetItemById(productId, &model.PmsProduct{})
	if err != nil {
		return nil, businessError.NewBusinessError("操作失败")
	}
	if !exists {
		return nil, businessError.NewBusinessError("商品不存在")
	}

	tx := s.beginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	params.PmsProduct.Id = uint(productId)
	s.productId = params.PmsProduct.Id

	err = s.dao.Save(params.PmsProduct)
	if err != nil {
		return s.errorHandler(err)
	}

	s.dao.DeleteAdditionalAttr(params.PmsProduct)
	err = s.saveAdditionalAttr(params)
	if err != nil {
		return s.errorHandler(err)
	}

	tx.Commit()

	product = params.PmsProduct

	return
}

func (s *ProductService) List(params *admin.ProductListRequest, pageSize int, pageOffset int) (list []*model.PmsProduct, count int64, err error) {
	defer func() {
		util.AddErrorToCtx(s.service.ctx, err)
	}()

	product := &model.PmsProduct{
		PublishStatus:     params.PublishStatus,
		VerifyStatus:      params.VerifyStatus,
		ProductSn:         params.ProductSn,
		BrandId:           params.BrandId,
		ProductCategoryId: params.ProductCategoryId,
	}
	if params.Keyword != "" {
		product.Name = fmt.Sprintf("%%%s%%", params.Keyword)
	}

	list, count, err = s.dao.List(product, pageSize, pageOffset)

	return
}

func (s *ProductService) SimpleList(keyword string) (list []*model.PmsProduct, err error) {
	defer func() {
		util.AddErrorToCtx(s.service.ctx, err)
	}()

	product := &model.PmsProduct{}
	if keyword != "" {
		product.Name = fmt.Sprintf("%%%s%%", keyword)
	}

	list, err = s.dao.SimpleList(product)

	return
}

func (s *ProductService) saveAdditionalAttr(params *admin.ProductRequest) (err error) {
	// 会员价格
	memberPrice := params.MemberPriceList
	err = s.saveMemberPrice(memberPrice)
	if err != nil {
		return
	}

	// 阶梯价格
	productLadderList := params.ProductLadderList
	err = s.saveProductLadder(productLadderList)
	if err != nil {
		return
	}

	// 满减价格
	productFullReductionList := params.ProductFullReductionList
	err = s.saveProductFullReduction(productFullReductionList)
	if err != nil {
		return
	}

	// 添加sku库存
	skuStockList := params.SkuStockList
	err = s.saveSkuStock(skuStockList)
	if err != nil {
		return
	}

	// 添加商品参数,添加自定义商品规格
	productAttributeValueList := params.ProductAttributeValueList
	err = s.saveProductAttributeValue(productAttributeValueList)
	if err != nil {
		return
	}

	// 关联专题
	subjectProductRelationList := params.SubjectProductRelationList
	err = s.saveSubjectProductRelation(subjectProductRelationList)
	if err != nil {
		return
	}

	// 关联优选
	prefrenceAreaProductRelationList := params.PrefrenceAreaProductRelationList
	err = s.savePrefrenceAreaProductRelation(prefrenceAreaProductRelationList)
	if err != nil {
		return
	}

	return
}

// 会员价格
func (s *ProductService) saveMemberPrice(memberPrice []*model.PmsMemberPrice) error {
	if len(memberPrice) > 0 {
		for _, price := range memberPrice {
			price.ProductId = s.productId
			price.Id = 0
		}
		memberPriceDao := dao.NewMemberPrice(s.tx)

		return memberPriceDao.Insert(&memberPrice)
	}

	return nil
}

// 阶梯价格
func (s *ProductService) saveProductLadder(productLadderList []*model.PmsProductLadder) error {
	if len(productLadderList) > 0 {
		for _, ladder := range productLadderList {
			ladder.ProductId = s.productId
			ladder.Id = 0
		}
		productLadderDao := dao.NewProductLadder(s.tx)

		return productLadderDao.Insert(&productLadderList)
	}

	return nil
}

// 满减价格
func (s *ProductService) saveProductFullReduction(productFullReductionList []*model.PmsProductFullReduction) error {
	if len(productFullReductionList) > 0 {
		for _, reduction := range productFullReductionList {
			reduction.ProductId = s.productId
			reduction.Id = 0
		}
		productFullReductionDao := dao.NewProductFullReduction(s.tx)

		return productFullReductionDao.Insert(&productFullReductionList)
	}

	return nil
}

// 添加sku库存
func (s *ProductService) saveSkuStock(skuStockList []*model.PmsSkuStock) error {
	if len(skuStockList) > 0 {
		for i, stock := range skuStockList {
			stock.ProductId = s.productId
			stock.Id = 0
			// sku 编码
			if stock.SkuCode == "" {
				stock.SetDefaultSkuCode(s.productId, i+1)
			}
		}
		skuStockDao := dao.NewSkuStock(s.tx)

		return skuStockDao.Insert(&skuStockList)
	}

	return nil
}

// 添加商品参数,添加自定义商品规格
func (s *ProductService) saveProductAttributeValue(productAttributeValueList []*model.PmsProductAttributeValue) error {
	if len(productAttributeValueList) > 0 {
		for _, value := range productAttributeValueList {
			value.ProductId = s.productId
			value.Id = 0
		}
		productAttributeValueDao := dao.NewProductAttributeValue(s.tx)

		return productAttributeValueDao.Insert(&productAttributeValueList)
	}

	return nil
}

// 关联专题
func (s *ProductService) saveSubjectProductRelation(subjectProductRelationList []*model.CmsSubjectProductRelation) error {
	if len(subjectProductRelationList) > 0 {
		for _, relation := range subjectProductRelationList {
			relation.ProductId = s.productId
			relation.Id = 0
		}
		subjectProductRelationDao := dao.NewSubjectProductRelation(s.tx)

		return subjectProductRelationDao.Insert(&subjectProductRelationList)
	}

	return nil
}

// 关联优选
func (s *ProductService) savePrefrenceAreaProductRelation(prefrenceAreaProductRelationList []*model.CmsPrefrenceAreaProductRelation) error {
	if len(prefrenceAreaProductRelationList) > 0 {
		for _, relation := range prefrenceAreaProductRelationList {
			relation.ProductId = s.productId
			relation.Id = 0
		}
		prefrenceAreaProductRelationDao := dao.NewPrefrenceAreaProductRelation(s.tx)

		return prefrenceAreaProductRelationDao.Insert(&prefrenceAreaProductRelationList)
	}

	return nil
}
