package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type AppendOrderStatus int

const (
	AppendOrderStatusAuditing = 0
	AppendOrderStatusApproved = 1
	AppendOrderStatusRejected = 2
)

type UserFollowAppendOrder struct {
	global.GVA_MODEL
	AppendOrderNo     string             `gorm:"column:append_order_no"`
	UserId            uint               `gorm:"column:user_id"`
	FollowOrderId     uint               `gorm:"column:follow_order_id"`
	AppendAmount      float64            `gorm:"column:append_amount"`
	AppendOrderStatus AppendOrderStatus  `gorm:"column:append_order_status"`
	UserFollowOrder   UserFollowOrder    `gorm:"foreignKey:FollowOrderId"`
	User              user.FrontendUsers `gorm:"foreignKey:UserId"`
}

func (UserFollowAppendOrder) TableName() string {
	return "user_follow_append_order"
}
