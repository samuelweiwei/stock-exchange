import service from '@/utils/request'
// @Tags Currencies
// @Summary 创建currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Currencies true "创建currencies表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /currencies/createCurrencies [post]
export const createCurrencies = (data) => {
  return service({
    url: '/currencies/createCurrencies',
    method: 'post',
    data
  })
}

// @Tags Currencies
// @Summary 删除currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Currencies true "删除currencies表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /currencies/deleteCurrencies [delete]
export const deleteCurrencies = (params) => {
  return service({
    url: '/currencies/deleteCurrencies',
    method: 'delete',
    params
  })
}

// @Tags Currencies
// @Summary 批量删除currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除currencies表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /currencies/deleteCurrencies [delete]
export const deleteCurrenciesByIds = (params) => {
  return service({
    url: '/currencies/deleteCurrenciesByIds',
    method: 'delete',
    params
  })
}

// @Tags Currencies
// @Summary 更新currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Currencies true "更新currencies表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /currencies/updateCurrencies [put]
export const updateCurrencies = (data) => {
  return service({
    url: '/currencies/updateCurrencies',
    method: 'put',
    data
  })
}

// @Tags Currencies
// @Summary 用id查询currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Currencies true "用id查询currencies表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /currencies/findCurrencies [get]
export const findCurrencies = (params) => {
  return service({
    url: '/currencies/findCurrencies',
    method: 'get',
    params
  })
}

// @Tags Currencies
// @Summary 分页获取currencies表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取currencies表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /currencies/getCurrenciesList [get]
export const getCurrenciesList = (params) => {
  return service({
    url: '/currencies/getCurrenciesList',
    method: 'get',
    params
  })
}

// @Tags Currencies
// @Summary 不需要鉴权的currencies表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.CurrenciesSearch true "分页获取currencies表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /currencies/getCurrenciesPublic [get]
export const getCurrenciesPublic = () => {
  return service({
    url: '/currencies/getCurrenciesPublic',
    method: 'get',
  })
}
