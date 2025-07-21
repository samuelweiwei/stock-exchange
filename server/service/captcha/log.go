package captcha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/captcha"
	"github.com/flipped-aurora/gin-vue-admin/server/model/captcha/request"
)

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

// Record 记录验证码
// Author [yourname](https://github.com/yourname)
func (l *LogService) Record(ca *captcha.Captcha) (err error) {
	err = global.GVA_DB.Create(ca).Error
	return err
}

func (l *LogService) AdminList(req request.CaptchaListSearch) (captchaList []*captcha.Captcha, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&captcha.Captcha{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if req.UserId != 0 {
		db = db.Where("user_id = ?", req.UserId)
	}
	if req.AccessIp != "" {
		db = db.Where("access_ip = ?", req.AccessIp)
	}
	if req.CaptchaCode != "" {
		db = db.Where("captcha_code = ?", req.CaptchaCode)
	}
	if req.UserType != 0 {
		db = db.Where("user_type = ?", req.UserType)
	}
	if req.ChannelType != 0 {
		db = db.Where("channel_type = ?", req.ChannelType)
	}
	if req.UniqueId != "" {
		db = db.Where("unique_id = ?", req.UniqueId)
	}
	if req.CountryCode != "" {
		db = db.Where("national_code = ?", req.CountryCode)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&captchaList).Error
	return captchaList, total, err

}
