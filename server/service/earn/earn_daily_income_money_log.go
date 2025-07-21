package earn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/earn"
	earnReq "github.com/flipped-aurora/gin-vue-admin/server/model/earn/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var (
	dummy earn.EarnDailyIncomeMoneyLog
)

type TotalEarnings struct {
	SubscribeId uint    `json:"subscribe_id,omitempty" gorm:"subscribe_id"`
	N           float64 `json:"totalEarnings" `
}

type EarnDailyIncomeMoneyLogService struct{}

// CreateEarnDailyIncomeMoneyLog 创建earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) CreateEarnDailyIncomeMoneyLog(earnDailyIncomeMoneyLog *earn.EarnDailyIncomeMoneyLog) (err error) {
	err = global.GVA_DB.Create(earnDailyIncomeMoneyLog).Error
	return err
}

// DeleteEarnDailyIncomeMoneyLog 删除earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) DeleteEarnDailyIncomeMoneyLog(id string) (err error) {
	err = global.GVA_DB.Delete(&earn.EarnDailyIncomeMoneyLog{}, "id = ?", id).Error
	return err
}

// DeleteEarnDailyIncomeMoneyLogByIds 批量删除earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) DeleteEarnDailyIncomeMoneyLogByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]earn.EarnDailyIncomeMoneyLog{}, "id in ?", ids).Error
	return err
}

// UpdateEarnDailyIncomeMoneyLog 更新earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) UpdateEarnDailyIncomeMoneyLog(earnDailyIncomeMoneyLog earn.EarnDailyIncomeMoneyLog) (err error) {
	err = global.GVA_DB.Model(&earn.EarnDailyIncomeMoneyLog{}).Where("id = ?", earnDailyIncomeMoneyLog.Id).Updates(&earnDailyIncomeMoneyLog).Error
	return err
}

// GetEarnDailyIncomeMoneyLog 根据id获取earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetEarnDailyIncomeMoneyLog(id string) (earnDailyIncomeMoneyLog earn.EarnDailyIncomeMoneyLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&earnDailyIncomeMoneyLog).Error
	return
}

// GetEarnDailyIncomeMoneyLogList 根据id获取earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetEarnDailyIncomeSummaryListGroupBySubscribe(idList []uint) (res []*TotalEarnings, err error) {
	res = make([]*TotalEarnings, 0)
	err = global.GVA_DB.Debug().Table(dummy.TableName()).Where("subscribe_id in(?)", idList).Group("subscribe_id").
		Select("sum(earnings) as n, subscribe_id ").Scan(&res).Error
	return
}

func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetFrontEarnProductDailyIncomeMoneyLogDetail(id uint, subscriptionId uint, page, pageSize int) (res []*earn.EarnDailyIncomeMoneyLog, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	db := global.GVA_DB.Where("uid = ? and id = ?", id, subscriptionId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&res).Error
	return
}
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) BatchSave(tx *gorm.DB, res []*earn.EarnDailyIncomeMoneyLog) (err error) {
	if tx == nil {
		tx = global.GVA_DB.Begin()
	}
	return tx.Create(&res).Error
}

// GetEarnDailyIncomeMoneyLogInfoList 分页获取earnDailyIncomeMoneyLog表记录
// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetEarnDailyIncomeMoneyLogInfoSummary(info earnReq.EarnDailyIncomeMoneyLogSearch) (list []earn.EarnDailyIncomeMoneyLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnDailyIncomeMoneyLog{})
	var earnDailyIncomeMoneyLogs []earn.EarnDailyIncomeMoneyLog
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&earnDailyIncomeMoneyLogs).Error
	return earnDailyIncomeMoneyLogs, total, err
}

// Author [yourname](https://github.com/yourname)
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetEarnDailyIncomeMoneyLogInfoList(users []uint, req earnReq.EarnDailyIncomeMoneyLogSearch) (list []*earn.EarnDailyIncomeMoneyLog, total int64, err error) {
	//limit := req.PageSize
	//offset := req.PageSize * (req.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnDailyIncomeMoneyLog{})
	var earnDailyIncomeMoneyLogs []*earn.EarnDailyIncomeMoneyLog

	/***
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}***/
	if req.BeginTime != 0 {
		db = db.Where("created_at >= ?", req.BeginTime)
		db = db.Where("created_at < ?", req.EndTime)
	}
	if len(users) > 0 {
		db = db.Where("uid in(?)", users)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Find(&earnDailyIncomeMoneyLogs).Error
	return earnDailyIncomeMoneyLogs, total, err
}

func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetFrontEarnDailyIncomeMoneyLogInfoList(users []uint, req earnReq.EarnDailyIncomeMoneyLogSearch) (list []*earn.EarnDailyIncomeMoneyLog, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&earn.EarnDailyIncomeMoneyLog{})
	var earnDailyIncomeMoneyLogs []*earn.EarnDailyIncomeMoneyLog

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if req.BeginTime != 0 {
		db = db.Where("created_at >= ?", req.BeginTime)
		db = db.Where("created_at < ?", req.EndTime)
	}
	if len(users) >= 0 {
		db = db.Where("uid in(?)", users)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Find(&earnDailyIncomeMoneyLogs).Error
	return earnDailyIncomeMoneyLogs, total, err
}

func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetUserProductionEarnings(tx *gorm.DB, subscribeId uint, uid uint) (earned decimal.Decimal, err error) {
	if tx == nil {
		tx = global.GVA_DB.Begin()
	}
	err = tx.Model(&earn.EarnDailyIncomeMoneyLog{}).Select("sum(earnings) as earnings ").Group("subscribe_id, uid").
		Where("subscribe_id = ? and uid = ?", subscribeId, uid).First(&earned).Error
	return earned, err
}
func (earnDailyIncomeMoneyLogService *EarnDailyIncomeMoneyLogService) GetEarnDailyIncomeMoneyLogPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
