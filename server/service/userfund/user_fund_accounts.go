package userfund

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type UserFundAccountsService struct {
	UserAccountFlowService *UserAccountFlowService
	tx                     *gorm.DB
	notAutoCommit          bool
	TotalChange            Decimal
	FrozenChange           Decimal
	AvailableChange        Decimal
}

func NewUserFundAccountService(tx *gorm.DB, autoCommit bool, options ...Option) *UserFundAccountsService {
	// 在构造函数中初始化 UserFundService
	r := &UserFundAccountsService{
		UserAccountFlowService: &UserAccountFlowService{}, // 自动初始化 UserFundService
		tx:                     tx,
		notAutoCommit:          false,
	}
	if !autoCommit {
		r.notAutoCommit = true
	}
	for _, option := range options {
		option(r)
	}
	return r
}

// CreateUserFundAccounts 创建userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) CreateUserFundAccounts(userFundAccounts *userfund.UserFundAccounts) (err error) {
	err = global.GVA_DB.Create(userFundAccounts).Error
	return err
}

// DeleteUserFundAccounts 删除userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) DeleteUserFundAccounts(id string) (err error) {
	err = global.GVA_DB.Delete(&userfund.UserFundAccounts{}, "id = ?", id).Error
	return err
}

// DeleteUserFundAccountsByIds 批量删除userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) DeleteUserFundAccountsByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.UserFundAccounts{}, "id in ?", ids).Error
	return err
}

// UpdateUserFundAccounts 更新userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) UpdateUserFundAccounts(userFundAccounts userfund.UserFundAccounts) (err error) {
	err = global.GVA_DB.Model(&userfund.UserFundAccounts{}).Where("id = ?", userFundAccounts.Id).Updates(&userFundAccounts).Error
	return err
}

// GetUserFundAccounts 根据id获取userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) GetUserFundAccounts(id string) (userFundAccounts userfund.UserFundAccounts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userFundAccounts).Error
	return
}

// GetUserFundAccounts 根据id获取userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) GetUserFundAccountsByUserId(userId int) (userFundAccounts userfund.UserFundAccounts, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).First(&userFundAccounts).Error
	return
}

// GetUserFundAccounts 根据id获取userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) GetUserFundAccountsByUserIdWithTx(userId int, tx *gorm.DB) (userFundAccounts userfund.UserFundAccounts, err error) {
	err = tx.Model(&userfund.UserFundAccounts{}).Where("user_id = ?", userId).First(&userFundAccounts).Error
	return
}

// GetUserFundAccountsInfoList 分页获取userFundAccounts表记录
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) GetUserFundAccountsInfoList(info userfundReq.UserFundAccountsSearch) (list []userfund.UserFundAccounts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userfund.UserFundAccounts{})
	var userFundAccountss []userfund.UserFundAccounts
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userFundAccountss).Error
	return userFundAccountss, total, err
}
func (userFundAccountsService *UserFundAccountsService) GetUserFundAccountsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Recharge 用户充值接口
// Author [yourname](https://github.com/yourname)
func (userFundAccountsService *UserFundAccountsService) Recharge() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&userfund.UserFundAccounts{})
	return db.Error
}

func (userFundAccountsService *UserFundAccountsService) UpdateUserFundAccountsAndNewFlow(userId int, actionType fund.ActionType, amount float64, orderId string) error {
	fmt.Printf("调用新增订单流水方法：")
	fmt.Printf("userId: %d\n", userId)
	fmt.Printf("actionType: %d\n", actionType)
	fmt.Printf("amount: %f\n", amount)
	fmt.Printf("orderId: %s\n", orderId)
	var rowsAffected int64
	tx := global.GVA_DB.Begin()
	if userFundAccountsService.tx != nil {
		tx = userFundAccountsService.tx
	}
	if tx.Error != nil {
		// 如果开启事务失败，则返回错误
		return tx.Error
	}
	//计算amount
	amountVal := NewFromFloat(amount)
	//用户资金表-余额更新

	// 使用行级锁获取 userFundAccount
	var userfundAccount userfund.UserFundAccounts
	//userfundAccount, err := userFundAccountsService.GetUserFundAccountsByUserIdWithTx(userId, tx)
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ?", userId).
		First(&userfundAccount)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应用户的资金账户记录")
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(err))
		tx.Rollback()
		return err
	}

	var frontendUsers user.FrontendUsers
	err := tx.Model(user.FrontendUsers{}).Where("id = ?", cast.ToString(userId)).First(&frontendUsers).Error
	if err != nil {
		global.GVA_LOG.Error("查询用户失败!", zap.Error(err))
		return err
	}

	changedAmount := getChangeAmount(actionType, amountVal)
	if changedAmount.LessThan(Zero) && actionType != fund.Withdraw {
		if userfundAccount.Balance.LessThan(amountVal) || userfundAccount.AvailableBalance.LessThan(amountVal) {
			global.GVA_LOG.Error("资金不足!", zap.Error(err))
			tx.Rollback()
			return fund.BalanceDoesNotEnoughError
		}
	}
	//新增流水
	// 根据 actionType 处理资金变动
	var beforeBalance Decimal
	var afterBalance Decimal

	beforeBalance = userfundAccount.Balance // 保存变动前的余额
	frozenBalanceBefore := userfundAccount.FrozenBalance
	frozenBalanceAfter := userfundAccount.FrozenBalance //不涉及冻结金额变化的情况下二者相同
	avaBalanceBefore := userfundAccount.AvailableBalance
	// 根据 actionType 判断变动金额
	switch actionType {
	case fund.Recharge:
		afterBalance = beforeBalance.Add(amountVal)
	case fund.Withdraw:
		afterBalance = beforeBalance.Sub(amountVal)
	case fund.TransferToContract:
		afterBalance = beforeBalance.Sub(amountVal)
	case fund.TransferFromContract:
		afterBalance = beforeBalance.Add(amountVal)
	case fund.TradeFollow:
		afterBalance = beforeBalance.Sub(amountVal)
	case fund.CancelTradeFollow:
		afterBalance = beforeBalance.Add(amountVal)
	case fund.OperationRefused:
		afterBalance = beforeBalance.Add(amountVal)
	case fund.SettleProfit:
		afterBalance = beforeBalance.Add(amountVal)
	case fund.AutoSettle:
		//分润
		afterBalance = beforeBalance.Add(amountVal)
	case fund.ProfitSharing:
		//分润
		afterBalance = beforeBalance.Add(amountVal)
	case fund.SysSend:
		//分润
		afterBalance = beforeBalance.Add(amountVal)
	case fund.ApplyOrderFollow:
		//分润
		afterBalance = beforeBalance.Sub(amountVal)
	case fund.RefusedApplyOrderFollow:
		//分润
		afterBalance = beforeBalance.Add(amountVal)
	case fund.StakeEarnProduct:
		afterBalance = beforeBalance //总额不变
	case fund.RedeemEarnProduct:
		afterBalance = beforeBalance.Add(userFundAccountsService.TotalChange)
	default:
		return fmt.Errorf("unknown action type: %s", actionType)
	}
	if actionType == fund.Withdraw {
		frozenBalanceAfter = userfundAccount.FrozenBalance.Add(changedAmount)
		userfundAccount.FrozenBalance = frozenBalanceAfter
		userfundAccount.Balance = userfundAccount.Balance.Add(changedAmount)
	} else if actionType == fund.StakeEarnProduct {
		frozenBalanceAfter = userfundAccount.FrozenBalance.Sub(changedAmount) //计算冻结后金额
		userfundAccount.FrozenBalance = frozenBalanceAfter
		userfundAccount.AvailableBalance = userfundAccount.AvailableBalance.Add(changedAmount)
	} else if actionType == fund.RedeemEarnProduct {
		frozenBalanceAfter = userfundAccount.FrozenBalance.Add(userFundAccountsService.FrozenChange)
		userfundAccount.FrozenBalance = frozenBalanceAfter
		userfundAccount.AvailableBalance = userfundAccount.AvailableBalance.Add(userFundAccountsService.AvailableChange)
		userfundAccount.Balance = afterBalance
	} else {
		avaBalance := userfundAccount.AvailableBalance.Add(changedAmount)
		userfundAccount.Balance = afterBalance
		userfundAccount.AvailableBalance = avaBalance
	}
	avaBalanceAfter := userfundAccount.AvailableBalance
	// 创建新的流水记录
	userAccountFlow := userfund.UserAccountFlow{
		GVA_MODEL:           global.GVA_MODEL{},
		UserId:              userfundAccount.UserId,
		ParentId:            int(frontendUsers.ParentId),
		RootId:              int(frontendUsers.RootUserid),
		TransactionType:     string(actionType),
		Amount:              amountVal,
		BalanceBefore:       beforeBalance,
		BalanceAfter:        afterBalance,
		FrozenBalanceBefore: frozenBalanceBefore,
		FrozenBalanceAfter:  frozenBalanceAfter,
		AvaBalanceBefore:    avaBalanceBefore,
		AvaBalanceAfter:     avaBalanceAfter,
		TransactionDate:     time.Now(),
		Description:         "",
		OrderId:             orderId,
		UserType:            userfundAccount.UserType,
	}

	defer func() {
		if !userFundAccountsService.notAutoCommit {
			tx.Rollback()
		}
	}()
	err = tx.Model(&userfund.UserAccountFlow{}).Create(&userAccountFlow).Error
	if err != nil {
		global.GVA_LOG.Error("新增用户流水失败!", zap.Error(err))
		tx.Rollback()
		return err
	}

	userfundAccount.UpdatedAt = func() *time.Time {
		now := time.Now()
		return &now
	}()

	//如果是充值
	if (userfundAccount.FirstChargeAmount.Decimal.Equal(Zero) || userfundAccount.FirstChargeTime.IsZero()) && actionType == fund.Recharge {
		userfundAccount.FirstChargeAmount.Decimal = changedAmount
		now := time.Now()
		userfundAccount.FirstChargeTime = &now
	}
	if userfundAccount.AvailableBalance.LessThan(NewFromFloat(0.0)) {
		global.GVA_LOG.Error("用户可用金额计算出错!")
		return err
	}
	result = tx.Model(&userfund.UserFundAccounts{}).Where("id = ?", userfundAccount.Id).Updates(&userfundAccount)

	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("查询用户资金账户失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}
	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应用户的资金账户记录")
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(err))
		tx.Rollback()
		return err
	}
	// 提交事务
	if !userFundAccountsService.notAutoCommit {
		err = tx.Commit().Error
	}
	if err != nil {
		// 如果提交事务失败，回滚事务并返回错误
		return err
	}
	return nil
}

func (userFundAccountsService *UserFundAccountsService) AmountSend(userIdStr string, amount float64) {
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}
	userFundAccountsService.UpdateUserFundAccountsAndNewFlow(userId, fund.SysSend, amount, "")
}

func getChangeAmount(actionType fund.ActionType, amount Decimal) Decimal {
	// 根据 actionType 判断变动金额
	switch actionType {
	case fund.Recharge:
		return amount
	case fund.Withdraw:
		return amount.Neg()
	case fund.TransferToContract:
		return amount.Neg()
	case fund.TransferFromContract:
		return amount
	case fund.TradeFollow:
		return amount.Neg()
	case fund.CancelTradeFollow:
		return amount
	case fund.OperationRefused:
		return amount
	case fund.SettleProfit:
		return amount
	case fund.AutoSettle:
		return amount
	case fund.ProfitSharing:
		return amount
	case fund.SysSend:
		return amount
	case fund.ApplyOrderFollow:
		return amount.Neg()
	case fund.RefusedApplyOrderFollow:
		return amount
	case fund.StakeEarnProduct:
		return amount.Neg()
	case fund.RedeemEarnProduct:
		return amount
	default:
		return Zero // 未知操作返回 0
	}
}

func (userFundAccountsService *UserFundAccountsService) UpdateUserFundAccountsWithFlowAndRechargeRecords(rechargeRecords userfund.RechargeRecords) error {
	var rowsAffected int64
	// 开启事务
	tx := global.GVA_DB.Begin()
	if tx.Error != nil {
		// 如果开启事务失败，则返回错误
		return tx.Error
	}
	//计算amount
	//用户资金表-余额更新
	userId := rechargeRecords.MemberId
	var userfundAccount userfund.UserFundAccounts
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ?", userId).
		First(&userfundAccount)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应用户的资金账户记录")
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(err))
		tx.Rollback()
		return err
	}
	//更新订单表
	//err := global.GVA_DB.Model(&userfund.RechargeRecords{}).Where("id = ?", rechargeRecords.ID).Updates(&rechargeRecords).Error
	rechargeRecords.UserType = userfundAccount.UserType
	result = tx.Model(&userfund.RechargeRecords{}).Where("id = ?", rechargeRecords.ID).Updates(&rechargeRecords)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("更新充值记录失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应充值记录")
		global.GVA_LOG.Error("未找到对应充值记录!", zap.Error(err))
		tx.Rollback()
		return err
	}

	changedAmount := getChangeAmount(fund.Recharge, rechargeRecords.ExchangedAmountUsdt)
	beforeBalance := userfundAccount.Balance // 保存变动前的余额
	afterBalance := beforeBalance.Add(rechargeRecords.ExchangedAmountUsdt)
	frozenBalanceBefore := userfundAccount.FrozenBalance
	frozenBalanceAfter := userfundAccount.FrozenBalance  //不涉及冻结金额变化的情况下二者相同
	avaBalanceBefore := userfundAccount.AvailableBalance // 保存变动前的余额
	avaBalanceAfter := avaBalanceBefore.Add(rechargeRecords.ExchangedAmountUsdt)
	// 创建新的流水记录
	userAccountFlow := userfund.UserAccountFlow{
		GVA_MODEL:           global.GVA_MODEL{},
		UserId:              userfundAccount.UserId,
		TransactionType:     string(fund.Recharge),
		Amount:              rechargeRecords.ExchangedAmountUsdt,
		RootId:              rechargeRecords.RootId,
		ParentId:            rechargeRecords.ParentId,
		BalanceBefore:       beforeBalance,
		BalanceAfter:        afterBalance,
		FrozenBalanceBefore: frozenBalanceBefore,
		FrozenBalanceAfter:  frozenBalanceAfter,
		AvaBalanceAfter:     avaBalanceAfter,
		AvaBalanceBefore:    avaBalanceBefore,
		TransactionDate:     time.Now(),
		Description:         "",
		OrderId:             rechargeRecords.OrderId,
		UserType:            userfundAccount.UserType,
	}
	err := tx.Model(&userfund.UserAccountFlow{}).Create(&userAccountFlow).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	userfundAccount.Balance = afterBalance
	avaBalance := userfundAccount.AvailableBalance.Add(changedAmount)
	userfundAccount.AvailableBalance = avaBalance
	userfundAccount.UpdatedAt = func() *time.Time {
		now := time.Now()
		return &now
	}()
	//如果是充值
	if userfundAccount.FirstChargeAmount.Decimal.Equal(Zero) || userfundAccount.FirstChargeTime.IsZero() {
		userfundAccount.FirstChargeAmount = NullDecimal{Decimal: changedAmount,
			Valid: true}
		now := time.Now()
		userfundAccount.FirstChargeTime = &now
	}
	result = tx.Model(&userfund.UserFundAccounts{}).Where("id = ?", userfundAccount.Id).Updates(&userfundAccount)

	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应用户的资金账户记录")
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(err))
		tx.Rollback()
		return err
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 如果提交事务失败，回滚事务并返回错误
		tx.Rollback()
		return err
	}
	return nil
}

func (userFundAccountsService *UserFundAccountsService) UpdateUserFundAccountsWithFlowAndWithdrawRecords(withdrawRecord userfund.WithdrawRecords) error {
	// 开启事务
	tx := global.GVA_DB.Begin()
	if tx.Error != nil {
		// 如果开启事务失败，则返回错误
		return tx.Error
	}
	var rowsAffected int64
	userId := withdrawRecord.MemberId
	var userfundAccount userfund.UserFundAccounts
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ?", userId).
		First(&userfundAccount)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}
	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应用户的资金账户记录")
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(err))
		tx.Rollback()
		return err
	}
	withdrawRecord.UserType = userfundAccount.UserType
	//更新订单表
	result = tx.Model(&userfund.WithdrawRecords{}).Where("id = ?", withdrawRecord.ID).Where("order_status = ? OR order_status = ?", enums.WITHDRAWING, enums.WITHDRAW_CHECKED).Updates(&withdrawRecord)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("更新订单表失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应订单记录")
		global.GVA_LOG.Error("未找到对应订单记录!", zap.Error(err))
		tx.Rollback()
		return err
	}

	//changedAmount := getChangeAmount(fund.Withdraw, withdrawRecord.ExchangedAmountTarget)

	beforeBalance := userfundAccount.Balance // 保存变动前的余额
	afterBalance := beforeBalance.Sub(withdrawRecord.WithdrawAmount)
	avaBalanceBefore := userfundAccount.AvailableBalance

	frozenBalanceBefore := userfundAccount.FrozenBalance
	frozenBalanceAfter := userfundAccount.FrozenBalance.Sub(withdrawRecord.WithdrawAmount) // 冻结金额减少
	userfundAccount.FrozenBalance = frozenBalanceAfter
	userfundAccount.Balance = userfundAccount.Balance.Sub(withdrawRecord.WithdrawAmount) //总金额减少
	userfundAccount.UpdatedAt = func() *time.Time {
		now := time.Now()
		return &now
	}()
	avaBalanceAfter := afterBalance.Sub(frozenBalanceAfter)
	commission := withdrawRecord.Commission.Add(withdrawRecord.ThirdCommission)
	// 创建新的流水记录
	userAccountFlow := userfund.UserAccountFlow{
		GVA_MODEL:           global.GVA_MODEL{},
		UserId:              userfundAccount.UserId,
		TransactionType:     string(fund.Withdraw),
		Amount:              withdrawRecord.WithdrawAmount,
		TotalCommission:     commission,
		BalanceBefore:       beforeBalance,
		BalanceAfter:        afterBalance,
		FrozenBalanceBefore: frozenBalanceBefore,
		FrozenBalanceAfter:  frozenBalanceAfter,
		AvaBalanceBefore:    avaBalanceBefore,
		AvaBalanceAfter:     avaBalanceAfter,
		TransactionDate:     time.Now(),
		Description:         "",
		OrderId:             withdrawRecord.OrderId,
		UserType:            userfundAccount.UserType,
	}
	err := tx.Model(&userfund.UserAccountFlow{}).Create(&userAccountFlow).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	result = tx.Model(&userfund.UserFundAccounts{}).Where("id = ?", userfundAccount.Id).Updates(&userfundAccount)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("更新自己账户表失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到自己账户记录")
		global.GVA_LOG.Error("未找到自己账户记录!", zap.Error(err))
		tx.Rollback()
		return err
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 如果提交事务失败，回滚事务并返回错误
		tx.Rollback()
		return err
	}
	return nil
}

func (userFundAccountsService *UserFundAccountsService) UpdateUserFundAccountsWithFlowAndRefusedWithdrawRecords(record userfund.WithdrawRecords, userAccount userfund.UserFundAccounts) error {
	// 开启事务
	tx := global.GVA_DB.Begin()
	if tx.Error != nil {
		// 如果开启事务失败，则返回错误
		return tx.Error
	}

	balance := userAccount.Balance // 保存变动前的余额
	avaBalanceBefore := userAccount.AvailableBalance
	frozenBalanceBefore := userAccount.FrozenBalance

	// 计算新的余额
	newFrozenBalance := userAccount.FrozenBalance.Sub(record.WithdrawAmount)       //原有的冻结金额需要减少
	newAvailableBalance := userAccount.AvailableBalance.Add(record.WithdrawAmount) //可用金额增加
	if newAvailableBalance.IsNegative() {
		tx.Rollback()
		return fmt.Errorf("可用金额资金计算失败")
	}
	updateAccount := userfund.UserFundAccounts{
		Id:               userAccount.Id,
		AvailableBalance: newAvailableBalance,
		FrozenBalance:    newFrozenBalance,
	}
	result := tx.Model(&userfund.UserFundAccounts{}).Where("id = ?", updateAccount.Id).Updates(&updateAccount)

	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}
	var rowsAffected int64
	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应用户的资金账户记录")
		global.GVA_LOG.Error("查询用户资金失败!", zap.Error(err))
		tx.Rollback()
		return err
	}

	approvalTime := time.Now()
	record.ApprovalTime = &approvalTime
	record.OrderStatus = enums.WITHDRAWED
	// 更新提现记录
	result = tx.Model(&userfund.WithdrawRecords{}).Where("id = ?", record.ID).Updates(&record)
	// 检查错误
	if result.Error != nil {
		global.GVA_LOG.Error("更新提现记录失败!", zap.Error(result.Error))
		tx.Rollback()
		return result.Error
	}

	// 检查是否有记录被查询到
	rowsAffected = result.RowsAffected
	if rowsAffected == 0 {
		err := errors.New("未找到对应的提现记录")
		global.GVA_LOG.Error("查询提现记录失败!", zap.Error(err))
		tx.Rollback()
		return err
	}

	// 创建新的流水记录
	userAccountFlow := userfund.UserAccountFlow{
		GVA_MODEL:           global.GVA_MODEL{},
		UserId:              updateAccount.UserId,
		TransactionType:     string(fund.Withdraw),
		Amount:              record.WithdrawAmount,
		BalanceBefore:       balance,
		BalanceAfter:        balance,
		FrozenBalanceBefore: frozenBalanceBefore,
		FrozenBalanceAfter:  newFrozenBalance,
		AvaBalanceBefore:    avaBalanceBefore,
		AvaBalanceAfter:     newAvailableBalance,
		TransactionDate:     time.Now(),
		Description:         "",
		OrderId:             record.OrderId,
		UserType:            userAccount.UserType,
	}
	err := tx.Model(&userfund.UserAccountFlow{}).Create(&userAccountFlow).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 如果提交事务失败，回滚事务并返回错误
		tx.Rollback()
		return err
	}
	return nil
}

func (userFundAccountsService *UserFundAccountsService) GetUserFundAccountsByUserIdWithtx(userId int, tx *gorm.DB) (userFundAccounts *userfund.UserFundAccounts, err error) {
	err = tx.Model(&userfund.UserFundAccounts{}).Where("user_id = ?", userId).First(&userFundAccounts).Error
	if err != nil {
		return nil, err
	}
	return userFundAccounts, nil
}

type Option func(*UserFundAccountsService)

func WithFrozenChange(n Decimal) Option {
	return func(c *UserFundAccountsService) {
		c.FrozenChange = n
	}
}

func WithAvailableChange(n Decimal) Option {
	return func(c *UserFundAccountsService) {
		c.AvailableChange = n
	}
}

func WithTotalBalanceChange(n Decimal) Option {
	return func(c *UserFundAccountsService) {
		c.TotalChange = n
	}
}
