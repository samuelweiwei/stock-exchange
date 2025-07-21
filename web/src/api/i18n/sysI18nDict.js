import service from '@/utils/request'
// @Tags SysI18nDict
// @Summary 创建sysI18nDict表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nDict true "创建sysI18nDict表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysI18nDict/createSysI18nDict [post]
export const createSysI18nDict = (data) => {
  return service({
    url: '/sysI18nDict/createSysI18nDict',
    method: 'post',
    data
  })
}

// @Tags SysI18nDict
// @Summary 删除sysI18nDict表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nDict true "删除sysI18nDict表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysI18nDict/deleteSysI18nDict [delete]
export const deleteSysI18nDict = (params) => {
  return service({
    url: '/sysI18nDict/deleteSysI18nDict',
    method: 'delete',
    params
  })
}

// @Tags SysI18nDict
// @Summary 批量删除sysI18nDict表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysI18nDict表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysI18nDict/deleteSysI18nDict [delete]
export const deleteSysI18nDictByIds = (params) => {
  return service({
    url: '/sysI18nDict/deleteSysI18nDictByIds',
    method: 'delete',
    params
  })
}

// @Tags SysI18nDict
// @Summary 更新sysI18nDict表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysI18nDict true "更新sysI18nDict表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysI18nDict/updateSysI18nDict [put]
export const updateSysI18nDict = (data) => {
  return service({
    url: '/sysI18nDict/updateSysI18nDict',
    method: 'put',
    data
  })
}

// @Tags SysI18nDict
// @Summary 用id查询sysI18nDict表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysI18nDict true "用id查询sysI18nDict表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysI18nDict/findSysI18nDict [get]
export const findSysI18nDict = (params) => {
  return service({
    url: '/sysI18nDict/findSysI18nDict',
    method: 'get',
    params
  })
}

// @Tags SysI18nDict
// @Summary 分页获取sysI18nDict表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysI18nDict表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysI18nDict/getSysI18nDictList [get]
export const getSysI18nDictList = (params) => {
  return service({
    url: '/sysI18nDict/getSysI18nDictList',
    method: 'get',
    params
  })
}

// @Tags SysI18nDict
// @Summary 不需要鉴权的sysI18nDict表接口
// @accept application/json
// @Produce application/json
// @Param data query i18nReq.SysI18nDictSearch true "分页获取sysI18nDict表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysI18nDict/getSysI18nDictPublic [get]
export const getSysI18nDictPublic = () => {
  return service({
    url: '/sysI18nDict/getSysI18nDictPublic',
    method: 'get',
  })
}
