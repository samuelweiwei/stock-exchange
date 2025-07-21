import service from '@/utils/request'
// @Tags SysI18nDictCategory
// @Summary 创建sysI18nDictCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nDictCategory true "创建sysI18nDictCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysI18nDictCategory/createSysI18nDictCategory [post]
export const createSysI18nDictCategory = (data) => {
  return service({
    url: '/sysI18nDictCategory/createSysI18nDictCategory',
    method: 'post',
    data
  })
}

// @Tags SysI18nDictCategory
// @Summary 删除sysI18nDictCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nDictCategory true "删除sysI18nDictCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysI18nDictCategory/deleteSysI18nDictCategory [delete]
export const deleteSysI18nDictCategory = (params) => {
  return service({
    url: '/sysI18nDictCategory/deleteSysI18nDictCategory',
    method: 'delete',
    params
  })
}

// @Tags SysI18nDictCategory
// @Summary 批量删除sysI18nDictCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysI18nDictCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysI18nDictCategory/deleteSysI18nDictCategory [delete]
export const deleteSysI18nDictCategoryByIds = (params) => {
  return service({
    url: '/sysI18nDictCategory/deleteSysI18nDictCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags SysI18nDictCategory
// @Summary 更新sysI18nDictCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nDictCategory true "更新sysI18nDictCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysI18nDictCategory/updateSysI18nDictCategory [put]
export const updateSysI18nDictCategory = (data) => {
  return service({
    url: '/sysI18nDictCategory/updateSysI18nDictCategory',
    method: 'put',
    data
  })
}

// @Tags SysI18nDictCategory
// @Summary 用id查询sysI18nDictCategory表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysI18nDictCategory true "用id查询sysI18nDictCategory表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysI18nDictCategory/findSysI18nDictCategory [get]
export const findSysI18nDictCategory = (params) => {
  return service({
    url: '/sysI18nDictCategory/findSysI18nDictCategory',
    method: 'get',
    params
  })
}

// @Tags SysI18nDictCategory
// @Summary 分页获取sysI18nDictCategory表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysI18nDictCategory表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysI18nDictCategory/getSysI18nDictCategoryList [get]
export const getSysI18nDictCategoryList = (params) => {
  return service({
    url: '/sysI18nDictCategory/getSysI18nDictCategoryList',
    method: 'get',
    params
  })
}

// @Tags SysI18nDictCategory
// @Summary 不需要鉴权的sysI18nDictCategory表接口
// @accept application/json
// @Produce application/json
// @Param data query i18nReq.SysI18nDictCategorySearch true "分页获取sysI18nDictCategory表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysI18nDictCategory/getSysI18nDictCategoryPublic [get]
export const getSysI18nDictCategoryPublic = () => {
  return service({
    url: '/sysI18nDictCategory/getSysI18nDictCategoryPublic',
    method: 'get',
  })
}
