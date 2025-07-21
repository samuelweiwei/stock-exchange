import service from '@/utils/request'
// @Tags EarnProducts
// @Summary 创建earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnProducts true "创建earnProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /earnProducts/createEarnProducts [post]
export const createEarnProducts = (data) => {
  return service({
    url: '/earnProducts/createEarnProducts',
    method: 'post',
    data
  })
}

// @Tags EarnProducts
// @Summary 删除earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnProducts true "删除earnProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnProducts/deleteEarnProducts [delete]
export const deleteEarnProducts = (params) => {
  return service({
    url: '/earnProducts/deleteEarnProducts',
    method: 'delete',
    params
  })
}

// @Tags EarnProducts
// @Summary 批量删除earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除earnProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /earnProducts/deleteEarnProducts [delete]
export const deleteEarnProductsByIds = (params) => {
  return service({
    url: '/earnProducts/deleteEarnProductsByIds',
    method: 'delete',
    params
  })
}

// @Tags EarnProducts
// @Summary 更新earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EarnProducts true "更新earnProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /earnProducts/updateEarnProducts [put]
export const updateEarnProducts = (data) => {
  return service({
    url: '/earnProducts/updateEarnProducts',
    method: 'put',
    data
  })
}

// @Tags EarnProducts
// @Summary 用id查询earnProducts表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EarnProducts true "用id查询earnProducts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /earnProducts/findEarnProducts [get]
export const findEarnProducts = (params) => {
  return service({
    url: '/earnProducts/findEarnProducts',
    method: 'get',
    params
  })
}

// @Tags EarnProducts
// @Summary 分页获取earnProducts表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取earnProducts表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /earnProducts/getEarnProductsList [get]
export const getEarnProductsList = (params) => {
  return service({
    url: '/earnProducts/getEarnProductsList',
    method: 'get',
    params
  })
}

// @Tags EarnProducts
// @Summary 不需要鉴权的earnProducts表接口
// @accept application/json
// @Produce application/json
// @Param data query earnReq.EarnProductsSearch true "分页获取earnProducts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /earnProducts/getEarnProductsPublic [get]
export const getEarnProductsPublic = () => {
  return service({
    url: '/earnProducts/getEarnProductsPublic',
    method: 'get',
  })
}
