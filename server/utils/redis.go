package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/redis/go-redis/v9"
)

// GetIpToRegionListFirstIp 获取待查询的ip
func GetIpToRegionListFirstIp() (ipString string, err error) {
	ipString, err = global.GVA_REDIS.RPop(context.Background(), constants.RedisKeyIpToRegionList).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		SendMsgToTgAsync(fmt.Sprintf("【获取Redis待查询ip出错】\n[redisKey]：%s\n[错误原因]：%s",
			constants.RedisKeyIpToRegionList, err.Error()))
	}
	return
}

// AddIpStringInIpToRegionList 添加待查询的ip到缓存
func AddIpStringInIpToRegionList(ipString string) (err error) {
	err = global.GVA_REDIS.RPush(context.Background(), constants.RedisKeyIpToRegionList, ipString).Err()
	if err != nil {
		SendMsgToTgAsync(fmt.Sprintf("【添加Redis待查询ip出错】\n[redisKey]：%s\n[错误原因]：%s",
			constants.RedisKeyIpToRegionList, err.Error()))
	}
	return
}

// GetRegionByIp 通过ip获取地区缓存
func GetRegionByIp(ipString string) (region string, err error) {
	region, err = global.GVA_REDIS.HGet(context.Background(), constants.RedisKeyIpRegionInfo, ipString).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		SendMsgToTgAsync(fmt.Sprintf("【获取Redis的ip地区信息出错】\n[redisKey]：%s\n[错误原因]：%s",
			constants.RedisKeyIpRegionInfo+"."+ipString, err.Error()))
	}
	return
}

// SetRegionInfo 设置ip-region缓存
func SetRegionInfo(ipString, region string) (err error) {

	err = global.GVA_REDIS.HSet(context.Background(), constants.RedisKeyIpRegionInfo, ipString, region).Err()
	if err != nil {
		SendMsgToTgAsync(fmt.Sprintf("【写入Redis的ip地区信息出错】\n[redisKeyValue]：%s\n[错误原因]：%s",
			constants.RedisKeyIpRegionInfo+"."+ipString+":"+region, err.Error()))
	}
	return
}
