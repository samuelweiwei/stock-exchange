import service from '@/utils/request'
// @Tags ContractEntrust
// @Summary 创建contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractEntrust true "创建contractEntrust表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /contractEntrust/createContractEntrust [post]
export const createContractEntrust = (data) => {
  return service({
    url: '/contractEntrust/createContractEntrust',
    method: 'post',
    data
  })
}

// @Tags ContractEntrust
// @Summary 删除contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractEntrust true "删除contractEntrust表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractEntrust/deleteContractEntrust [delete]
export const deleteContractEntrust = (params) => {
  return service({
    url: '/contractEntrust/deleteContractEntrust',
    method: 'delete',
    params
  })
}

// @Tags ContractEntrust
// @Summary 批量删除contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除contractEntrust表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractEntrust/deleteContractEntrust [delete]
export const deleteContractEntrustByIds = (params) => {
  return service({
    url: '/contractEntrust/deleteContractEntrustByIds',
    method: 'delete',
    params
  })
}

// @Tags ContractEntrust
// @Summary 更新contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractEntrust true "更新contractEntrust表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contractEntrust/updateContractEntrust [put]
export const updateContractEntrust = (data) => {
  return service({
    url: '/contractEntrust/updateContractEntrust',
    method: 'put',
    data
  })
}

// @Tags ContractEntrust
// @Summary 用id查询contractEntrust表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContractEntrust true "用id查询contractEntrust表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contractEntrust/findContractEntrust [get]
export const findContractEntrust = (params) => {
  return service({
    url: '/contractEntrust/findContractEntrust',
    method: 'get',
    params
  })
}

// @Tags ContractEntrust
// @Summary 分页获取contractEntrust表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取contractEntrust表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contractEntrust/getContractEntrustList [get]
export const getContractEntrustList = (params) => {
  return service({
    url: '/contractEntrust/getContractEntrustList',
    method: 'get',
    params
  })
}

// @Tags ContractEntrust
// @Summary 不需要鉴权的contractEntrust表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractEntrustSearch true "分页获取contractEntrust表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractEntrust/getContractEntrustPublic [get]
export const getContractEntrustPublic = () => {
  return service({
    url: '/contractEntrust/getContractEntrustPublic',
    method: 'get',
  })
}
