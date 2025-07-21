import service from '@/utils/request'
// @Tags UserFundAccounts
// @Summary 创建userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFundAccounts true "创建userFundAccounts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userFundAccounts/createUserFundAccounts [post]
export const createUserFundAccounts = (data) => {
  return service({
    url: '/userFundAccounts/createUserFundAccounts',
    method: 'post',
    data
  })
}

// @Tags UserFundAccounts
// @Summary 删除userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFundAccounts true "删除userFundAccounts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFundAccounts/deleteUserFundAccounts [delete]
export const deleteUserFundAccounts = (params) => {
  return service({
    url: '/userFundAccounts/deleteUserFundAccounts',
    method: 'delete',
    params
  })
}

// @Tags UserFundAccounts
// @Summary 批量删除userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userFundAccounts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userFundAccounts/deleteUserFundAccounts [delete]
export const deleteUserFundAccountsByIds = (params) => {
  return service({
    url: '/userFundAccounts/deleteUserFundAccountsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserFundAccounts
// @Summary 更新userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserFundAccounts true "更新userFundAccounts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userFundAccounts/updateUserFundAccounts [put]
export const updateUserFundAccounts = (data) => {
  return service({
    url: '/userFundAccounts/updateUserFundAccounts',
    method: 'put',
    data
  })
}

// @Tags UserFundAccounts
// @Summary 用id查询userFundAccounts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserFundAccounts true "用id查询userFundAccounts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userFundAccounts/findUserFundAccounts [get]
export const findUserFundAccounts = (params) => {
  return service({
    url: '/userFundAccounts/findUserFundAccounts',
    method: 'get',
    params
  })
}

// @Tags UserFundAccounts
// @Summary 分页获取userFundAccounts表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userFundAccounts表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userFundAccounts/getUserFundAccountsList [get]
export const getUserFundAccountsList = (params) => {
  return service({
    url: '/userFundAccounts/getUserFundAccountsList',
    method: 'get',
    params
  })
}

// @Tags UserFundAccounts
// @Summary 不需要鉴权的userFundAccounts表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.UserFundAccountsSearch true "分页获取userFundAccounts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userFundAccounts/getUserFundAccountsPublic [get]
export const getUserFundAccountsPublic = () => {
  return service({
    url: '/userFundAccounts/getUserFundAccountsPublic',
    method: 'get',
  })
}
// Recharge 用户充值接口
// @Tags UserFundAccounts
// @Summary 用户充值接口
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /userFundAccounts/recharge [POST]
export const recharge = () => {
  return service({
    url: '/userFundAccounts/recharge',
    method: 'POST'
  })
}
