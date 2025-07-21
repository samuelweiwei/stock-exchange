/**
* @Author: Jackey
* @Date: 12/29/24 8:17 pm
 */

package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"net/http"
	"strings"
)

type IpInfoResponse struct {
	Status  string `json:"status"`
	Country string `json:"country"`
	City    string `json:"city"`
}

const (
	ResponseSuccessString = "success"

	IgnoreIpStrings = "#127.0.0.1#"
)

func FetchIpToRegion(ctx context.Context) (err error) {
	var (
		//1.获取需要查询的ip
		ipString, _ = utils.GetIpToRegionListFirstIp()
	)
	//2.判断ip是否有效
	if len(ipString) == 0 {
		return
	} else if strings.Contains(IgnoreIpStrings, ipString) {
		return
	}

	//3.先判断缓存是否已存在ip信息
	region, _ := utils.GetRegionByIp(ipString)
	if len(region) > 0 {
		return
	}

	//4.请求三方获取ip信息。如果失败就将ip重新写入待查询ip缓存
	region, err = startQuery(ipString)
	if err != nil {
		_ = utils.AddIpStringInIpToRegionList(ipString)
		utils.SendMsgToTg(fmt.Sprintf("【获取ip地址信息失败】\n[ip]：%s\n[失败原因]：%s",
			ipString, err.Error()))
	}

	//5.如果有查询到ip信息就写入到缓存，并且更新用户的登录记录
	if len(region) > 0 {
		_ = utils.SetRegionInfo(ipString, region)
		//更新用户登录记录的ip信息
		err = service.ServiceGroupApp.UserServiceGroup.FrontendUserLoginLogService.UpdateFrontendUserLoginLogRegionByIp(ipString, region)
		if err != nil {
			utils.SendMsgToTg(fmt.Sprintf("【更新用户登录记录ip信息失败】\n[ip]：%s \n[信息]：%s\n[失败原因]：%s",
				ipString, region, err.Error()))
		}
	}
	return
}

func startQuery(ipString string) (regionString string, err error) {

	//创建必要的参数，发送请求
	var (
		urlString = fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ipString)
		resp      *http.Response
		ipInfo    IpInfoResponse
	)
	resp, err = http.Get(urlString)
	if err != nil {
		err = errors.New("发送请求失败->" + err.Error())
		return
	}

	//解析响应数据
	if err = json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		err = errors.New("解析响应失败->" + err.Error())
		return
	}

	//如果响应数据成功，就更新数据
	if ipInfo.Status == ResponseSuccessString {
		regionString = fmt.Sprintf("%s %s", ipInfo.Country, ipInfo.City)
	} else {
		err = errors.New(fmt.Sprintf("返回数据错误->%v", ipInfo))
	}
	return
}
