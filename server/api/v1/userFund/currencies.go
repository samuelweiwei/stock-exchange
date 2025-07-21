package userFund

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"strconv"
)

type CurrenciesApi struct{}

// CreateCurrencies 创建currencies表
// @Tags Currencies
// @Summary 创建currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.Currencies true "创建currencies表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /currencies/create [post]
func (currenciesApi *CurrenciesApi) CreateCurrencies(c *gin.Context) {
	var currencies userfund.Currencies
	err := c.ShouldBindJSON(&currencies)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = currenciesService.CreateCurrencies(&currencies)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCurrencies 删除currencies表
// @Tags Currencies
// @Summary 删除currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.Currencies true "删除currencies表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /currencies/delete [delete]
func (currenciesApi *CurrenciesApi) DeleteCurrencies(c *gin.Context) {
	id := c.Query("id")
	err := currenciesService.DeleteCurrencies(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCurrenciesByIds 批量删除currencies表
// @Tags Currencies
// @Summary 批量删除currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /currencies/deleteByIds [delete]
func (currenciesApi *CurrenciesApi) DeleteCurrenciesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := currenciesService.DeleteCurrenciesByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCurrencies 更新currencies表
// @Tags Currencies
// @Summary 更新currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userfund.Currencies true "更新currencies表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /currencies/update [put]
func (currenciesApi *CurrenciesApi) UpdateCurrencies(c *gin.Context) {
	var currencies userfund.Currencies
	err := c.ShouldBindJSON(&currencies)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = currenciesService.UpdateCurrencies(currencies)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCurrencies 用id查询currencies表
// @Tags Currencies
// @Summary 用id查询currencies表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfund.Currencies true "用id查询currencies表"
// @Success 200 {object} response.Response{data=userfund.Currencies,msg=string} "查询成功"
// @Router /currencies/detail/{id} [get]
func (currenciesApi *CurrenciesApi) FindCurrencies(c *gin.Context) {
	id := c.Param("id")
	recurrencies, err := currenciesService.GetCurrencies(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recurrencies, c)
}

// GetCurrenciesList 分页获取currencies表列表
// @Tags Currencies
// @Summary 分页获取currencies表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userfundReq.CurrenciesSearch true "页码, 每页大小, 搜索条件"
// @Param currency query string false "货币代码"
// @Param coinType query int false "货币类型：1-数字货币 2-法币"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /currencies/list [get]
func (currenciesApi *CurrenciesApi) GetCurrenciesList(c *gin.Context) {
	var pageInfo userfundReq.CurrenciesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := currenciesService.GetCurrenciesInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	// 遍历记录列表
	for i := range list {
		record := &list[i] // 使用指针
		// 检查 locker 属性是否等于当前登录用户
		record.CreateAtInt = record.CreatedAt.UnixMilli()
		record.UpdatedAtInt = record.UpdatedAt.UnixMilli()
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetCurrenciesPublic 获取货币列表
// @Tags Currencies
// @Summary 前端-获取货币列表
// @accept application/json
// @Produce application/json
// @Param coinType query int false "货币类型：1数字货币 2法币"
// @Success 200 {object} response.Response{data=[]userfund.Currencies,msg=string} "获取成功"
// @Router /currencies/frontend/list [get]
func (currenciesApi *CurrenciesApi) GetCurrenciesPublic(c *gin.Context) {
	var searchInfo userfundReq.CurrenciesSearch
	err := c.ShouldBindQuery(&searchInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, err := currenciesService.GetCurrenciesPublic(searchInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	// 假设你有一个 Redis 客户端实例 redisClient
	for i := range list {
		// 判断是否是数字货币类型，
		currency := &list[i]
		if *currency.CoinType == 1 && currency.Currency != "USDT" {
			redisKey := fmt.Sprintf("symbol:crypto:%s", currency.Currency+"/USD")
			// 从Redis获取价格
			if price, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result(); err == nil {
				if priceFloat, err := strconv.ParseFloat(price, 64); err == nil {
					currency.PriceUsdt = decimal.NewFromFloat(priceFloat)
				}
			}
		}
	}
	response.OkWithData(list, c)
}
