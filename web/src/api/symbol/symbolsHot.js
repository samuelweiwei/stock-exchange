import service from '@/utils/request'
// @Tags SymbolsHot
// @Summary 创建symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SymbolsHot true "创建symbolsHot表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /symbolsHot/createSymbolsHot [post]
export const createSymbolsHot = (data) => {
  return service({
    url: '/symbolsHot/createSymbolsHot',
    method: 'post',
    data
  })
}

// @Tags SymbolsHot
// @Summary 删除symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SymbolsHot true "删除symbolsHot表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /symbolsHot/deleteSymbolsHot [delete]
export const deleteSymbolsHot = (params) => {
  return service({
    url: '/symbolsHot/deleteSymbolsHot',
    method: 'delete',
    params
  })
}

// @Tags SymbolsHot
// @Summary 批量删除symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除symbolsHot表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /symbolsHot/deleteSymbolsHot [delete]
export const deleteSymbolsHotByIds = (params) => {
  return service({
    url: '/symbolsHot/deleteSymbolsHotByIds',
    method: 'delete',
    params
  })
}

// @Tags SymbolsHot
// @Summary 更新symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SymbolsHot true "更新symbolsHot表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /symbolsHot/updateSymbolsHot [put]
export const updateSymbolsHot = (data) => {
  return service({
    url: '/symbolsHot/updateSymbolsHot',
    method: 'put',
    data
  })
}

// @Tags SymbolsHot
// @Summary 用id查询symbolsHot表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SymbolsHot true "用id查询symbolsHot表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /symbolsHot/findSymbolsHot [get]
export const findSymbolsHot = (params) => {
  return service({
    url: '/symbolsHot/findSymbolsHot',
    method: 'get',
    params
  })
}

// @Tags SymbolsHot
// @Summary 分页获取symbolsHot表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取symbolsHot表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /symbolsHot/getSymbolsHotList [get]
export const getSymbolsHotList = (params) => {
  return service({
    url: '/symbolsHot/getSymbolsHotList',
    method: 'get',
    params
  })
}

// @Tags SymbolsHot
// @Summary 不需要鉴权的symbolsHot表接口
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsHotSearch true "分页获取symbolsHot表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /symbolsHot/getSymbolsHotPublic [get]
export const getSymbolsHotPublic = () => {
  return service({
    url: '/symbolsHot/getSymbolsHotPublic',
    method: 'get',
  })
}
