package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"gorm.io/gorm"
	"time"
)

type EarnSubscribeLogService struct{}

// CreateEarnSubscribeLog 创建earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) CreateEarnSubscribeLog(tx *gorm.DB, earnSubscribeLog *earn.EarnSubscribeLog) (err error) {
	if tx == nil {
		tx = global.GVA_DB
	}
	err = tx.Create(earnSubscribeLog).Error
	return err
}

// DeleteEarnSubscribeLog 删除earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) DeleteEarnSubscribeLog(id string) (err error) {
	err = global.GVA_DB.Delete(&earn.EarnSubscribeLog{}, "id = ?", id).Error
	return err
}

// DeleteEarnSubscribeLogByIds 批量删除earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) DeleteEarnSubscribeLogByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]earn.EarnSubscribeLog{}, "id in ?", ids).Error
	return err
}

// UpdateEarnSubscribeLog 更新earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) UpdateEarnSubscribeLog(earnSubscribeLog earn.EarnSubscribeLog) (err error) {
	earnSubscribeLog.UpdatedAtX = time.Now().UnixMilli()
	err = global.GVA_DB.Model(&earn.EarnSubscribeLog{}).Where("id = ?", earnSubscribeLog.Id).Updates(&earnSubscribeLog).Error
	return err
}

// UpdateEarnSubscribeLog 更新earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) Fine(tx *gorm.DB, earnSubscribeLog earn.EarnSubscribeLog) (err error) {
	if tx == nil {
		tx = global.GVA_DB
	}
	earnSubscribeLog.UpdatedAtX = time.Now().UnixMilli()
	err = tx.Model(&earn.EarnSubscribeLog{}).Select("status", "fine", "updatedAt").Where("id = ?", earnSubscribeLog.Id).Save(&earnSubscribeLog).Error
	return err
}

// GetUserEarnSubscribeLog 根据id获取earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) GetUserEarnSubscribeLog(id string, uid uint) (earnSubscribeLog earn.EarnSubscribeLog, err error) {
	err = global.GVA_DB.Where("id = ? and uid = ?", id, uid).First(&earnSubscribeLog).Error
	return
}

// GetUserEarnSubscribeLog 根据id获取earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) GetEarnSubscribeLog(id string) (earnSubscribeLog earn.EarnSubscribeLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&earnSubscribeLog).Error
	return
}

// GetEarnSubscribeLog 根据id获取earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) GetUserEarnSubscribeLogList(id uint) (earnSubscribeLog []*earn.EarnSubscribeLog, err error) {
	err = global.GVA_DB.Where("uid = ?", id).Order("id desc").Find(&earnSubscribeLog).Error
	return
}

// GetEarnSubscribeLogInfoList 分页获取earnSubscribeLog表记录
// Author [yourname](https://github.com/yourname)
func (earnSubscribeLogService *EarnSubscribeLogService) GetEarnSubscribeLogInfoList(info earnReq.EarnSubscribeLogSearch) (list []earn.EarnSubscribeLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnSubscribeLog{})
	var earnSubscribeLogs []earn.EarnSubscribeLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if info.Status != earnReq.UnKnown {
		db = db.Where("status = ?", info.Status)
	}
	if info.Uid != 0 {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.UserType != 0 {
		db = db.Where("user_type = ?", info.UserType)
	}
	if info.RedeemInAdvance != earnReq.RedeemUnKnown {
		if info.RedeemInAdvance == earnReq.RedeemNormal {
			db = db.Where("redeem_in_advance in (?)", []earnReq.RedeemStatus{earnReq.RedeemUnKnown, earnReq.RedeemNormal})
		} else {
			db = db.Where("redeem_in_advance = ?", info.RedeemInAdvance)
		}
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc ").Find(&earnSubscribeLogs).Error
	return earnSubscribeLogs, total, err
}

func (earnSubscribeLogService *EarnSubscribeLogService) FindExpiredSubscribeLog(tx *gorm.DB, endAt int64) (res []*earn.EarnSubscribeLog, err error) {
	if tx == nil {
		tx = global.GVA_DB
	}
	err = tx.Model(&earn.EarnSubscribeLog{}).Where("end_at < ? and status = ?", endAt, earn.Staking).Find(&res).Error
	return res, err
}

func (earnSubscribeLogService *EarnSubscribeLogService) FindStakingSubscribeLog(tx *gorm.DB, endAt int64) (res []*earn.EarnSubscribeLog, err error) {
	if tx == nil {
		tx = global.GVA_DB
	}
	err = tx.Model(&earn.EarnSubscribeLog{}).Where("end_at > ? and status = ?", endAt, earn.Staking).Find(&res).Error
	return res, err
}

func (earnSubscribeLogService *EarnSubscribeLogService) GetEarnSubscribeLogPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
