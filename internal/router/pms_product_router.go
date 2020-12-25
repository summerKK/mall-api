package router

import (
	"github.com/gin-gonic/gin"
	"github.com/summerKK/mall-api/internal/router/api/admin"
)

func pmsProductRouter(r *gin.RouterGroup) {
	r1 := r.Group("/product")
	// 创建商品
	r1.POST("/create", admin.PmsProductController.Create)
	// 更新商品
	r1.POST("/update/:id", admin.PmsProductController.Update)
	// 查询商品
	r1.GET("/list", admin.PmsProductController.List)
	// 根据商品名称或货号模糊查询
	r1.GET("/simpleList", admin.PmsProductController.SimpleList)
	// 批量修改删除状态
	r1.POST("/batchUpdate/deleteStatus", admin.PmsProductController.BatchSetDeleteStatus)
	// 批量设为新品
	r1.POST("/batchUpdate/newStatus", admin.PmsProductController.BatchSetNewStatus)
	// 批量上下架
	r1.POST("/batchUpdate/publishStatus", admin.PmsProductController.BatchSetPublishStatus)
	// 批量推荐商品
	r1.POST("/batchUpdate/recommendStatus", admin.PmsProductController.BatchSetRecommendStatus)
	// 批量修改审核状态
	r1.POST("/batchUpdate/verifyStatus", admin.PmsProductController.BatchSetVerifyStatus)
	// 根据商品id获取商品编辑信息
	r1.GET("/updateInfo/:id", admin.PmsProductController.GetUpdateInfo)
}
