import service from '@/utils/request'
// @Tags Symbols
// @Summary 创建symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Symbols true "创建symbols表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /symbols/createSymbols [post]
export const createSymbols = (data) => {
  return service({
    url: '/symbols/createSymbols',
    method: 'post',
    data
  })
}

// @Tags Symbols
// @Summary 删除symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Symbols true "删除symbols表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /symbols/deleteSymbols [delete]
export const deleteSymbols = (params) => {
  return service({
    url: '/symbols/deleteSymbols',
    method: 'delete',
    params
  })
}

// @Tags Symbols
// @Summary 批量删除symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除symbols表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /symbols/deleteSymbols [delete]
export const deleteSymbolsByIds = (params) => {
  return service({
    url: '/symbols/deleteSymbolsByIds',
    method: 'delete',
    params
  })
}

// @Tags Symbols
// @Summary 更新symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Symbols true "更新symbols表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /symbols/updateSymbols [put]
export const updateSymbols = (data) => {
  return service({
    url: '/symbols/updateSymbols',
    method: 'put',
    data
  })
}

// @Tags Symbols
// @Summary 用id查询symbols表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Symbols true "用id查询symbols表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /symbols/findSymbols [get]
export const findSymbols = (params) => {
  return service({
    url: '/symbols/findSymbols',
    method: 'get',
    params
  })
}

// @Tags Symbols
// @Summary 分页获取symbols表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取symbols表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /symbols/getSymbolsList [get]
export const getSymbolsList = (params) => {
  return service({
    url: '/symbols/getSymbolsList',
    method: 'get',
    params
  })
}

// @Tags Symbols
// @Summary 不需要鉴权的symbols表接口
// @accept application/json
// @Produce application/json
// @Param data query symbolReq.SymbolsSearch true "分页获取symbols表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /symbols/getSymbolsPublic [get]
export const getSymbolsPublic = () => {
  return service({
    url: '/symbols/getSymbolsPublic',
    method: 'get',
  })
}
