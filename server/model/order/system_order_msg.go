package order

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type MsgType uint
type MsgReadStatus uint

const (
	MsgTypeNewFollowOrder MsgType = 0
	MsgTypeNewAppendOrder MsgType = 1
)
const (
	MsgReadStatusUnRead MsgReadStatus = 0
	MsgReadStatusRead   MsgReadStatus = 1
)

type SystemOrderMsg struct {
	global.GVA_MODEL
	Type       MsgType       `gorm:"column:type"`
	OrderId    uint          `gorm:"column:order_id"`
	UserId     uint          `gorm:"column:user_id"`
	ReadStatus MsgReadStatus `gorm:"column:read_status"`
}

func (SystemOrderMsg) TableName() string {
	return "system_order_msg"
}
