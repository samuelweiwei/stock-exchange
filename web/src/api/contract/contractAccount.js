import service from '@/utils/request'
// @Tags ContractAccount
// @Summary 创建contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractAccount true "创建contractAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /contractAccount/createContractAccount [post]
export const createContractAccount = (data) => {
  return service({
    url: '/contractAccount/createContractAccount',
    method: 'post',
    data
  })
}

// @Tags ContractAccount
// @Summary 删除contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractAccount true "删除contractAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractAccount/deleteContractAccount [delete]
export const deleteContractAccount = (params) => {
  return service({
    url: '/contractAccount/deleteContractAccount',
    method: 'delete',
    params
  })
}

// @Tags ContractAccount
// @Summary 批量删除contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除contractAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractAccount/deleteContractAccount [delete]
export const deleteContractAccountByIds = (params) => {
  return service({
    url: '/contractAccount/deleteContractAccountByIds',
    method: 'delete',
    params
  })
}

// @Tags ContractAccount
// @Summary 更新contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractAccount true "更新contractAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contractAccount/updateContractAccount [put]
export const updateContractAccount = (data) => {
  return service({
    url: '/contractAccount/updateContractAccount',
    method: 'put',
    data
  })
}

// @Tags ContractAccount
// @Summary 用id查询contractAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContractAccount true "用id查询contractAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contractAccount/findContractAccount [get]
export const findContractAccount = (params) => {
  return service({
    url: '/contractAccount/findContractAccount',
    method: 'get',
    params
  })
}

// @Tags ContractAccount
// @Summary 分页获取contractAccount表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取contractAccount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contractAccount/getContractAccountList [get]
export const getContractAccountList = (params) => {
  return service({
    url: '/contractAccount/getContractAccountList',
    method: 'get',
    params
  })
}

// @Tags ContractAccount
// @Summary 不需要鉴权的contractAccount表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractAccountSearch true "分页获取contractAccount表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractAccount/getContractAccountPublic [get]
export const getContractAccountPublic = () => {
  return service({
    url: '/contractAccount/getContractAccountPublic',
    method: 'get',
  })
}
// ChangeAccountMargin 合约资金划转
// @Tags ContractAccount
// @Summary 合约资金划转
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /contractAccount/changeAccountMargin [POST]
export const changeAccountMargin = () => {
  return service({
    url: '/contractAccount/changeAccountMargin',
    method: 'POST'
  })
}
