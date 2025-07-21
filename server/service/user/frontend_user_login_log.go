package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type FrontendUserLoginLogService struct{}

// CreateFrontendUserLoginLog 创建frontendUserLoginLog表记录
// Author [yourname](https://github.com/yourname)
func (frontendUserLoginLogService *FrontendUserLoginLogService) CreateFrontendUserLoginLog(frontendUserLoginLog *user.FrontendUserLoginLog) (err error) {
	rowsAffected := global.GVA_DB.Where("login_ip = ?", frontendUserLoginLog.LoginIp).First(&user.FrontendUserLoginLog{}).RowsAffected
	if err != nil {
		return
	}
	if rowsAffected > 0 {
		frontendUserLoginLog.IsSameIp = 1
	} else {
		frontendUserLoginLog.IsSameIp = 0
	}
	err = global.GVA_DB.Create(frontendUserLoginLog).Error
	return err
}

// DeleteFrontendUserLoginLog 删除frontendUserLoginLog表记录
// Author [yourname](https://github.com/yourname)
func (frontendUserLoginLogService *FrontendUserLoginLogService) DeleteFrontendUserLoginLog(id string) (err error) {
	err = global.GVA_DB.Delete(&user.FrontendUserLoginLog{}, "id = ?", id).Error
	return err
}

// DeleteFrontendUserLoginLogByIds 批量删除frontendUserLoginLog表记录
// Author [yourname](https://github.com/yourname)
func (frontendUserLoginLogService *FrontendUserLoginLogService) DeleteFrontendUserLoginLogByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]user.FrontendUserLoginLog{}, "id in ?", ids).Error
	return err
}

// UpdateFrontendUserLoginLog 更新frontendUserLoginLog表记录
// Author [yourname](https://github.com/yourname)
func (frontendUserLoginLogService *FrontendUserLoginLogService) UpdateFrontendUserLoginLog(frontendUserLoginLog user.FrontendUserLoginLog) (err error) {
	err = global.GVA_DB.Model(&user.FrontendUserLoginLog{}).Where("id = ?", frontendUserLoginLog.Id).Updates(&frontendUserLoginLog).Error
	return err
}

func (frontendUserLoginLogService *FrontendUserLoginLogService) UpdateFrontendUserLoginLogRegionByIp(loginIp, loginRegion string) (err error) {
	err = global.GVA_DB.Model(&user.FrontendUserLoginLog{}).Where("login_ip = ?", loginIp).Update("login_region", loginRegion).Error
	return err
}

// GetFrontendUserLoginLog 根据id获取frontendUserLoginLog表记录
// Author [yourname](https://github.com/yourname)
func (frontendUserLoginLogService *FrontendUserLoginLogService) GetFrontendUserLoginLog(id string) (frontendUserLoginLog user.FrontendUserLoginLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&frontendUserLoginLog).Error
	return
}

// GetFrontendUserLoginLogInfoList 分页获取frontendUserLoginLog表记录
// Author [yourname](https://github.com/yourname)
func (frontendUserLoginLogService *FrontendUserLoginLogService) GetFrontendUserLoginLogInfoList(info userReq.FrontendUserLoginLogSearch) (list []user.FrontendUserLoginLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.FrontendUserLoginLog{})
	var frontendUserLoginLogs []user.FrontendUserLoginLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartLoginTime != nil && info.EndLoginTime != nil {
		db = db.Where("login_time BETWEEN ? AND ?", info.StartLoginTime, info.EndLoginTime)
	}
	if info.Uid != nil {
		db = db.Where("uid = ?", *info.Uid)
	}
	if info.IsSameIp != nil {
		db = db.Where("is_same_ip = ?", *info.IsSameIp)
	}
	if info.LoginIp != nil {
		db = db.Where("login_ip = ?", *info.LoginIp)
	}
	if info.LoginRegion != nil {
		db = db.Where("login_region like ?", "%"+*info.LoginRegion+"%")
	}
	if info.UserAgent != nil {
		db = db.Where("user_agent like ?", "%"+*info.UserAgent+"%")
	}

	// 关联用户表查询
	db = db.Joins("JOIN frontend_users ON frontend_users.id = frontend_user_login_log.uid").
		Select("frontend_user_login_log.*, frontend_users.phone, frontend_users.email, frontend_users.username")

	// 用户信息查询
	if info.Phone != nil {
		db = db.Where("frontend_users.phone = ?", *info.Phone)
	}
	if info.Email != nil {
		db = db.Where("frontend_users.email = ?", *info.Email)
	}
	if info.UserName != nil {
		db = db.Where("frontend_users.username like ?", "%"+*info.UserName+"%")
	}

	err = db.Order("id desc").Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&frontendUserLoginLogs).Error
	return frontendUserLoginLogs, total, err
}
