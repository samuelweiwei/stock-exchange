import service from '@/utils/request'
// @Tags FrontendUsers
// @Summary 创建frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrontendUsers true "创建frontendUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /frontendUsers/createFrontendUsers [post]
export const createFrontendUsers = (data) => {
  return service({
    url: '/frontendUsers/createFrontendUsers',
    method: 'post',
    data
  })
}

// @Tags FrontendUsers
// @Summary 删除frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrontendUsers true "删除frontendUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frontendUsers/deleteFrontendUsers [delete]
export const deleteFrontendUsers = (params) => {
  return service({
    url: '/frontendUsers/deleteFrontendUsers',
    method: 'delete',
    params
  })
}

// @Tags FrontendUsers
// @Summary 批量删除frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frontendUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frontendUsers/deleteFrontendUsers [delete]
export const deleteFrontendUsersByIds = (params) => {
  return service({
    url: '/frontendUsers/deleteFrontendUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags FrontendUsers
// @Summary 更新frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrontendUsers true "更新frontendUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /frontendUsers/updateFrontendUsers [put]
export const updateFrontendUsers = (data) => {
  return service({
    url: '/frontendUsers/updateFrontendUsers',
    method: 'put',
    data
  })
}

// @Tags FrontendUsers
// @Summary 用id查询frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FrontendUsers true "用id查询frontendUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /frontendUsers/findFrontendUsers [get]
export const findFrontendUsers = (params) => {
  return service({
    url: '/frontendUsers/findFrontendUsers',
    method: 'get',
    params
  })
}

// @Tags FrontendUsers
// @Summary 分页获取frontendUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取frontendUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /frontendUsers/getFrontendUsersList [get]
export const getFrontendUsersList = (params) => {
  return service({
    url: '/frontendUsers/getFrontendUsersList',
    method: 'get',
    params
  })
}

// @Tags FrontendUsers
// @Summary 不需要鉴权的frontendUsers表接口
// @accept application/json
// @Produce application/json
// @Param data query userReq.FrontendUsersSearch true "分页获取frontendUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /frontendUsers/getFrontendUsersPublic [get]
export const getFrontendUsersPublic = () => {
  return service({
    url: '/frontendUsers/getFrontendUsersPublic',
    method: 'get',
  })
}
