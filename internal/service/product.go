package service

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/global"
	"github.com/summerKK/mall-api/internal/dao"
	"github.com/summerKK/mall-api/internal/dto/admin"
	"github.com/summerKK/mall-api/internal/model"
	businessError "github.com/summerKK/mall-api/pkg/error"
	"github.com/summerKK/mall-api/pkg/util"
)

type ProductService struct {
	service *service
	dao     *dao.ProductDao
}

func NewProductService(ctx *gin.Context) *ProductService {
	service := NewService(ctx)
	return &ProductService{
		service: service,
		dao:     dao.NewProduct(global.DBEngine),
	}
}

func (s *ProductService) Create(params *admin.ProductRequest) (product *model.PmsProduct, err error) {

	// 开启事物
	db := s.dao.GetDb()
	// 在一个事物里面要使用同一个tx
	tx := db.Begin()
	// 使用事物创建商品
	s.dao.SetDb(tx)

	errorhandler := func(err error) (*model.PmsProduct, error) {
		// 事物回滚
		tx.Rollback()

		util.AddErrorToCtx(s.service.ctx, err)
		return nil, businessError.NewBusinessError("创建失败")
	}

	// 商品创建
	product = params.PmsProduct
	err = s.dao.Insert(product)
	if err != nil {
		return errorhandler(err)
	}

	// 会员价格
	memberPrice := params.MemberPriceList
	if len(memberPrice) > 0 {
		for _, price := range memberPrice {
			price.ProductId = product.Id
		}
		memberPriceDao := dao.NewMemberPrice(tx)
		err = memberPriceDao.Insert(&memberPrice)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 阶梯价格
	productLadderList := params.ProductLadderList
	if len(productLadderList) > 0 {
		for _, ladder := range productLadderList {
			ladder.ProductId = product.Id
		}
		productLadderDao := dao.NewProductLadder(tx)
		err = productLadderDao.Insert(&productLadderList)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 满减价格
	productFullReductionList := params.ProductFullReductionList
	if len(productFullReductionList) > 0 {
		for _, reduction := range productFullReductionList {
			reduction.ProductId = product.Id
		}
		productFullReductionDao := dao.NewProductFullReduction(tx)
		err = productFullReductionDao.Insert(&productFullReductionList)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 添加sku库存
	skuStockList := params.SkuStockList
	if len(skuStockList) > 0 {
		for i, stock := range skuStockList {
			stock.ProductId = product.Id
			// sku 编码
			if stock.SkuCode == "" {
				stock.SetDefaultSkuCode(product.Id, i+1)
			}
		}
		skuStockDao := dao.NewSkuStock(tx)
		err = skuStockDao.Insert(&skuStockList)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 添加商品参数,添加自定义商品规格
	productAttributeValueList := params.ProductAttributeValueList
	if len(productAttributeValueList) > 0 {
		for _, value := range productAttributeValueList {
			value.ProductId = product.Id
		}
		productAttributeValueDao := dao.NewProductAttributeValue(tx)
		err = productAttributeValueDao.Insert(&productAttributeValueList)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 关联专题
	subjectProductRelationList := params.SubjectProductRelationList
	if len(subjectProductRelationList) > 0 {
		for _, relation := range subjectProductRelationList {
			relation.ProductId = product.Id
		}
		subjectProductRelationDao := dao.NewSubjectProductRelation(tx)
		err = subjectProductRelationDao.Insert(&subjectProductRelationList)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 关联优选
	prefrenceAreaProductRelationList := params.PrefrenceAreaProductRelationList
	if len(prefrenceAreaProductRelationList) > 0 {
		for _, relation := range prefrenceAreaProductRelationList {
			relation.ProductId = product.Id
		}
		prefrenceAreaProductRelationDao := dao.NewPrefrenceAreaProductRelation(tx)
		err = prefrenceAreaProductRelationDao.Insert(&prefrenceAreaProductRelationList)
		if err != nil {
			return errorhandler(err)
		}
	}

	// 提交事物
	tx.Commit()

	return
}

func (s *ProductService) Update(params *admin.ProductRequest, productId int) (product *model.PmsProduct, err error) {
	return
}
