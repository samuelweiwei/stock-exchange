package captcha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"strings"
)

type Captcha struct {
	NotificationChannel NotificationChannel
	Store               Store
	Driver              Driver
}

// NewCaptcha creates a captcha instance from driver and store
func NewCaptcha(notificationChannel NotificationChannel, driver Driver, store Store) *Captcha {
	return &Captcha{NotificationChannel: notificationChannel, Driver: driver, Store: store}
}

// Generate generates a random id, base64 image string or an error if any
func (c *Captcha) Generate(target string) (id, captcha string, err error) {
	id, captcha = c.Driver.GenerateIdCaptcha(target)
	err = c.NotificationChannel.SetCaptcha(target, captcha)
	if err != nil {
		global.GVA_LOG.Info("Captcha SetCaptcha Error", zap.Error(err))
		return "", "", err
	}
	err = c.Store.Set(id, captcha)
	if err != nil {
		global.GVA_LOG.Info("Captcha Store Error", zap.Error(err))
		return "", "", err
	}
	return
}

// Verify by a given id key and remove the captcha value in store,
// return boolean value.
// if you has multiple captcha instances which share a same store.
// You may want to call `store.Verify` method instead.
func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
	vv := c.Store.Get(id, clear)
	//fix issue for some redis key-value string value
	vv = strings.TrimSpace(vv)
	return vv == strings.TrimSpace(answer)
}
