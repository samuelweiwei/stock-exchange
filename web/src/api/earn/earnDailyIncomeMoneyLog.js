import service from '@/utils/request'
// @Tags EarnDailyIncomeMoneyLog
// @Summary 创建earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnDailyIncomeMoneyLog true "创建earnDailyIncomeMoneyLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /earnDailyIncomeMoneyLog/createEarnDailyIncomeMoneyLog [post]
export const createEarnDailyIncomeMoneyLog = (data) => {
  return service({
    url: '/earnDailyIncomeMoneyLog/createEarnDailyIncomeMoneyLog',
    method: 'post',
    data
  })
}

// @Tags EarnDailyIncomeMoneyLog
// @Summary 删除earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnDailyIncomeMoneyLog true "删除earnDailyIncomeMoneyLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnDailyIncomeMoneyLog/deleteEarnDailyIncomeMoneyLog [delete]
export const deleteEarnDailyIncomeMoneyLog = (params) => {
  return service({
    url: '/earnDailyIncomeMoneyLog/deleteEarnDailyIncomeMoneyLog',
    method: 'delete',
    params
  })
}

// @Tags EarnDailyIncomeMoneyLog
// @Summary 批量删除earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除earnDailyIncomeMoneyLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnDailyIncomeMoneyLog/deleteEarnDailyIncomeMoneyLog [delete]
export const deleteEarnDailyIncomeMoneyLogByIds = (params) => {
  return service({
    url: '/earnDailyIncomeMoneyLog/deleteEarnDailyIncomeMoneyLogByIds',
    method: 'delete',
    params
  })
}

// @Tags EarnDailyIncomeMoneyLog
// @Summary 更新earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnDailyIncomeMoneyLog true "更新earnDailyIncomeMoneyLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /earnDailyIncomeMoneyLog/updateEarnDailyIncomeMoneyLog [put]
export const updateEarnDailyIncomeMoneyLog = (data) => {
  return service({
    url: '/earnDailyIncomeMoneyLog/updateEarnDailyIncomeMoneyLog',
    method: 'put',
    data
  })
}

// @Tags EarnDailyIncomeMoneyLog
// @Summary 用id查询earnDailyIncomeMoneyLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EarnDailyIncomeMoneyLog true "用id查询earnDailyIncomeMoneyLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /earnDailyIncomeMoneyLog/findEarnDailyIncomeMoneyLog [get]
export const findEarnDailyIncomeMoneyLog = (params) => {
  return service({
    url: '/earnDailyIncomeMoneyLog/findEarnDailyIncomeMoneyLog',
    method: 'get',
    params
  })
}

// @Tags EarnDailyIncomeMoneyLog
// @Summary 分页获取earnDailyIncomeMoneyLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取earnDailyIncomeMoneyLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /earnDailyIncomeMoneyLog/getEarnDailyIncomeMoneyLogList [get]
export const getEarnDailyIncomeMoneyLogList = (params) => {
  return service({
    url: '/earnDailyIncomeMoneyLog/getEarnDailyIncomeMoneyLogList',
    method: 'get',
    params
  })
}

// @Tags EarnDailyIncomeMoneyLog
// @Summary 不需要鉴权的earnDailyIncomeMoneyLog表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnDailyIncomeMoneyLogSearch true "分页获取earnDailyIncomeMoneyLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnDailyIncomeMoneyLog/getEarnDailyIncomeMoneyLogPublic [get]
export const getEarnDailyIncomeMoneyLogPublic = () => {
  return service({
    url: '/earnDailyIncomeMoneyLog/getEarnDailyIncomeMoneyLogPublic',
    method: 'get',
  })
}
