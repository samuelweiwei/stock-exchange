import service from '@/utils/request'
// @Tags ContractLeverage
// @Summary 创建contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractLeverage true "创建contractLeverage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /contractLeverage/createContractLeverage [post]
export const createContractLeverage = (data) => {
  return service({
    url: '/contractLeverage/createContractLeverage',
    method: 'post',
    data
  })
}

// @Tags ContractLeverage
// @Summary 删除contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractLeverage true "删除contractLeverage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractLeverage/deleteContractLeverage [delete]
export const deleteContractLeverage = (params) => {
  return service({
    url: '/contractLeverage/deleteContractLeverage',
    method: 'delete',
    params
  })
}

// @Tags ContractLeverage
// @Summary 批量删除contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除contractLeverage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractLeverage/deleteContractLeverage [delete]
export const deleteContractLeverageByIds = (params) => {
  return service({
    url: '/contractLeverage/deleteContractLeverageByIds',
    method: 'delete',
    params
  })
}

// @Tags ContractLeverage
// @Summary 更新contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractLeverage true "更新contractLeverage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contractLeverage/updateContractLeverage [put]
export const updateContractLeverage = (data) => {
  return service({
    url: '/contractLeverage/updateContractLeverage',
    method: 'put',
    data
  })
}

// @Tags ContractLeverage
// @Summary 用id查询contractLeverage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContractLeverage true "用id查询contractLeverage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contractLeverage/findContractLeverage [get]
export const findContractLeverage = (params) => {
  return service({
    url: '/contractLeverage/findContractLeverage',
    method: 'get',
    params
  })
}

// @Tags ContractLeverage
// @Summary 分页获取contractLeverage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取contractLeverage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contractLeverage/getContractLeverageList [get]
export const getContractLeverageList = (params) => {
  return service({
    url: '/contractLeverage/getContractLeverageList',
    method: 'get',
    params
  })
}

// @Tags ContractLeverage
// @Summary 不需要鉴权的contractLeverage表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractLeverageSearch true "分页获取contractLeverage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractLeverage/getContractLeveragePublic [get]
export const getContractLeveragePublic = () => {
  return service({
    url: '/contractLeverage/getContractLeveragePublic',
    method: 'get',
  })
}