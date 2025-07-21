import service from '@/utils/request'
// @Tags RechargeChannels
// @Summary 创建rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeChannels true "创建rechargeChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rechargeChannels/createRechargeChannels [post]
export const createRechargeChannels = (data) => {
  return service({
    url: '/rechargeChannels/createRechargeChannels',
    method: 'post',
    data
  })
}

// @Tags RechargeChannels
// @Summary 删除rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeChannels true "删除rechargeChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeChannels/deleteRechargeChannels [delete]
export const deleteRechargeChannels = (params) => {
  return service({
    url: '/rechargeChannels/deleteRechargeChannels',
    method: 'delete',
    params
  })
}

// @Tags RechargeChannels
// @Summary 批量删除rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除rechargeChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeChannels/deleteRechargeChannels [delete]
export const deleteRechargeChannelsByIds = (params) => {
  return service({
    url: '/rechargeChannels/deleteRechargeChannelsByIds',
    method: 'delete',
    params
  })
}

// @Tags RechargeChannels
// @Summary 更新rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeChannels true "更新rechargeChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rechargeChannels/updateRechargeChannels [put]
export const updateRechargeChannels = (data) => {
  return service({
    url: '/rechargeChannels/updateRechargeChannels',
    method: 'put',
    data
  })
}

// @Tags RechargeChannels
// @Summary 用id查询rechargeChannels表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RechargeChannels true "用id查询rechargeChannels表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rechargeChannels/findRechargeChannels [get]
export const findRechargeChannels = (params) => {
  return service({
    url: '/rechargeChannels/findRechargeChannels',
    method: 'get',
    params
  })
}

// @Tags RechargeChannels
// @Summary 分页获取rechargeChannels表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取rechargeChannels表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rechargeChannels/getRechargeChannelsList [get]
export const getRechargeChannelsList = (params) => {
  return service({
    url: '/rechargeChannels/getRechargeChannelsList',
    method: 'get',
    params
  })
}

// @Tags RechargeChannels
// @Summary 不需要鉴权的rechargeChannels表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.RechargeChannelsSearch true "分页获取rechargeChannels表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rechargeChannels/getRechargeChannelsPublic [get]
export const getRechargeChannelsPublic = () => {
  return service({
    url: '/rechargeChannels/getRechargeChannelsPublic',
    method: 'get',
  })
}
// Paynotify 支付回调接口
// @Tags RechargeChannels
// @Summary 支付回调接口
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /rechargeChannels/paynotify [POST]
export const paynotify = () => {
  return service({
    url: '/rechargeChannels/paynotify',
    method: 'POST'
  })
}
