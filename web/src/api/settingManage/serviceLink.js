import service from '@/utils/request'
// @Tags ServiceLink
// @Summary 创建serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceLink true "创建serviceLink表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serviceLink/createServiceLink [post]
export const createServiceLink = (data) => {
  return service({
    url: '/serviceLink/createServiceLink',
    method: 'post',
    data
  })
}

// @Tags ServiceLink
// @Summary 删除serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceLink true "删除serviceLink表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceLink/deleteServiceLink [delete]
export const deleteServiceLink = (params) => {
  return service({
    url: '/serviceLink/deleteServiceLink',
    method: 'delete',
    params
  })
}

// @Tags ServiceLink
// @Summary 批量删除serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除serviceLink表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceLink/deleteServiceLink [delete]
export const deleteServiceLinkByIds = (params) => {
  return service({
    url: '/serviceLink/deleteServiceLinkByIds',
    method: 'delete',
    params
  })
}

// @Tags ServiceLink
// @Summary 更新serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceLink true "更新serviceLink表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serviceLink/updateServiceLink [put]
export const updateServiceLink = (data) => {
  return service({
    url: '/serviceLink/updateServiceLink',
    method: 'put',
    data
  })
}

// @Tags ServiceLink
// @Summary 用id查询serviceLink表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ServiceLink true "用id查询serviceLink表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serviceLink/findServiceLink [get]
export const findServiceLink = (params) => {
  return service({
    url: '/serviceLink/findServiceLink',
    method: 'get',
    params
  })
}

// @Tags ServiceLink
// @Summary 分页获取serviceLink表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取serviceLink表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceLink/getServiceLinkList [get]
export const getServiceLinkList = (params) => {
  return service({
    url: '/serviceLink/getServiceLinkList',
    method: 'get',
    params
  })
}

// @Tags ServiceLink
// @Summary 不需要鉴权的serviceLink表接口
// @accept application/json
// @Produce application/json
// @Param data query settingManageReq.ServiceLinkSearch true "分页获取serviceLink表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /serviceLink/getServiceLinkPublic [get]
export const getServiceLinkPublic = () => {
  return service({
    url: '/serviceLink/getServiceLinkPublic',
    method: 'get',
  })
}
