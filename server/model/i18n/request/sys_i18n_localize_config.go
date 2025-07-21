package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysI18nLocalizeConfigSearch struct {
	TagLang          string `json:"langTag,omitempty" form:"langTag"`
	MessageId        string `json:"messageId,omitempty" form:"messageId"`
	request.PageInfo `json:"request_._page_info"`
}
