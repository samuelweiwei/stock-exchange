import service from '@/utils/request'
// @Tags SafetyEducationCategories
// @Summary 创建safetyEducationCategories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SafetyEducationCategories true "创建safetyEducationCategories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /safetyEducationCategories/createSafetyEducationCategories [post]
export const createSafetyEducationCategories = (data) => {
  return service({
    url: '/safetyEducationCategories/createSafetyEducationCategories',
    method: 'post',
    data
  })
}

// @Tags SafetyEducationCategories
// @Summary 删除safetyEducationCategories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SafetyEducationCategories true "删除safetyEducationCategories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /safetyEducationCategories/deleteSafetyEducationCategories [delete]
export const deleteSafetyEducationCategories = (params) => {
  return service({
    url: '/safetyEducationCategories/deleteSafetyEducationCategories',
    method: 'delete',
    params
  })
}

// @Tags SafetyEducationCategories
// @Summary 批量删除safetyEducationCategories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除safetyEducationCategories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /safetyEducationCategories/deleteSafetyEducationCategories [delete]
export const deleteSafetyEducationCategoriesByIds = (params) => {
  return service({
    url: '/safetyEducationCategories/deleteSafetyEducationCategoriesByIds',
    method: 'delete',
    params
  })
}

// @Tags SafetyEducationCategories
// @Summary 更新safetyEducationCategories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SafetyEducationCategories true "更新safetyEducationCategories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /safetyEducationCategories/updateSafetyEducationCategories [put]
export const updateSafetyEducationCategories = (data) => {
  return service({
    url: '/safetyEducationCategories/updateSafetyEducationCategories',
    method: 'put',
    data
  })
}

// @Tags SafetyEducationCategories
// @Summary 用id查询safetyEducationCategories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SafetyEducationCategories true "用id查询safetyEducationCategories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /safetyEducationCategories/findSafetyEducationCategories [get]
export const findSafetyEducationCategories = (params) => {
  return service({
    url: '/safetyEducationCategories/findSafetyEducationCategories',
    method: 'get',
    params
  })
}

// @Tags SafetyEducationCategories
// @Summary 分页获取safetyEducationCategories表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取safetyEducationCategories表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /safetyEducationCategories/getSafetyEducationCategoriesList [get]
export const getSafetyEducationCategoriesList = (params) => {
  return service({
    url: '/safetyEducationCategories/getSafetyEducationCategoriesList',
    method: 'get',
    params
  })
}

// @Tags SafetyEducationCategories
// @Summary 不需要鉴权的safetyEducationCategories表接口
// @accept application/json
// @Produce application/json
// @Param data query stockReq.SafetyEducationCategoriesSearch true "分页获取safetyEducationCategories表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /safetyEducationCategories/getSafetyEducationCategoriesPublic [get]
export const getSafetyEducationCategoriesPublic = () => {
  return service({
    url: '/safetyEducationCategories/getSafetyEducationCategoriesPublic',
    method: 'get',
  })
}
