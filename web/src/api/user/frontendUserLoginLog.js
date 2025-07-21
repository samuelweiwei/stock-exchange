import service from '@/utils/request'
// @Tags FrontendUserLoginLog
// @Summary 创建frontendUserLoginLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrontendUserLoginLog true "创建frontendUserLoginLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /frontendUserLoginLog/createFrontendUserLoginLog [post]
export const createFrontendUserLoginLog = (data) => {
  return service({
    url: '/frontendUserLoginLog/createFrontendUserLoginLog',
    method: 'post',
    data
  })
}

// @Tags FrontendUserLoginLog
// @Summary 删除frontendUserLoginLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrontendUserLoginLog true "删除frontendUserLoginLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frontendUserLoginLog/deleteFrontendUserLoginLog [delete]
export const deleteFrontendUserLoginLog = (params) => {
  return service({
    url: '/frontendUserLoginLog/deleteFrontendUserLoginLog',
    method: 'delete',
    params
  })
}

// @Tags FrontendUserLoginLog
// @Summary 批量删除frontendUserLoginLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frontendUserLoginLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frontendUserLoginLog/deleteFrontendUserLoginLog [delete]
export const deleteFrontendUserLoginLogByIds = (params) => {
  return service({
    url: '/frontendUserLoginLog/deleteFrontendUserLoginLogByIds',
    method: 'delete',
    params
  })
}

// @Tags FrontendUserLoginLog
// @Summary 更新frontendUserLoginLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrontendUserLoginLog true "更新frontendUserLoginLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /frontendUserLoginLog/updateFrontendUserLoginLog [put]
export const updateFrontendUserLoginLog = (data) => {
  return service({
    url: '/frontendUserLoginLog/updateFrontendUserLoginLog',
    method: 'put',
    data
  })
}

// @Tags FrontendUserLoginLog
// @Summary 用id查询frontendUserLoginLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FrontendUserLoginLog true "用id查询frontendUserLoginLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /frontendUserLoginLog/findFrontendUserLoginLog [get]
export const findFrontendUserLoginLog = (params) => {
  return service({
    url: '/frontendUserLoginLog/findFrontendUserLoginLog',
    method: 'get',
    params
  })
}

// @Tags FrontendUserLoginLog
// @Summary 分页获取frontendUserLoginLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取frontendUserLoginLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /frontendUserLoginLog/getFrontendUserLoginLogList [get]
export const getFrontendUserLoginLogList = (params) => {
  return service({
    url: '/frontendUserLoginLog/getFrontendUserLoginLogList',
    method: 'get',
    params
  })
}

// @Tags FrontendUserLoginLog
// @Summary 不需要鉴权的frontendUserLoginLog表接口
// @accept application/json
// @Produce application/json
// @Param data query userReq.FrontendUserLoginLogSearch true "分页获取frontendUserLoginLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /frontendUserLoginLog/getFrontendUserLoginLogPublic [get]
export const getFrontendUserLoginLogPublic = () => {
  return service({
    url: '/frontendUserLoginLog/getFrontendUserLoginLogPublic',
    method: 'get',
  })
}
