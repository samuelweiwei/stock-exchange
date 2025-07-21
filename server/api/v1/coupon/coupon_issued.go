package coupon

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/coupon"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	couponReq "github.com/flipped-aurora/gin-vue-admin/server/model/coupon/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const MaxPageSize = 999

type CouponIssuedApi struct{}

// CreateCouponIssued 发放优惠券
// @Tags CouponIssued
// @Summary 发放优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body couponReq.IssueCoupon true "发放优惠券"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /coupon/issue [post]
func (couponIssuedApi *CouponIssuedApi) CreateCouponIssued(c *gin.Context) {
	var (
		couponIssued couponReq.IssueCoupon
		res          []string
	)
	err := c.ShouldBindJSON(&couponIssued)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err = couponIssuedService.CreateCouponIssued(&couponIssued)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		switch err.Error() {
		case i18n.CouponDoesNotExist:
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CouponDoesNotExist, response.ERROR, res), c)
		case i18n.UserDoesNotExist:
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UserDoesNotExist, response.ERROR), c)
		case i18n.CouponDoesNotActive:
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CouponDoesNotActive, response.ERROR), c)
		default:
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CreationFailed, response.ERROR), c)
		}
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// DeleteCouponIssued 删除couponIssued表
// @Tags CouponIssued
// @Summary 删除couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coupon.CouponIssued true "删除couponIssued表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /coupon/issued/delete [delete]
func (couponIssuedApi *CouponIssuedApi) DeleteCouponIssued(c *gin.Context) {
	id := c.Query("id")
	err := couponIssuedService.DeleteCouponIssued(id)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.DeleteFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// UpdateCouponIssued 更新优惠券
// @Tags CouponIssued
// @Summary 使用优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coupon.CouponIssued true "更新优惠券"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /coupon/issued/update [put]
func (couponIssuedApi *CouponIssuedApi) UpdateCouponIssued(c *gin.Context) {
	var couponIssued coupon.CouponIssued
	err := c.ShouldBindJSON(&couponIssued)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = couponIssuedService.UpdateCouponIssued(couponIssued)
	if err != nil {
		_ = c.Error(err)
		response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.UPDATE_FAILED), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// UseCoupon 使用优惠券
// @Tags CouponIssued
// @Summary 使用优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coupon.CouponIssued true "使用优惠券"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /coupon/issued/use [put]
func (couponIssuedApi *CouponIssuedApi) UseCoupon(c *gin.Context) {
	var couponIssued couponReq.UseIssuedCoupon
	err := c.ShouldBindJSON(&couponIssued)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = couponIssuedService.Use(&couponIssued)
	if err != nil {
		_ = c.Error(err)
		response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.UPDATE_FAILED), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// FindCouponIssued 用id查询couponIssued表
// @Tags CouponIssued
// @Summary 用id查询couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "查询绑定用户的优惠券"
// FindCouponIssued @Success 200 {object} response.Response{data=coupon.CouponIssued,msg=string} "查询成功"
// @Router /coupon/issued/find [get]
func (couponIssuedApi *CouponIssuedApi) FindCouponIssued(c *gin.Context) {
	id := c.Query("id")
	recouponIssued, err := couponIssuedService.GetCouponIssued(id)
	if err != nil {
		_ = c.Error(err)
		response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithData(recouponIssued, c)
}

// GetCouponIssuedList 分页获取couponIssued表列表
// @Tags CouponIssued
// @Summary 分页获取couponIssued表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponIssuedSearch true "分页获取couponIssued表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /coupon/front/issued/list [get]
func (couponIssuedApi *CouponIssuedApi) GetCouponIssuedList(c *gin.Context) {
	var pageInfo couponReq.CouponIssuedSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = uint64(utils.GetUserIDFrontUser(c))
	list, total, err := couponIssuedService.GetCouponIssuedInfoList(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.SuccessfullyObtained, response.SUCCESS), c)
}

// AdminGetCouponIssuedList 分页获取couponIssued表列表
// @Tags CouponIssued
// @Summary 分页获取couponIssued表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponIssuedSearch true "分页获取couponIssued表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /coupon/issued/list [get]
func (couponIssuedApi *CouponIssuedApi) AdminGetCouponIssuedList(c *gin.Context) {
	var (
		req couponReq.CouponIssuedSearch
	)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := couponIssuedService.AdminGetCouponIssuedInfoList(req)
	if err != nil {
		_ = c.Error(err)
		response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.SuccessfullyObtained, response.SUCCESS), c)
}

func (couponIssuedApi *CouponIssuedApi) getUsers(couponReq *couponReq.CouponIssuedSearch) (res map[uint]*user.FrontendUsers, userIdList []uint, err error) {
	var (
		userReqSearch userReq.FrontendUsersSearch
	)
	userReqSearch.UserType = couponReq.UserType
	userReqSearch.Email = couponReq.Email
	userReqSearch.RootId = couponReq.SuperiorId
	userReqSearch.Page = 1
	userReqSearch.PageSize = MaxPageSize
	userReqSearch.UserIDList = fmt.Sprintf("%v", couponReq.UserId)
	infoList, _, err := service.ServiceGroupApp.UserServiceGroup.FrontendUsersService.GetFrontendUsersInfoList(userReqSearch)
	if err != nil {
		return res, userIdList, err
	}
	for _, v := range infoList {
		res[v.ID] = &v
		userIdList = append(userIdList, v.ID)
	}
	return res, userIdList, nil
}

// Options  优惠券下拉列表
// @Tags CouponIssued
// @Summary 优惠券下拉列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.NormalResult,msg=object} "获取成功"
// @Router /coupon/issued/options [get]
func (couponIssuedApi *CouponIssuedApi) Options(c *gin.Context) {
	var pageInfo couponReq.CouponIssuedSearch
	pageInfo.Page = 1
	pageInfo.PageSize = 999

	list, _, err := couponIssuedService.GetCouponIssuedInfoList(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.NormalResult{
		List: list,
	}, i18n.Message(request.GetLanguageTag(c), i18n.SuccessfullyObtained, response.SUCCESS), c)
}
