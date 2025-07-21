package coupon

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	couponReq "github.com/flipped-aurora/gin-vue-admin/server/model/coupon/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CouponApi struct{}

func fillField(cp *coupon.Coupon) (*coupon.Coupon, error) {
	n := len(cp.Period)
	switch n {
	case 1:
		cp.ValidDays = int(cp.Period[0])
		cp.ValidStart = 0
		cp.ValidEnd = 0
	case 2:
		cp.ValidStart = 0
		cp.ValidEnd = 0
		cp.ValidStart = cp.Period[0]
		cp.ValidEnd = cp.Period[1]
		cp.ValidDays = 0

	default:
		return cp, nil
	}
	return cp, nil
}

// CreateCoupon 创建coupon表
// @Tags Coupon
// @Summary 创建coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coupon.Coupon true "创建coupon表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /coupon/create [post]
func (cpApi *CouponApi) CreateCoupon(c *gin.Context) {
	var cp = new(coupon.Coupon)
	err := c.ShouldBindJSON(cp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cp, err = fillField(cp)
	if err != nil {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CreationFailed, response.ERROR), c)
		return
	}
	err = cpService.CreateCoupon(cp)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CreationFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// DeleteCoupon 删除coupon表
// @Tags Coupon
// @Summary 删除coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coupon.Coupon true "删除coupon表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /coupon/delete [delete]
func (cpApi *CouponApi) DeleteCoupon(c *gin.Context) {
	id := c.Query("id")
	err := cpService.DeleteCoupon(id)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.DeleteFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// UpdateCoupon 更新coupon表
// @Tags Coupon
// @Summary 更新coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coupon.Coupon true "更新coupon表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /coupon/update [put]
func (cpApi *CouponApi) UpdateCoupon(c *gin.Context) {
	var cp = new(coupon.Coupon)
	err := c.ShouldBindJSON(cp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cp, err = fillField(cp)
	if err != nil {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UpdateFailed, response.ERROR), c)
	}
	err = cpService.UpdateCoupon(*cp)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UpdateFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// FindCoupon 用id查询coupon表
// @Tags Coupon
// @Summary 用id查询coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "用id查询coupon表"
// @Success 200 {object} response.Response{data=coupon.Coupon,msg=object} "查询成功"
// @Router /coupon/find [get]
func (cpApi *CouponApi) FindCoupon(c *gin.Context) {
	id := c.Query("id")
	recp, err := cpService.GetCoupon(id)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithData(recp, c)
}

// Options 获取coupon列表
// @Tags Coupon
// @Summary 获取coupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponSearch true "下拉列表获取优惠券"
// @Success 200 {object} response.Response{data=response.NormalResult,msg=object} "获取成功"
// @Router /coupon/options [Get]
func (cpApi *CouponApi) Options(c *gin.Context) {
	var pageInfo couponReq.CouponSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.Page = 1
	pageInfo.PageSize = 999
	list, _, err := cpService.GetCouponInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.NormalResult{
		List: list,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetCouponList  分页获取coupon列表
// @Tags Coupon
// @Summary 获取coupon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponSearch true "下拉列表获取优惠券"
// @Success 200 {object} response.Response{data=response.PageResult,msg=object} "获取成功"
// @Router /coupon/list [Post]
func (cpApi *CouponApi) GetCouponList(c *gin.Context) {
	var pageInfo couponReq.CouponSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cpService.GetCouponInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}
