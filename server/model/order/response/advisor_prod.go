package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/order"

// AdvisorProdPageData 导师产品分页数据
// @Description 导师产品分页数据
type AdvisorProdPageData struct {
	AdvisorProdId  uint                  `json:"advisorProdId"`  //导师产品ID
	AdvisorName    string                `json:"advisorName"`    //导师产品名称
	ProductName    string                `json:"productName"`    //产品名称
	FollowPeriod   int                   `json:"followPeriod"`   //跟单周期，单位天
	CommissionRate float64               `json:"commissionRate"` //导师佣金比例，单位%
	MaxAmount      float64               `json:"maxAmount"`      //最大金额
	MinAmount      float64               `json:"minAmount"`      //最小金额
	AmountStep     float64               `json:"amountStep"`     //金额步长
	AutoRenew      order.AutoRenewStatus `json:"autoRenew"`      //是否允许自动续期：0-否；1-是
	ActiveStatus   order.ActiveStatus    `json:"activeStatus"`   //启用状态：0-否；1-是
}

// AdvisorProdListData 导师产品列表数据
// @Description 导师产品列表数据
type AdvisorProdListData struct {
	AdvisorProdId  uint                       `json:"advisorProdId"`  //导师产品ID
	ProductName    string                     `json:"productName"`    //导师产品名称
	FollowPeriod   int                        `json:"followPeriod"`   //跟单周期，单位天
	CommissionRate float64                    `json:"commissionRate"` //导师佣金比例，单位%
	MaxAmount      float64                    `json:"maxAmount"`      //最大金额
	MinAmount      float64                    `json:"minAmount"`      //最小金额
	AutoRenew      order.AutoRenewStatus      `json:"autoRenew"`      //是否允许自动续期：0-否；1-是
	AutoRenewText  string                     `json:"autoRenewText"`  //是否允许自动续期文案
	UsableCoupons  []*AdvisorProdUsableCoupon `json:"usableCoupons"`  //可用优惠券列表
}

// AdvisorProdUsableCoupon 导师产品可用优惠券
// @Description 导师产品可用优惠券
type AdvisorProdUsableCoupon struct {
	CouponRecordId uint   `json:"couponRecordId"` //优惠券发放记录
	CouponName     string `json:"couponName"`     //优惠券名称
}

// AdvisorProdDetail 导师产品详情
// @Description 导师产品详情
type AdvisorProdDetail struct {
	AdvisorProdId  uint                  `json:"advisorProdId"`  //导师产品ID
	AdvisorId      uint                  `json:"advisorId"`      //导师ID
	ProductName    string                `json:"productName"`    //产品名称
	FollowPeriod   int                   `json:"followPeriod"`   //跟单周期，单位天
	CommissionRate float64               `json:"commissionRate"` //导师佣金比例，单位%
	MaxAmount      float64               `json:"maxAmount"`      //最大金额
	MinAmount      float64               `json:"minAmount"`      //最小金额
	AmountStep     float64               `json:"amountStep"`     //金额步长
	AutoRenew      order.AutoRenewStatus `json:"autoRenew"`      //是否允许自动续期：0-否；1-是
	ActiveStatus   order.ActiveStatus    `json:"activeStatus"`   //启用状态： 0-否；1-是
	CouponIdList   []uint                `json:"couponIdList"`   //允许客户使用优惠券ID集合
}
