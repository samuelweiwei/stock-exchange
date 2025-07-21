package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type LoginResponse struct {
	User      user.FrontendUsers `json:"user"`
	Token     string             `json:"token"`
	ExpiresAt int64              `json:"expiresAt"`
}

type SubUserResponse struct {
	ID       uint   `json:"ID"`                       // user ID
	NickName string `json:"nickName" form:"nickName"` // 用户名
	// 是否充值
	IsRecharge int `json:"isRecharge"` // 0 未充值 1 已充值
	// 下级总人数
	SubUserCount int64 `json:"subUserCount"`
}
