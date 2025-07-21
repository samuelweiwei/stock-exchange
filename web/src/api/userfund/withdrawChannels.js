import service from '@/utils/request'
// @Tags WithdrawChannels
// @Summary 创建withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WithdrawChannels true "创建withdrawChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /withdrawChannels/createWithdrawChannels [post]
export const createWithdrawChannels = (data) => {
  return service({
    url: '/withdrawChannels/createWithdrawChannels',
    method: 'post',
    data
  })
}

// @Tags WithdrawChannels
// @Summary 删除withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WithdrawChannels true "删除withdrawChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdrawChannels/deleteWithdrawChannels [delete]
export const deleteWithdrawChannels = (params) => {
  return service({
    url: '/withdrawChannels/deleteWithdrawChannels',
    method: 'delete',
    params
  })
}

// @Tags WithdrawChannels
// @Summary 批量删除withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除withdrawChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdrawChannels/deleteWithdrawChannels [delete]
export const deleteWithdrawChannelsByIds = (params) => {
  return service({
    url: '/withdrawChannels/deleteWithdrawChannelsByIds',
    method: 'delete',
    params
  })
}

// @Tags WithdrawChannels
// @Summary 更新withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WithdrawChannels true "更新withdrawChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /withdrawChannels/updateWithdrawChannels [put]
export const updateWithdrawChannels = (data) => {
  return service({
    url: '/withdrawChannels/updateWithdrawChannels',
    method: 'put',
    data
  })
}

// @Tags WithdrawChannels
// @Summary 用id查询withdrawChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WithdrawChannels true "用id查询withdrawChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /withdrawChannels/findWithdrawChannels [get]
export const findWithdrawChannels = (params) => {
  return service({
    url: '/withdrawChannels/findWithdrawChannels',
    method: 'get',
    params
  })
}

// @Tags WithdrawChannels
// @Summary 分页获取withdrawChannels表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取withdrawChannels表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /withdrawChannels/getWithdrawChannelsList [get]
export const getWithdrawChannelsList = (params) => {
  return service({
    url: '/withdrawChannels/getWithdrawChannelsList',
    method: 'get',
    params
  })
}

// @Tags WithdrawChannels
// @Summary 不需要鉴权的withdrawChannels表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.WithdrawChannelsSearch true "分页获取withdrawChannels表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /withdrawChannels/getWithdrawChannelsPublic [get]
export const getWithdrawChannelsPublic = () => {
  return service({
    url: '/withdrawChannels/getWithdrawChannelsPublic',
    method: 'get',
  })
}
