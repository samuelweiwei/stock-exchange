package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
)

// AdvisorProdCreateReq 导师产品创建请求
// @Description 导师产品创建请求
type AdvisorProdCreateReq struct {
	AdvisorId      uint                  `json:"advisorId"`      //导师ID
	ProductName    string                `json:"productName"`    //产品名称
	AmountStep     float64               `json:"amountStep"`     //金额步长
	MaxAmount      float64               `json:"maxAmount"`      //最大金额
	MinAmount      float64               `json:"minAmount"`      //最小金额
	FollowPeriod   int                   `json:"followPeriod"`   //跟单周期，单位天
	CommissionRate float64               `json:"commissionRate"` //导师佣金比例，单位%
	AutoRenew      order.AutoRenewStatus `json:"autoRenew"`      //是否开启自动续期：0-否；1-是
	ActiveStatus   order.ActiveStatus    `json:"activeStatus"`   //导师启用状态：0-否；1-是
	CouponIdList   []uint                `json:"couponIdList"`   //允许使用优惠券ID
}

// AdvisorProdUpdateReq 导师产品更新请求
// @Description 导师产品更新请求
type AdvisorProdUpdateReq struct {
	AdvisorProdCreateReq
	Id uint `json:"-"`
}

// AdvisorProdPageQueryReq 导师产品分页查询请求
type AdvisorProdPageQueryReq struct {
	request.PageInfo
}
