import service from '@/utils/request'
// @Tags CouponRecord
// @Summary 创建couponRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CouponRecord true "创建couponRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /couponRecord/createCouponRecord [post]
export const createCouponRecord = (data) => {
  return service({
    url: '/couponRecord/createCouponRecord',
    method: 'post',
    data
  })
}

// @Tags CouponRecord
// @Summary 删除couponRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CouponRecord true "删除couponRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /couponRecord/deleteCouponRecord [delete]
export const deleteCouponRecord = (params) => {
  return service({
    url: '/couponRecord/deleteCouponRecord',
    method: 'delete',
    params
  })
}

// @Tags CouponRecord
// @Summary 批量删除couponRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除couponRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /couponRecord/deleteCouponRecord [delete]
export const deleteCouponRecordByIds = (params) => {
  return service({
    url: '/couponRecord/deleteCouponRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags CouponRecord
// @Summary 更新couponRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CouponRecord true "更新couponRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /couponRecord/updateCouponRecord [put]
export const updateCouponRecord = (data) => {
  return service({
    url: '/couponRecord/updateCouponRecord',
    method: 'put',
    data
  })
}

// @Tags CouponRecord
// @Summary 用id查询couponRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CouponRecord true "用id查询couponRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /couponRecord/findCouponRecord [get]
export const findCouponRecord = (params) => {
  return service({
    url: '/couponRecord/findCouponRecord',
    method: 'get',
    params
  })
}

// @Tags CouponRecord
// @Summary 分页获取couponRecord表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取couponRecord表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /couponRecord/getCouponRecordList [get]
export const getCouponRecordList = (params) => {
  return service({
    url: '/couponRecord/getCouponRecordList',
    method: 'get',
    params
  })
}

// @Tags CouponRecord
// @Summary 不需要鉴权的couponRecord表接口
// @accept application/json
// @Produce application/json
// @Param data query couponReq.CouponRecordSearch true "分页获取couponRecord表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /couponRecord/getCouponRecordPublic [get]
export const getCouponRecordPublic = () => {
  return service({
    url: '/couponRecord/getCouponRecordPublic',
    method: 'get',
  })
}
