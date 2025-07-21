// 自动生成模板Advisor
package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ActiveStatus int

const (
	Active   ActiveStatus = 1
	Inactive ActiveStatus = 0
)

// advisor表 结构体  Advisor
type Advisor struct {
	global.GVA_MODEL
	NickName           string       `gorm:"column:nick_name"`
	Duty               string       `gorm:"column:duty"`
	Intro              string       `gorm:"column:intro"`
	Exp                *int         `gorm:"column:exp"`
	AvatarUrl          string       `gorm:"column:avatar_url"`
	SevenDayReturn     *float64     `gorm:"column:seven_day_return"`
	SevenDayReturnRate *float64     `gorm:"column:seven_day_return_rate"`
	ActiveStatus       ActiveStatus `gorm:"column:active_status"`
}

// TableName advisor表 Advisor自定义表名 advisor
func (Advisor) TableName() string {
	return "advisor"
}
