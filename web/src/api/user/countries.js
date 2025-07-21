import service from '@/utils/request'
// @Tags Countries
// @Summary 创建countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Countries true "创建countries表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /countries/createCountries [post]
export const createCountries = (data) => {
  return service({
    url: '/countries/createCountries',
    method: 'post',
    data
  })
}

// @Tags Countries
// @Summary 删除countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Countries true "删除countries表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /countries/deleteCountries [delete]
export const deleteCountries = (params) => {
  return service({
    url: '/countries/deleteCountries',
    method: 'delete',
    params
  })
}

// @Tags Countries
// @Summary 批量删除countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除countries表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /countries/deleteCountries [delete]
export const deleteCountriesByIds = (params) => {
  return service({
    url: '/countries/deleteCountriesByIds',
    method: 'delete',
    params
  })
}

// @Tags Countries
// @Summary 更新countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Countries true "更新countries表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /countries/updateCountries [put]
export const updateCountries = (data) => {
  return service({
    url: '/countries/updateCountries',
    method: 'put',
    data
  })
}

// @Tags Countries
// @Summary 用id查询countries表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Countries true "用id查询countries表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /countries/findCountries [get]
export const findCountries = (params) => {
  return service({
    url: '/countries/findCountries',
    method: 'get',
    params
  })
}

// @Tags Countries
// @Summary 分页获取countries表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取countries表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /countries/getCountriesList [get]
export const getCountriesList = (params) => {
  return service({
    url: '/countries/getCountriesList',
    method: 'get',
    params
  })
}

// @Tags Countries
// @Summary 不需要鉴权的countries表接口
// @accept application/json
// @Produce application/json
// @Param data query userReq.CountriesSearch true "分页获取countries表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /countries/getCountriesPublic [get]
export const getCountriesPublic = () => {
  return service({
    url: '/countries/getCountriesPublic',
    method: 'get',
  })
}
