package captcha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type NotificationChannel interface {
	SetCaptcha(target, captcha string) error
}

type PhoneMessageChannel struct {
	sms SMS
}

func NewPhoneMessageChannel() NotificationChannel {
	return &PhoneMessageChannel{
		sms: NewSMS(constants.SMSAppKey, constants.SMSAppSecret, constants.SMSAppCode),
	}
}

func (c *PhoneMessageChannel) SetCaptcha(target, captcha string) error {
	global.GVA_LOG.Info("PhoneMessageChannel SetCaptcha ", zap.String("target", target), zap.String("captcha", captcha))
	err := c.sms.Send(target, captcha)
	global.GVA_LOG.Info("PhoneMessageChannel SetCaptcha ", zap.Error(err), zap.String("target", target), zap.String("captcha", captcha))
	return nil
}

type EmailMessageChannel struct {
	email Email
}

func NewEmailMessageChannel() NotificationChannel {
	return &EmailMessageChannel{
		email: NewAokSend(),
	}
}

func (c *EmailMessageChannel) SetCaptcha(target, captcha string) (err error) {
	global.GVA_LOG.Info("EmailMessageChannel SetCaptcha ", zap.String("target", target), zap.String("captcha", captcha))
	err = c.email.Send(target, captcha)
	return nil
}
