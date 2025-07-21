import service from '@/utils/request'
// @Tags RechargeRecords
// @Summary 创建rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeRecords true "创建rechargeRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rechargeRecords/createRechargeRecords [post]
export const createRechargeRecords = (data) => {
  return service({
    url: '/rechargeRecords/createRechargeRecords',
    method: 'post',
    data
  })
}

// @Tags RechargeRecords
// @Summary 删除rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeRecords true "删除rechargeRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeRecords/deleteRechargeRecords [delete]
export const deleteRechargeRecords = (params) => {
  return service({
    url: '/rechargeRecords/deleteRechargeRecords',
    method: 'delete',
    params
  })
}

// @Tags RechargeRecords
// @Summary 批量删除rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除rechargeRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rechargeRecords/deleteRechargeRecords [delete]
export const deleteRechargeRecordsByIds = (params) => {
  return service({
    url: '/rechargeRecords/deleteRechargeRecordsByIds',
    method: 'delete',
    params
  })
}

// @Tags RechargeRecords
// @Summary 更新rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RechargeRecords true "更新rechargeRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rechargeRecords/updateRechargeRecords [put]
export const updateRechargeRecords = (data) => {
  return service({
    url: '/rechargeRecords/updateRechargeRecords',
    method: 'put',
    data
  })
}

// @Tags RechargeRecords
// @Summary 用id查询rechargeRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RechargeRecords true "用id查询rechargeRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rechargeRecords/findRechargeRecords [get]
export const findRechargeRecords = (params) => {
  return service({
    url: '/rechargeRecords/findRechargeRecords',
    method: 'get',
    params
  })
}

// @Tags RechargeRecords
// @Summary 分页获取rechargeRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取rechargeRecords表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rechargeRecords/getRechargeRecordsList [get]
export const getRechargeRecordsList = (params) => {
  return service({
    url: '/rechargeRecords/getRechargeRecordsList',
    method: 'get',
    params
  })
}

// @Tags RechargeRecords
// @Summary 不需要鉴权的rechargeRecords表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.RechargeRecordsSearch true "分页获取rechargeRecords表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rechargeRecords/getRechargeRecordsPublic [get]
export const getRechargeRecordsPublic = () => {
  return service({
    url: '/rechargeRecords/getRechargeRecordsPublic',
    method: 'get',
  })
}
