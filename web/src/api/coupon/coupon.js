import service from '@/utils/request'
// @Tags Coupon
// @Summary 创建coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coupon true "创建coupon表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cp/createCoupon [post]
export const createCoupon = (data) => {
  return service({
    url: '/cp/createCoupon',
    method: 'post',
    data
  })
}

// @Tags Coupon
// @Summary 删除coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coupon true "删除coupon表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cp/deleteCoupon [delete]
export const deleteCoupon = (params) => {
  return service({
    url: '/cp/deleteCoupon',
    method: 'delete',
    params
  })
}

// @Tags Coupon
// @Summary 批量删除coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除coupon表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cp/deleteCoupon [delete]
export const deleteCouponByIds = (params) => {
  return service({
    url: '/cp/deleteCouponByIds',
    method: 'delete',
    params
  })
}

// @Tags Coupon
// @Summary 更新coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coupon true "更新coupon表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cp/updateCoupon [put]
export const updateCoupon = (data) => {
  return service({
    url: '/cp/updateCoupon',
    method: 'put',
    data
  })
}

// @Tags Coupon
// @Summary 用id查询coupon表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Coupon true "用id查询coupon表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cp/findCoupon [get]
export const findCoupon = (params) => {
  return service({
    url: '/cp/findCoupon',
    method: 'get',
    params
  })
}

// @Tags Coupon
// @Summary 分页获取coupon表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取coupon表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cp/getCouponList [get]
export const getCouponList = (params) => {
  return service({
    url: '/cp/getCouponList',
    method: 'get',
    params
  })
}

// @Tags Coupon
// @Summary 不需要鉴权的coupon表接口
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponSearch true "分页获取coupon表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cp/getCouponPublic [get]
export const getCouponPublic = () => {
  return service({
    url: '/cp/getCouponPublic',
    method: 'get',
  })
}
