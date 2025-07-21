import service from '@/utils/request'
// @Tags ContractOrder
// @Summary 创建contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractOrder true "创建contractOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /contractOrder/createContractOrder [post]
export const createContractOrder = (data) => {
  return service({
    url: '/contractOrder/createContractOrder',
    method: 'post',
    data
  })
}

// @Tags ContractOrder
// @Summary 删除contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractOrder true "删除contractOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractOrder/deleteContractOrder [delete]
export const deleteContractOrder = (params) => {
  return service({
    url: '/contractOrder/deleteContractOrder',
    method: 'delete',
    params
  })
}

// @Tags ContractOrder
// @Summary 批量删除contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除contractOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractOrder/deleteContractOrder [delete]
export const deleteContractOrderByIds = (params) => {
  return service({
    url: '/contractOrder/deleteContractOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags ContractOrder
// @Summary 更新contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractOrder true "更新contractOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contractOrder/updateContractOrder [put]
export const updateContractOrder = (data) => {
  return service({
    url: '/contractOrder/updateContractOrder',
    method: 'put',
    data
  })
}

// @Tags ContractOrder
// @Summary 用id查询contractOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContractOrder true "用id查询contractOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contractOrder/findContractOrder [get]
export const findContractOrder = (params) => {
  return service({
    url: '/contractOrder/findContractOrder',
    method: 'get',
    params
  })
}

// @Tags ContractOrder
// @Summary 分页获取contractOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取contractOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contractOrder/getContractOrderList [get]
export const getContractOrderList = (params) => {
  return service({
    url: '/contractOrder/getContractOrderList',
    method: 'get',
    params
  })
}

// @Tags ContractOrder
// @Summary 不需要鉴权的contractOrder表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractOrderSearch true "分页获取contractOrder表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractOrder/getContractOrderPublic [get]
export const getContractOrderPublic = () => {
  return service({
    url: '/contractOrder/getContractOrderPublic',
    method: 'get',
  })
}
// CloseContractOrder 平仓方法
// @Tags ContractOrder
// @Summary 平仓方法
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /contractOrder/closeContractOrder [POST]
export const  closeContractOrder = () => {
  return service({
    url: '/contractOrder/closeContractOrder',
    method: 'POST'
  })
}
