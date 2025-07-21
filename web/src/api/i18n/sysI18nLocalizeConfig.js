import service from '@/utils/request'
// @Tags SysI18nLocalizeConfig
// @Summary 创建sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nLocalizeConfig true "创建sysI18nLocalizeConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysI18nLocalizeConfig/createSysI18nLocalizeConfig [post]
export const createSysI18nLocalizeConfig = (data) => {
  return service({
    url: '/sysI18nLocalizeConfig/createSysI18nLocalizeConfig',
    method: 'post',
    data
  })
}

// @Tags SysI18nLocalizeConfig
// @Summary 删除sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nLocalizeConfig true "删除sysI18nLocalizeConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysI18nLocalizeConfig/deleteSysI18nLocalizeConfig [delete]
export const deleteSysI18nLocalizeConfig = (params) => {
  return service({
    url: '/sysI18nLocalizeConfig/deleteSysI18nLocalizeConfig',
    method: 'delete',
    params
  })
}

// @Tags SysI18nLocalizeConfig
// @Summary 批量删除sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysI18nLocalizeConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysI18nLocalizeConfig/deleteSysI18nLocalizeConfig [delete]
export const deleteSysI18nLocalizeConfigByIds = (params) => {
  return service({
    url: '/sysI18nLocalizeConfig/deleteSysI18nLocalizeConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags SysI18nLocalizeConfig
// @Summary 更新sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nLocalizeConfig true "更新sysI18nLocalizeConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysI18nLocalizeConfig/updateSysI18nLocalizeConfig [put]
export const updateSysI18nLocalizeConfig = (data) => {
  return service({
    url: '/sysI18nLocalizeConfig/updateSysI18nLocalizeConfig',
    method: 'put',
    data
  })
}

// @Tags SysI18nLocalizeConfig
// @Summary 用id查询sysI18nLocalizeConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysI18nLocalizeConfig true "用id查询sysI18nLocalizeConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysI18nLocalizeConfig/findSysI18nLocalizeConfig [get]
export const findSysI18nLocalizeConfig = (params) => {
  return service({
    url: '/sysI18nLocalizeConfig/findSysI18nLocalizeConfig',
    method: 'get',
    params
  })
}

// @Tags SysI18nLocalizeConfig
// @Summary 分页获取sysI18nLocalizeConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysI18nLocalizeConfig表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysI18nLocalizeConfig/getSysI18nLocalizeConfigList [get]
export const getSysI18nLocalizeConfigList = (params) => {
  return service({
    url: '/sysI18nLocalizeConfig/getSysI18nLocalizeConfigList',
    method: 'get',
    params
  })
}

// @Tags SysI18nLocalizeConfig
// @Summary 不需要鉴权的sysI18nLocalizeConfig表接口
// @accept application/json
// @Produce application/json
// @Param data query i18nReq.SysI18nLocalizeConfigSearch true "分页获取sysI18nLocalizeConfig表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysI18nLocalizeConfig/getSysI18nLocalizeConfigPublic [get]
export const getSysI18nLocalizeConfigPublic = () => {
  return service({
    url: '/sysI18nLocalizeConfig/getSysI18nLocalizeConfigPublic',
    method: 'get',
  })
}
