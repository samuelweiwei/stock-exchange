import service from '@/utils/request'
// @Tags EarnSubscribeLog
// @Summary 创建earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnSubscribeLog true "创建earnSubscribeLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /earnSubscribeLog/createEarnSubscribeLog [post]
export const createEarnSubscribeLog = (data) => {
  return service({
    url: '/earnSubscribeLog/createEarnSubscribeLog',
    method: 'post',
    data
  })
}

// @Tags EarnSubscribeLog
// @Summary 删除earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnSubscribeLog true "删除earnSubscribeLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnSubscribeLog/deleteEarnSubscribeLog [delete]
export const deleteEarnSubscribeLog = (params) => {
  return service({
    url: '/earnSubscribeLog/deleteEarnSubscribeLog',
    method: 'delete',
    params
  })
}

// @Tags EarnSubscribeLog
// @Summary 批量删除earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除earnSubscribeLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnSubscribeLog/deleteEarnSubscribeLog [delete]
export const deleteEarnSubscribeLogByIds = (params) => {
  return service({
    url: '/earnSubscribeLog/deleteEarnSubscribeLogByIds',
    method: 'delete',
    params
  })
}

// @Tags EarnSubscribeLog
// @Summary 更新earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnSubscribeLog true "更新earnSubscribeLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /earnSubscribeLog/updateEarnSubscribeLog [put]
export const updateEarnSubscribeLog = (data) => {
  return service({
    url: '/earnSubscribeLog/updateEarnSubscribeLog',
    method: 'put',
    data
  })
}

// @Tags EarnSubscribeLog
// @Summary 用id查询earnSubscribeLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EarnSubscribeLog true "用id查询earnSubscribeLog表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /earnSubscribeLog/findEarnSubscribeLog [get]
export const findEarnSubscribeLog = (params) => {
  return service({
    url: '/earnSubscribeLog/findEarnSubscribeLog',
    method: 'get',
    params
  })
}

// @Tags EarnSubscribeLog
// @Summary 分页获取earnSubscribeLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取earnSubscribeLog表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /earnSubscribeLog/getEarnSubscribeLogList [get]
export const getEarnSubscribeLogList = (params) => {
  return service({
    url: '/earnSubscribeLog/getEarnSubscribeLogList',
    method: 'get',
    params
  })
}

// @Tags EarnSubscribeLog
// @Summary 不需要鉴权的earnSubscribeLog表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnSubscribeLogSearch true "分页获取earnSubscribeLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnSubscribeLog/getEarnSubscribeLogPublic [get]
export const getEarnSubscribeLogPublic = () => {
  return service({
    url: '/earnSubscribeLog/getEarnSubscribeLogPublic',
    method: 'get',
  })
}
