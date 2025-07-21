package i18n

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/i18n"
	"github.com/flipped-aurora/gin-vue-admin/server/model/i18n/request"
	i18ng "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"sync"
	"time"
)

var (
	defaultLocalization *DefaultLocalization
)

type DefaultLocalization struct {
	ErrorCodeMessages map[string]*i18n.SysI18nLocalizeConfig // the map of error code's messages
	IdMessages        map[string]*i18n.SysI18nLocalizeConfig // the map of message id's messages
	mux               sync.RWMutex
	bundle            *i18ng.Bundle
}

func NewLocalization() *DefaultLocalization {
	dl := &DefaultLocalization{}
	err := dl.Load()
	if err != nil {
		return nil
	}
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		for {
			select {
			case <-ticker.C:
				err = dl.Load()
				fmt.Println(err)
			}
		}
	}()
	return dl
}

func (l *DefaultLocalization) Load() error {
	var (
		list              []*i18n.SysI18nLocalizeConfig
		req               request.SysI18nLocalizeConfigSearch
		err               error
		errorCodeMessages = make(map[string]*i18n.SysI18nLocalizeConfig)
		idMessages        = make(map[string]*i18n.SysI18nLocalizeConfig)
		bundle            = i18ng.NewBundle(language.English)
	)
	list, _, err = i18nService.SysI18nLocalizeConfigService.GetSysI18nLocalizeConfigInfoList(req)
	for _, v := range list {
		errorCodeMessages[l.GetCodeLanguageKey(v.LangTag, *v.ErrorCode)] = v
		idMessages[l.getIdLanguageKey(v.LangTag, v.MessageId)] = v
		_ = bundle.AddMessages(language.Make(v.LangTag), &i18ng.Message{
			ID:    v.MessageId,
			Other: v.TemplateData,
			One:   v.TemplateData,
		})
	}
	l.mux.Lock()
	defer l.mux.Unlock()
	l.IdMessages = idMessages
	l.ErrorCodeMessages = errorCodeMessages
	l.bundle = bundle
	return err
}
func (l *DefaultLocalization) GetCodeMessage(lang, code interface{}) *i18n.SysI18nLocalizeConfig {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return l.ErrorCodeMessages[l.GetCodeLanguageKey(lang, code)]
}

func (l *DefaultLocalization) GetIdMessage(lang, id string) *i18n.SysI18nLocalizeConfig {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return l.IdMessages[l.getIdLanguageKey(lang, id)]
}

func (l *DefaultLocalization) getIdLanguageKey(lang, id string) string {
	return fmt.Sprintf("%v_%v", lang, id)
}

func (l *DefaultLocalization) GetBundle() *i18ng.Bundle {
	return l.bundle
}

func (l *DefaultLocalization) GetCodeLanguageKey(lang, code interface{}) string {
	return fmt.Sprintf("%v_%v", lang, code)
}

func (l *DefaultLocalization) X(langTag string, sc *i18n.SysI18nLocalizeConfig, args ...interface{}) (string, error) {
	var (
		paramError   = errors.New("param error")
		TemplateData = make(map[string]interface{})
	)

	n := len(args)
	if n%2 != 0 {
		return "", paramError
	}
	for i := 0; i < n; {
		TemplateData[fmt.Sprint(args[i])] = args[i+1]
		i = i + 2
	}
	loc := i18ng.NewLocalizer(defaultLocalization.bundle, langTag)
	message := loc.MustLocalize(&i18ng.LocalizeConfig{
		DefaultMessage: &i18ng.Message{
			ID: sc.MessageId,
		},
		TemplateData: TemplateData,
	})
	return message, nil
}

/****
国际化支持场景：
1、业务标题
2、错误码
3、动态参数
****/

func Message(langTag, messageId string, code uint64, args ...interface{}) string {
	var (
		sc *i18n.SysI18nLocalizeConfig
	)
	if messageId != "" {
		sc = defaultLocalization.GetIdMessage(langTag, messageId)
	} else {
		sc = defaultLocalization.GetCodeMessage(langTag, code)
	}
	if sc == nil {
		return ""
	}
	message, _ := defaultLocalization.X(langTag, sc, args...)
	if message == "" {
		message, _ = defaultLocalization.X("en", sc, args...)
	}
	return message
}

func StartUp() {
	defaultLocalization = NewLocalization()
}
