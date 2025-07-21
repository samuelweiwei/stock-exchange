import service from '@/utils/request'
// @Tags SymbolsCustom
// @Summary 创建symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SymbolsCustom true "创建symbolsCustom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /symbolsCustom/createSymbolsCustom [post]
export const createSymbolsCustom = (data) => {
  return service({
    url: '/symbolsCustom/createSymbolsCustom',
    method: 'post',
    data
  })
}

// @Tags SymbolsCustom
// @Summary 删除symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SymbolsCustom true "删除symbolsCustom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /symbolsCustom/deleteSymbolsCustom [delete]
export const deleteSymbolsCustom = (params) => {
  return service({
    url: '/symbolsCustom/deleteSymbolsCustom',
    method: 'delete',
    params
  })
}

// @Tags SymbolsCustom
// @Summary 批量删除symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除symbolsCustom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /symbolsCustom/deleteSymbolsCustom [delete]
export const deleteSymbolsCustomByIds = (params) => {
  return service({
    url: '/symbolsCustom/deleteSymbolsCustomByIds',
    method: 'delete',
    params
  })
}

// @Tags SymbolsCustom
// @Summary 更新symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SymbolsCustom true "更新symbolsCustom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /symbolsCustom/updateSymbolsCustom [put]
export const updateSymbolsCustom = (data) => {
  return service({
    url: '/symbolsCustom/updateSymbolsCustom',
    method: 'put',
    data
  })
}

// @Tags SymbolsCustom
// @Summary 用id查询symbolsCustom表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SymbolsCustom true "用id查询symbolsCustom表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /symbolsCustom/findSymbolsCustom [get]
export const findSymbolsCustom = (params) => {
  return service({
    url: '/symbolsCustom/findSymbolsCustom',
    method: 'get',
    params
  })
}

// @Tags SymbolsCustom
// @Summary 分页获取symbolsCustom表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取symbolsCustom表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /symbolsCustom/getSymbolsCustomList [get]
export const getSymbolsCustomList = (params) => {
  return service({
    url: '/symbolsCustom/getSymbolsCustomList',
    method: 'get',
    params
  })
}

// @Tags SymbolsCustom
// @Summary 不需要鉴权的symbolsCustom表接口
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsCustomSearch true "分页获取symbolsCustom表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /symbolsCustom/getSymbolsCustomPublic [get]
export const getSymbolsCustomPublic = () => {
  return service({
    url: '/symbolsCustom/getSymbolsCustomPublic',
    method: 'get',
  })
}
