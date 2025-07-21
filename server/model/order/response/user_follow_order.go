package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/guregu/null/v5"
)

// MyFollowOrderPageData 我的跟单分页数据
// @Description 我的跟单分页数据
type MyFollowOrderPageData struct {
	FollowOrderId         string                  `json:"followOrderId"`         //跟单ID
	FollowOrderNo         string                  `json:"followOrderNo"`         //跟单号
	FollowOrderStatus     order.FollowOrderStatus `json:"followOrderStatus"`     //跟单状态： 0-审核中;3-已驳回;5-跟单中;8-已撤回;10-已结束
	PeriodStart           null.Int64              `json:"periodStart"`           //跟单周期开始日期，Unix时间戳
	PeriodEnd             null.Int64              `json:"periodEnd"`             //跟单周期结束日期，Unix时间戳
	AdvisorName           string                  `json:"advisorName"`           //导师名称
	ProductName           string                  `json:"productName"`           //产品名称
	FollowAmount          float64                 `json:"followAmount"`          //跟单金额
	AdvisorCommissionRate float64                 `json:"advisorCommissionRate"` //导师佣金比例
	ApplyTime             null.Int64              `json:"applyTime"`             //申请时间，Unix时间戳
	FinishTime            null.Int64              `json:"finishTime"`            //结束时间，Unix时间戳
	ReturnAmount          float64                 `json:"returnAmount"`          //交易盈亏
	RetrievableAmount     float64                 `json:"retrievableAmount"`     //可提盈金额
	PlatformCommission    float64                 `json:"platformCommission"`    //平台佣金
}

// MyFollowOrderDetail 我的跟单详情
// @Description 我的跟单详情
type MyFollowOrderDetail struct {
	FollowOrderId     string                      `json:"followOrderId"`     //跟单ID
	FollowOrderNo     string                      `json:"followOrderNo"`     //跟单号
	FollowOrderStatus order.FollowOrderStatus     `json:"followOrderStatus"` //跟单状态：0-审核中;3-已驳回;5-跟单中;8-已撤回;10-已结束
	StepAmount        float64                     `json:"stepAmount"`        //金额步长
	RetrievableAmount float64                     `json:"retrievableAmount"` //可提盈金额
	TotalTradeReturn  float64                     `json:"totalTradeReturn"`  //交易总盈亏
	TotalActualReturn float64                     `json:"totalActualReturn"` //实际总盈亏
	FollowAmount      float64                     `json:"followAmount"`      //跟单总金额
	AutoRenew         order.AutoRenewStatus       `json:"autoRenew"`         //是否自动续期：0-否；1-是
	ProductName       string                      `json:"productName"`       //产品名称
	AdvisorName       string                      `json:"advisorName"`       //导师名称
	AdvisorAvatarUrl  string                      `json:"advisorAvatarUrl"`  //导师头像
	CouponName        *string                     `json:"couponName"`        //优惠券名称
	CouponAmount      *float64                    `json:"couponAmount"`      //优惠券金额
	Details           []*MyFollowOrderStockDetail `json:"details"`           //跟单明细
}

// MyFollowOrderStockDetail 我的跟单明细
// @Description 我的跟单明细
type MyFollowOrderStockDetail struct {
	StockName          string     `json:"stockName"`          //股票名称
	BuyPrice           float64    `json:"buyPrice"`           //买入价格
	BuyTime            int64      `json:"buyTime"`            //买入时间，Unix时间戳
	StockNum           float64    `json:"stockNum"`           //股票数量
	SellPrice          null.Float `json:"sellPrice"`          //卖出价格
	SellTime           null.Int64 `json:"sellTime"`           //卖出时间，Unix时间戳
	TradeReturn        float64    `json:"tradeReturn"`        //交易盈亏
	ActualReturn       float64    `json:"actualReturn"`       //实际盈亏
	AdvisorCommission  float64    `json:"advisorCommission"`  //导师佣金
	PlatformCommission float64    `json:"platformCommission"` //平台佣金
}

// UserFollowOrderPageData 用户跟单分页数据
// @Description 用户跟单分页数据
type UserFollowOrderPageData struct {
	FollowOrderId          string                       `json:"followOrderId"`          //用户跟单ID
	AdvisorName            string                       `json:"advisorName"`            //导师名称
	UserName               string                       `json:"userName"`               //客户名称
	UserId                 string                       `json:"userId"`                 //客户ID
	UserType               uint                         `json:"userType"`               //用户类型
	UserTypeText           string                       `json:"userTypeText"`           //用户类型文案
	Phone                  string                       `json:"phone"`                  //客户手机
	ProductName            string                       `json:"productName"`            //产品名称
	AutoRenew              order.AutoRenewStatus        `json:"autoRenew"`              //是否自动续期： 0-否；1-是
	AdvisorCommissionRate  float64                      `json:"advisorCommissionRate"`  //导师佣金比例
	FollowOrderStatus      order.FollowOrderStatus      `json:"followOrderStatus"`      //跟单状态： 0-审核中;3-已驳回;5-跟单中;8-已撤回;10-已结束
	FollowOrderStatusText  string                       `json:"followOrderStatusText"`  //跟单状态文案
	StockStatus            order.FollowOrderStockStatus `json:"stockStatus"`            //跟单持仓状态：0-未开单；1-持仓中；3-已过期
	StockStatusText        string                       `json:"stockStatusText"`        //跟单持仓状态文案
	CouponRecordId         null.Int64                   `json:"couponRecordId"`         //优惠券发放记录ID
	Amount                 float64                      `json:"amount"`                 //跟单总金额
	UserEmail              string                       `json:"userEmail"`              //客户邮箱
	PlatformCommissionRate float64                      `json:"platformCommissionRate"` //平台佣金比例
	ApplyTime              int64                        `json:"applyTime"`              //申请时间
	RootUserId             uint                         `json:"rootUserId"`             //根用户ID
	ParentUserId           uint                         `json:"parentUserId"`           //上级用户ID
}

// UserFollowConfirmDetail 用户跟单确认详情
// @Description 用户跟单确认详情
type UserFollowConfirmDetail struct {
	FollowOrderId   string  `json:"followOrderId"`   //跟单ID
	StockOrderId    string  `json:"stockOrderId"`    //导师开单ID
	UserId          string  `json:"userId"`          //客户ID
	StockName       string  `json:"stockName"`       //股票名称
	BuyPrice        float64 `json:"buyPrice"`        //买入价格
	AvailableAmount float64 `json:"availableAmount"` //可用金额
	MaxStockNum     float64 `json:"maxStockNum"`     //最大股票数量
}

// UserFollowOrderProfitSummary 用户跟单收益报表
// @Description 用户跟单收益报表
type UserFollowOrderProfitSummary struct {
	TotalProfit     float64 `json:"totalProfit"`     //总收益
	TodayProfit     float64 `json:"todayProfit"`     //今日收益
	UserShareProfit float64 `json:"userShareProfit"` //合伙人分成
}

// UserFollowOrderRetrieveRecord 用户跟单提现记录
// @Description 用户跟单提现记录
type UserFollowOrderRetrieveRecord struct {
	Amount       float64 `json:"amount"`       //提现金额
	RetrieveTime int64   `json:"retrieveTime"` //提现时间
}
