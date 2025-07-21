import service from '@/utils/request'
// @Tags NationalityInfo
// @Summary 创建nationalityInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NationalityInfo true "创建nationalityInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /nationalityInfo/createNationalityInfo [post]
export const createNationalityInfo = (data) => {
  return service({
    url: '/nationalityInfo/createNationalityInfo',
    method: 'post',
    data
  })
}

// @Tags NationalityInfo
// @Summary 删除nationalityInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NationalityInfo true "删除nationalityInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nationalityInfo/deleteNationalityInfo [delete]
export const deleteNationalityInfo = (params) => {
  return service({
    url: '/nationalityInfo/deleteNationalityInfo',
    method: 'delete',
    params
  })
}

// @Tags NationalityInfo
// @Summary 批量删除nationalityInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除nationalityInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nationalityInfo/deleteNationalityInfo [delete]
export const deleteNationalityInfoByIds = (params) => {
  return service({
    url: '/nationalityInfo/deleteNationalityInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags NationalityInfo
// @Summary 更新nationalityInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NationalityInfo true "更新nationalityInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /nationalityInfo/updateNationalityInfo [put]
export const updateNationalityInfo = (data) => {
  return service({
    url: '/nationalityInfo/updateNationalityInfo',
    method: 'put',
    data
  })
}

// @Tags NationalityInfo
// @Summary 用id查询nationalityInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NationalityInfo true "用id查询nationalityInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /nationalityInfo/findNationalityInfo [get]
export const findNationalityInfo = (params) => {
  return service({
    url: '/nationalityInfo/findNationalityInfo',
    method: 'get',
    params
  })
}

// @Tags NationalityInfo
// @Summary 分页获取nationalityInfo表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取nationalityInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nationalityInfo/getNationalityInfoList [get]
export const getNationalityInfoList = (params) => {
  return service({
    url: '/nationalityInfo/getNationalityInfoList',
    method: 'get',
    params
  })
}

// @Tags NationalityInfo
// @Summary 不需要鉴权的nationalityInfo表接口
// @accept application/json
// @Produce application/json
// @Param data query userReq.NationalityInfoSearch true "分页获取nationalityInfo表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /nationalityInfo/getNationalityInfoPublic [get]
export const getNationalityInfoPublic = () => {
  return service({
    url: '/nationalityInfo/getNationalityInfoPublic',
    method: 'get',
  })
}
