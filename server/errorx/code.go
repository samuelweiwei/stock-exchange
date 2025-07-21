package errorx

// 通用异常码
const (
	InternalServerError = 500
)

// 跟单相关业务校验错误码
const (
	IllegalAdvisorExpr = 100000 + iota
	AdvisorNotFound
	AdvisorProdNotFound
	AdvisorStockOrderNotFound
	AdvisorStockOrderAlreadyFinished
	AdvisorHasOpenedStockOrder
	FollowOrderNotFound
	IllegalFollowAmount
	AppendOrderNotFound
	AdvisorProdDisabled
	FollowAmountShouldGtMin
	FollowAmountShouldLtMax
	CouponNotAllowed
	CouponNotFound
	CouponNotAllowedForProd
	CouponAlreadyUsed
	CouponIsExpired
	FollowAmountShouldGtCouponAmount
	MarketNotOpened
	InsufficientRetrievableAmount
	FollowOrderApprovedFailed
	FollowOrderCancelFailed
	StockNumShouldGtZero
	InsufficientFollowAvailableAmount
	FollowPeriodFinished
	AdvisorHasNotOpenStockOrder
	SymbolNotFount
	FollowOrderRejectFailed
)
