import service from '@/utils/request'
// @Tags CouponIssued
// @Summary 创建couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CouponIssued true "创建couponIssued表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /couponIssued/createCouponIssued [post]
export const createCouponIssued = (data) => {
  return service({
    url: '/couponIssued/createCouponIssued',
    method: 'post',
    data
  })
}

// @Tags CouponIssued
// @Summary 删除couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CouponIssued true "删除couponIssued表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /couponIssued/deleteCouponIssued [delete]
export const deleteCouponIssued = (params) => {
  return service({
    url: '/couponIssued/deleteCouponIssued',
    method: 'delete',
    params
  })
}

// @Tags CouponIssued
// @Summary 批量删除couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除couponIssued表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /couponIssued/deleteCouponIssued [delete]
export const deleteCouponIssuedByIds = (params) => {
  return service({
    url: '/couponIssued/deleteCouponIssuedByIds',
    method: 'delete',
    params
  })
}

// @Tags CouponIssued
// @Summary 更新couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CouponIssued true "更新couponIssued表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /couponIssued/updateCouponIssued [put]
export const updateCouponIssued = (data) => {
  return service({
    url: '/couponIssued/updateCouponIssued',
    method: 'put',
    data
  })
}

// @Tags CouponIssued
// @Summary 用id查询couponIssued表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CouponIssued true "用id查询couponIssued表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /couponIssued/findCouponIssued [get]
export const findCouponIssued = (params) => {
  return service({
    url: '/couponIssued/findCouponIssued',
    method: 'get',
    params
  })
}

// @Tags CouponIssued
// @Summary 分页获取couponIssued表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取couponIssued表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /couponIssued/getCouponIssuedList [get]
export const getCouponIssuedList = (params) => {
  return service({
    url: '/couponIssued/getCouponIssuedList',
    method: 'get',
    params
  })
}

// @Tags CouponIssued
// @Summary 不需要鉴权的couponIssued表接口
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponIssuedSearch true "分页获取couponIssued表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /couponIssued/getCouponIssuedPublic [get]
export const getCouponIssuedPublic = () => {
  return service({
    url: '/couponIssued/getCouponIssuedPublic',
    method: 'get',
  })
}
