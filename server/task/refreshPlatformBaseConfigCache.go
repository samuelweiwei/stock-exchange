package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

// RefreshPlatformBaseConfigCache 刷新平台基础设置缓存
func RefreshPlatformBaseConfigCache(db *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("Recover refresh platform base config cache err: %v", err))
		}
	}()
	var systemConfig system.SystemConfig
	err := db.Model(&systemConfig).First(&systemConfig).Error
	if err != nil {
		return err
	}

	mSet := map[string]interface{}{
		constants.RedisKeyPlatformCommissionRate:  systemConfig.Config.PlatformCommissionRate,
		constants.RedisKeyFirstGradeShareRate:     systemConfig.Config.FirstGradeShareRate,
		constants.RedisKeySecondGradeShareRate:    systemConfig.Config.SecondGradeShareRate,
		constants.RedisKeyThirdGradeShareRate:     systemConfig.Config.ThirdGradeShareRate,
		constants.RedisKeyWithdrawCommissionType:  systemConfig.Config.WithdrawCommissionType,
		constants.RedisKeyWithdrawCommissionQuota: systemConfig.Config.WithdrawCommissionQuota,
		constants.RedisKeyWithdrawCommissionRate:  systemConfig.Config.WithdrawCommissionRate,
	}
	info, _ := json.Marshal(systemConfig.Config)
	err = global.GVA_REDIS.Set(context.Background(), constants.RedisKeySystemConfigInfo, info, 0).Err()
	if err != nil {
		return err
	}
	err = global.GVA_REDIS.MSet(context.Background(), mSet).Err()
	if err != nil {
		return err
	}

	return nil
}
