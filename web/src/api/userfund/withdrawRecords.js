import service from '@/utils/request'
// @Tags WithdrawRecords
// @Summary 创建withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WithdrawRecords true "创建withdrawRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /withdrawRecords/createWithdrawRecords [post]
export const createWithdrawRecords = (data) => {
  return service({
    url: '/withdrawRecords/createWithdrawRecords',
    method: 'post',
    data
  })
}

// @Tags WithdrawRecords
// @Summary 删除withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WithdrawRecords true "删除withdrawRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdrawRecords/deleteWithdrawRecords [delete]
export const deleteWithdrawRecords = (params) => {
  return service({
    url: '/withdrawRecords/deleteWithdrawRecords',
    method: 'delete',
    params
  })
}

// @Tags WithdrawRecords
// @Summary 批量删除withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除withdrawRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /withdrawRecords/deleteWithdrawRecords [delete]
export const deleteWithdrawRecordsByIds = (params) => {
  return service({
    url: '/withdrawRecords/deleteWithdrawRecordsByIds',
    method: 'delete',
    params
  })
}

// @Tags WithdrawRecords
// @Summary 更新withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WithdrawRecords true "更新withdrawRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /withdrawRecords/updateWithdrawRecords [put]
export const updateWithdrawRecords = (data) => {
  return service({
    url: '/withdrawRecords/updateWithdrawRecords',
    method: 'put',
    data
  })
}

// @Tags WithdrawRecords
// @Summary 用id查询withdrawRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WithdrawRecords true "用id查询withdrawRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /withdrawRecords/findWithdrawRecords [get]
export const findWithdrawRecords = (params) => {
  return service({
    url: '/withdrawRecords/findWithdrawRecords',
    method: 'get',
    params
  })
}

// @Tags WithdrawRecords
// @Summary 分页获取withdrawRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取withdrawRecords表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /withdrawRecords/getWithdrawRecordsList [get]
export const getWithdrawRecordsList = (params) => {
  return service({
    url: '/withdrawRecords/getWithdrawRecordsList',
    method: 'get',
    params
  })
}

// @Tags WithdrawRecords
// @Summary 不需要鉴权的withdrawRecords表接口
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.WithdrawRecordsSearch true "分页获取withdrawRecords表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /withdrawRecords/getWithdrawRecordsPublic [get]
export const getWithdrawRecordsPublic = () => {
  return service({
    url: '/withdrawRecords/getWithdrawRecordsPublic',
    method: 'get',
  })
}
