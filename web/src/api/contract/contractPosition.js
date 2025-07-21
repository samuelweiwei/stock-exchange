import service from '@/utils/request'
// @Tags ContractPosition
// @Summary 创建contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractPosition true "创建contractPosition表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /contractPosition/createContractPosition [post]
export const createContractPosition = (data) => {
  return service({
    url: '/contractPosition/createContractPosition',
    method: 'post',
    data
  })
}

// @Tags ContractPosition
// @Summary 删除contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractPosition true "删除contractPosition表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractPosition/deleteContractPosition [delete]
export const deleteContractPosition = (params) => {
  return service({
    url: '/contractPosition/deleteContractPosition',
    method: 'delete',
    params
  })
}

// @Tags ContractPosition
// @Summary 批量删除contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除contractPosition表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contractPosition/deleteContractPosition [delete]
export const deleteContractPositionByIds = (params) => {
  return service({
    url: '/contractPosition/deleteContractPositionByIds',
    method: 'delete',
    params
  })
}

// @Tags ContractPosition
// @Summary 更新contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ContractPosition true "更新contractPosition表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contractPosition/updateContractPosition [put]
export const updateContractPosition = (data) => {
  return service({
    url: '/contractPosition/updateContractPosition',
    method: 'put',
    data
  })
}

// @Tags ContractPosition
// @Summary 用id查询contractPosition表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ContractPosition true "用id查询contractPosition表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contractPosition/findContractPosition [get]
export const findContractPosition = (params) => {
  return service({
    url: '/contractPosition/findContractPosition',
    method: 'get',
    params
  })
}

// @Tags ContractPosition
// @Summary 分页获取contractPosition表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取contractPosition表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contractPosition/getContractPositionList [get]
export const getContractPositionList = (params) => {
  return service({
    url: '/contractPosition/getContractPositionList',
    method: 'get',
    params
  })
}

// @Tags ContractPosition
// @Summary 不需要鉴权的contractPosition表接口
// @accept application/json
// @Produce application/json
// @Param data query contractReq.ContractPositionSearch true "分页获取contractPosition表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /contractPosition/getContractPositionPublic [get]
export const getContractPositionPublic = () => {
  return service({
    url: '/contractPosition/getContractPositionPublic',
    method: 'get',
  })
}
