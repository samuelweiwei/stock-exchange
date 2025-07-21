package system

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf config.Server, err error

type SystemConfigService struct{}

var SystemConfigServiceApp = new(SystemConfigService)

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return global.GVA_CONFIG, nil
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

func (systemConfigService *SystemConfigService) SavePlatformSystemConfig(req *request.SystemConfigSaveReq) (err error) {
	var systemConfig system.SystemConfig
	err = global.GVA_DB.Model(&system.SystemConfig{}).Limit(1).First(&systemConfig).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		systemConfig.Config = system.Config{}
		err = nil
	} else if err != nil {
		return err
	}

	if req.PlatformCommissionRate.Valid {
		systemConfig.Config.PlatformCommissionRate = req.PlatformCommissionRate.Float64
	}
	if req.FirstGradeShareRate.Valid {
		systemConfig.Config.FirstGradeShareRate = req.FirstGradeShareRate.Float64
	}
	if req.SecondGradeShareRate.Valid {
		systemConfig.Config.SecondGradeShareRate = req.SecondGradeShareRate.Float64
	}
	if req.ThirdGradeShareRate.Valid {
		systemConfig.Config.ThirdGradeShareRate = req.ThirdGradeShareRate.Float64
	}
	if req.WithdrawCommissionType.Valid {
		systemConfig.Config.WithdrawCommissionType = req.WithdrawCommissionType.Int64
	}
	if req.WithdrawCommissionQuota.Valid {
		systemConfig.Config.WithdrawCommissionQuota = req.WithdrawCommissionQuota.Decimal
	}
	if req.WithdrawCommissionRate.Valid {
		systemConfig.Config.WithdrawCommissionRate = req.WithdrawCommissionRate.Decimal
	}
	if req.IosAppDownloadUrl.Valid {
		systemConfig.Config.IosAppDownloadUrl = req.IosAppDownloadUrl.String
	}
	if req.AndroidAppDownloadUrl.Valid {
		systemConfig.Config.AndroidAppDownloadUrl = req.AndroidAppDownloadUrl.String
	}
	if req.BackWhiteIpStrings.Valid {
		systemConfig.Config.BackWhiteIpStrings = req.BackWhiteIpStrings.String
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) (txErr error) {
		if systemConfig.ID == 0 {
			txErr = global.GVA_DB.Create(&systemConfig).Error
		} else {
			txErr = global.GVA_DB.Save(&systemConfig).Error
		}
		if txErr != nil {
			return txErr
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
		txErr = global.GVA_REDIS.Set(context.Background(), constants.RedisKeySystemConfigInfo, info, 0).Err()
		if txErr != nil {
			return
		}
		txErr = global.GVA_REDIS.MSet(context.Background(), mSet).Err()
		return
	})

	return
}

func (systemConfigService *SystemConfigService) GetPlatformSystemConfig() (resp *response.SystemConfig, err error) {
	var systemConfig system.SystemConfig
	err = global.GVA_DB.Model(&system.SystemConfig{}).Limit(1).First(&systemConfig).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp = &response.SystemConfig{}
		err = nil
	} else if err == nil {
		resp = &response.SystemConfig{
			BackWhiteIpStrings: systemConfig.Config.BackWhiteIpStrings,
			DomainInfo: response.DomainInfo{
				IosAppDownloadUrl:       systemConfig.Config.IosAppDownloadUrl,
				AndroidAppDownloadUrl:   systemConfig.Config.AndroidAppDownloadUrl,
				PlatformCommissionRate:  systemConfig.Config.PlatformCommissionRate,
				FirstGradeShareRate:     systemConfig.Config.FirstGradeShareRate,
				SecondGradeShareRate:    systemConfig.Config.SecondGradeShareRate,
				ThirdGradeShareRate:     systemConfig.Config.ThirdGradeShareRate,
				WithdrawCommissionType:  systemConfig.Config.WithdrawCommissionType,
				WithdrawCommissionQuota: systemConfig.Config.WithdrawCommissionQuota,
				WithdrawCommissionRate:  systemConfig.Config.WithdrawCommissionRate,
			},
		}
	}
	return
}

func (systemConfigService *SystemConfigService) GetDomainInfo() (resp *response.DomainInfo, err error) {
	var (
		systemConfig system.SystemConfig
		sysConfig    *system.Config
	)
	defer func() {
		resp = &response.DomainInfo{}
		if sysConfig != nil {
			resp.PlatformCommissionRate = sysConfig.PlatformCommissionRate
			resp.FirstGradeShareRate = sysConfig.FirstGradeShareRate
			resp.SecondGradeShareRate = sysConfig.SecondGradeShareRate
			resp.ThirdGradeShareRate = sysConfig.ThirdGradeShareRate
			resp.WithdrawCommissionType = sysConfig.WithdrawCommissionType
			resp.WithdrawCommissionQuota = sysConfig.WithdrawCommissionQuota
			resp.WithdrawCommissionRate = sysConfig.WithdrawCommissionRate
			resp.IosAppDownloadUrl = sysConfig.IosAppDownloadUrl
			resp.AndroidAppDownloadUrl = sysConfig.AndroidAppDownloadUrl
		}
	}()

	//先取缓存
	err = global.GVA_REDIS.Get(context.Background(), constants.RedisKeySystemConfigInfo).Scan(&sysConfig)
	if err == nil && sysConfig != nil {
		return
	}

	//如果缓存没值，取数据库
	err = global.GVA_DB.Model(&system.SystemConfig{}).Limit(1).First(&systemConfig).Error
	if err == nil {
		sysConfig = &systemConfig.Config
	}

	return
}
