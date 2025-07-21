import service from '@/utils/request'
// @Tags UserAccountFlow
// @Summary 创建userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAccountFlow true "创建userAccountFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userAccountFlow/createUserAccountFlow [post]
export const createUserAccountFlow = (data) => {
  return service({
    url: '/userAccountFlow/createUserAccountFlow',
    method: 'post',
    data
  })
}

// @Tags UserAccountFlow
// @Summary 删除userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAccountFlow true "删除userAccountFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAccountFlow/deleteUserAccountFlow [delete]
export const deleteUserAccountFlow = (params) => {
  return service({
    url: '/userAccountFlow/deleteUserAccountFlow',
    method: 'delete',
    params
  })
}

// @Tags UserAccountFlow
// @Summary 批量删除userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userAccountFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAccountFlow/deleteUserAccountFlow [delete]
export const deleteUserAccountFlowByIds = (params) => {
  return service({
    url: '/userAccountFlow/deleteUserAccountFlowByIds',
    method: 'delete',
    params
  })
}

// @Tags UserAccountFlow
// @Summary 更新userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAccountFlow true "更新userAccountFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAccountFlow/updateUserAccountFlow [put]
export const updateUserAccountFlow = (data) => {
  return service({
    url: '/userAccountFlow/updateUserAccountFlow',
    method: 'put',
    data
  })
}

// @Tags UserAccountFlow
// @Summary 用id查询userAccountFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserAccountFlow true "用id查询userAccountFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAccountFlow/findUserAccountFlow [get]
export const findUserAccountFlow = (params) => {
  return service({
    url: '/userAccountFlow/findUserAccountFlow',
    method: 'get',
    params
  })
}

// @Tags UserAccountFlow
// @Summary 分页获取userAccountFlow表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userAccountFlow表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAccountFlow/getUserAccountFlowList [get]
export const getUserAccountFlowList = (params) => {
  return service({
    url: '/userAccountFlow/getUserAccountFlowList',
    method: 'get',
    params
  })
}

// @Tags UserAccountFlow
// @Summary 不需要鉴权的userAccountFlow表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserAccountFlowSearch true "分页获取userAccountFlow表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userAccountFlow/getUserAccountFlowPublic [get]
export const getUserAccountFlowPublic = () => {
  return service({
    url: '/userAccountFlow/getUserAccountFlowPublic',
    method: 'get',
  })
}
