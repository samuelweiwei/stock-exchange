import service from '@/utils/request'
// @Tags EarnInterestRates
// @Summary 创建earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnInterestRates true "创建earnInterestRates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /earnInterestRates/createEarnInterestRates [post]
export const createEarnInterestRates = (data) => {
  return service({
    url: '/earnInterestRates/createEarnInterestRates',
    method: 'post',
    data
  })
}

// @Tags EarnInterestRates
// @Summary 删除earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnInterestRates true "删除earnInterestRates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnInterestRates/deleteEarnInterestRates [delete]
export const deleteEarnInterestRates = (params) => {
  return service({
    url: '/earnInterestRates/deleteEarnInterestRates',
    method: 'delete',
    params
  })
}

// @Tags EarnInterestRates
// @Summary 批量删除earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除earnInterestRates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnInterestRates/deleteEarnInterestRates [delete]
export const deleteEarnInterestRatesByIds = (params) => {
  return service({
    url: '/earnInterestRates/deleteEarnInterestRatesByIds',
    method: 'delete',
    params
  })
}

// @Tags EarnInterestRates
// @Summary 更新earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnInterestRates true "更新earnInterestRates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /earnInterestRates/updateEarnInterestRates [put]
export const updateEarnInterestRates = (data) => {
  return service({
    url: '/earnInterestRates/updateEarnInterestRates',
    method: 'put',
    data
  })
}

// @Tags EarnInterestRates
// @Summary 用id查询earnInterestRates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EarnInterestRates true "用id查询earnInterestRates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /earnInterestRates/findEarnInterestRates [get]
export const findEarnInterestRates = (params) => {
  return service({
    url: '/earnInterestRates/findEarnInterestRates',
    method: 'get',
    params
  })
}

// @Tags EarnInterestRates
// @Summary 分页获取earnInterestRates表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取earnInterestRates表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /earnInterestRates/getEarnInterestRatesList [get]
export const getEarnInterestRatesList = (params) => {
  return service({
    url: '/earnInterestRates/getEarnInterestRatesList',
    method: 'get',
    params
  })
}

// @Tags EarnInterestRates
// @Summary 不需要鉴权的earnInterestRates表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnInterestRatesSearch true "分页获取earnInterestRates表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnInterestRates/getEarnInterestRatesPublic [get]
export const getEarnInterestRatesPublic = () => {
  return service({
    url: '/earnInterestRates/getEarnInterestRatesPublic',
    method: 'get',
  })
}
