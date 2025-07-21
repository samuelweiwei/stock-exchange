package userfund

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/enums/fund"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	. "github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

type WithdrawRecordsService struct{}

// CreateWithdrawRecords 创建withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) CreateWithdrawRecords(withdrawRecords *userfund.WithdrawRecords) (err error) {
	err = global.GVA_DB.Create(withdrawRecords).Error
	return err
}

// DeleteWithdrawRecords 删除withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) DeleteWithdrawRecords(ID string) (err error) {
	err = global.GVA_DB.Delete(&userfund.WithdrawRecords{}, "id = ?", ID).Error
	return err
}

// DeleteWithdrawRecordsByIds 批量删除withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) DeleteWithdrawRecordsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.WithdrawRecords{}, "id in ?", IDs).Error
	return err
}

// UpdateWithdrawRecords 更新withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) UpdateWithdrawRecords(withdrawRecords userfund.WithdrawRecords) (err error) {
	err = global.GVA_DB.Model(&userfund.WithdrawRecords{}).Where("id = ?", withdrawRecords.ID).Updates(&withdrawRecords).Error
	return err
}

// GetWithdrawRecords 根据ID获取withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecords(ID string) (withdrawRecords userfund.WithdrawRecords, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&withdrawRecords).Error
	return
}

// GetWithdrawRecordsInfoList 分页获取withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordsInfoList(info userfundReq.WithdrawRecordsSearch) (list []userfund.WithdrawRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userfund.WithdrawRecords{})
	var withdrawRecordss []userfund.WithdrawRecords
	// 如果有条件搜索 下方会自动创建搜索语句
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 时间范围查询
	if info.StartWithdrawTime > 0 && info.EndWithdrawTime > 0 {
		startTime := time.UnixMilli(info.StartWithdrawTime)
		endTime := time.UnixMilli(info.EndWithdrawTime)
		db = db.Where("withdraw_time BETWEEN ? AND ?", startTime, endTime)
	}
	// 时间范围查询
	if info.StartUpdateTime > 0 && info.EndUpdateTime > 0 {
		startTime := time.UnixMilli(info.StartUpdateTime)
		endTime := time.UnixMilli(info.EndUpdateTime)
		db = db.Where("updated_at BETWEEN ? AND ?", startTime, endTime)
	}
	if info.WithdrawType != "" {
		db = db.Where("withdraw_type = ?", info.WithdrawType)
	}
	if info.OrderId != "" {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.MemberId > 0 {
		db = db.Where("member_id = ?", info.MemberId)
	}
	if info.MemberPhone != "" {
		db = db.Where("member_phone = ?", info.MemberPhone)
	}
	if info.MemberEmail != "" {
		db = db.Where("member_email = ?", info.MemberEmail)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.ReviewStatus != "" {
		db = db.Where("review_status = ?", info.ReviewStatus)
	}
	if info.FromAddress != "" {
		db = db.Where("from_address = ?", info.FromAddress)
	}
	if info.ToAddress != "" {
		db = db.Where("to_address = ?", info.ToAddress)
	}
	if info.OrderStatus != "" {
		db = db.Where("order_status = ?", info.OrderStatus)
	}
	if info.UserType != 0 {
		db = db.Where("user_type = ?", info.UserType)
	}
	if info.ParentId > 0 {
		db = db.Where("parent_id = ?", info.ParentId)
	}
	if info.RootId > 0 {
		db = db.Where("root_id = ?", info.RootId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("created_at desc").Find(&withdrawRecordss).Error
	return withdrawRecordss, total, err
}
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetwithdrawRecords 根据ID获取withdrawRecords表记录
// Author [yourname](https://github.com/yourname)
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordsByOrderId(ID string) (withdrawRecords userfund.WithdrawRecords, err error) {
	err = global.GVA_DB.Where("order_id = ?", ID).First(&withdrawRecords).Error
	return
}

func (withdrawRecordsService *WithdrawRecordsService) CountAllByUserId(userId int64) Decimal {
	// 创建db
	db := global.GVA_DB.Model(&userfund.WithdrawRecords{})
	db = db.Where("member_id = ?", userId).Where("order_status = ?", "2")
	var totalAmount Decimal
	var totalAmountStr string
	rows, err := db.Select("IFNULL(SUM(withdraw_amount), 0) as total_amount").Rows()
	if err != nil {
		return Zero // 处理错误
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&totalAmountStr)
		if err != nil {
			return Zero // 处理扫描错误
		}
	} else {
		// 处理没有查询到数据的情况
		totalAmountStr = "0" // 或者其他默认值
	}
	totalAmount, err2 := NewFromString(totalAmountStr)
	if err2 != nil {
		fmt.Println("Decimal conversion error:", err)
		return Zero
	}
	return totalAmount
}

// GetWithdrawRecordsByUserId 根据用户ID获取提现记录
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordsByUserId(userId int64) ([]userfund.WithdrawRecords, error) {
	var records []userfund.WithdrawRecords
	err := global.GVA_DB.Where("member_id = ?", userId).Order("created_at desc").Find(&records).Error
	return records, err
}

// GetWithdrawRecordsByUserIdWithPagination 根据用户ID获取提现记录（带分页）
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordsByUserIdWithPagination(userId int64, page int, pageSize int) (list []userfund.WithdrawRecords, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)

	// 查询总数
	db := global.GVA_DB.Model(&userfund.WithdrawRecords{})
	db = db.Where("member_id = ?", userId)
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询分页数据
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&list).Error

	return list, total, err
}

// GetWithdrawRecordDetail 获取用户提现记录详情
func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordDetail(userId int64, recordId string) (withdrawRecords userfund.WithdrawRecords, err error) {
	// 查询记录并确保属于当前用户
	err = global.GVA_DB.Where("id = ? AND member_id = ?", recordId, userId).First(&withdrawRecords).Error
	if err != nil {
		return withdrawRecords, err
	}
	return withdrawRecords, nil
}

func (withdrawRecordsService *WithdrawRecordsService) CreateWithdrawRecordsAndFrozenBalance(withdrawRecords userfund.WithdrawRecords, userfundAccount userfund.UserFundAccounts) error {
	// 开启事务
	tx := global.GVA_DB.Begin()
	if tx.Error != nil {
		// 如果开启事务失败，则返回错误
		return tx.Error
	}
	err := tx.Model(&userfund.WithdrawRecords{}).Create(&withdrawRecords).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	frozenBalanceBefore := userfundAccount.FrozenBalance
	avaBalanceBefore := userfundAccount.AvailableBalance
	newAvailableBalance := userfundAccount.AvailableBalance.Sub(withdrawRecords.WithdrawAmount)
	var newFrozenBalance Decimal

	if userfundAccount.FrozenBalance == Zero {
		newFrozenBalance = withdrawRecords.WithdrawAmount
	} else {
		newFrozenBalance = userfundAccount.FrozenBalance.Add(withdrawRecords.WithdrawAmount)
	}
	// 创建更新对象
	updateAccount := userfund.UserFundAccounts{
		Id:               userfundAccount.Id,
		AvailableBalance: newAvailableBalance,
		FrozenBalance:    newFrozenBalance,
	}
	result := tx.Model(&userfund.UserFundAccounts{}).Where("id = ?", updateAccount.Id).Updates(&updateAccount)
	var rowsAffected int64

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
	//更新流水
	userAccountFlow := userfund.UserAccountFlow{
		GVA_MODEL:           global.GVA_MODEL{},
		UserId:              userfundAccount.UserId,
		TransactionType:     string(fund.ApplyWithdraw),
		Amount:              withdrawRecords.WithdrawAmount,
		BalanceBefore:       userfundAccount.Balance,
		BalanceAfter:        userfundAccount.Balance,
		FrozenBalanceBefore: frozenBalanceBefore,
		FrozenBalanceAfter:  newFrozenBalance,
		AvaBalanceBefore:    avaBalanceBefore,
		AvaBalanceAfter:     newAvailableBalance,
		TransactionDate:     time.Now(),
		Description:         "",
		OrderId:             withdrawRecords.OrderId,
		UserType:            userfundAccount.UserType,
	}
	err = tx.Model(&userfund.UserAccountFlow{}).Create(&userAccountFlow).Error
	if err != nil {
		global.GVA_LOG.Error("新增用户流水失败!", zap.Error(err))
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

func (withdrawRecordsService *WithdrawRecordsService) GetWithdrawRecordsInfoListByRootUser(info userfundReq.WithdrawRecordsSearch) ([]userfund.WithdrawRecords, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userfund.WithdrawRecords{})
	var withdrawRecordss []userfund.WithdrawRecords
	var total int64

	// 构建递归查询的SQL
	recursiveQuery := `
    WITH RECURSIVE user_hierarchy AS (
        SELECT id, root_userid
        FROM frontend_users
        WHERE id = ?
        
        UNION ALL
        
        SELECT u.id, u.root_userid
        FROM frontend_users u
        INNER JOIN user_hierarchy uh ON u.root_userid = uh.id
    )
    SELECT o.*
    FROM withdraw_records o
    INNER JOIN user_hierarchy uh ON o.member_id = uh.id
	WHERE 1=1
    `

	recursiveQueryCount := `
    WITH RECURSIVE user_hierarchy AS (
        SELECT id, root_userid
        FROM frontend_users
        WHERE id = ?
        
        UNION ALL
        
        SELECT u.id, u.root_userid
        FROM frontend_users u
        INNER JOIN user_hierarchy uh ON u.root_userid = uh.id
    )
    SELECT count(*)
    FROM withdraw_records o
    INNER JOIN user_hierarchy uh ON o.member_id = uh.id
	WHERE 1=1
    `
	// 添加其他查询条件
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		recursiveQuery += " AND o.created_at BETWEEN ? AND ?"
		recursiveQueryCount += " AND o.created_at BETWEEN ? AND ?"
	}
	if info.WithdrawType != "" {
		recursiveQuery += " AND o.withdraw_type = ?"
		recursiveQueryCount += " AND o.withdraw_type = ?"
	}
	if info.OrderId != "" {
		recursiveQuery += " AND o.order_id = ?"
		recursiveQueryCount += " AND o.order_id = ?"
	}

	if info.MemberPhone != "" {
		recursiveQuery += " AND o.member_phone = ?"
		recursiveQueryCount += " AND o.member_phone = ?"
	}
	if info.MemberEmail != "" {
		recursiveQuery += " AND o.member_email = ?"
		recursiveQueryCount += " AND o.member_email = ?"
	}
	if info.Currency != "" {
		recursiveQuery += " AND o.currency = ?"
		recursiveQueryCount += " AND o.currency = ?"
	}
	if info.ReviewStatus != "" {
		recursiveQuery += " AND o.review_status = ?"
		recursiveQueryCount += " AND o.review_status = ?"
	}
	if info.OrderStatus != "" {
		recursiveQuery += " AND o.order_status = ?"
		recursiveQueryCount += " AND o.order_status = ?"
	}
	if info.FromAddress != "" {
		recursiveQuery += " AND o.from_address = ?"
		recursiveQueryCount += " AND o.from_address = ?"
	}
	if info.ToAddress != "" {
		recursiveQuery += " AND o.to_address = ?"
		recursiveQueryCount += " AND o.to_address = ?"
	}

	// 执行递归查询，获取符合条件的记录总数
	//err := db.Raw(recursiveQueryCount, info.MemberId).Scan(&total).Error
	err := db.Raw(recursiveQueryCount, append([]interface{}{info.MemberId}, getQueryParams2(info)...)...).Scan(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果 limit 不为 0，则设置分页
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 获取具体的充值记录
	err = db.Raw(recursiveQuery, append([]interface{}{info.MemberId}, getQueryParams2(info)...)...).Limit(limit).Offset(offset).Scan(&withdrawRecordss).Error
	return withdrawRecordss, total, err
}

func getQueryParams2(info userfundReq.WithdrawRecordsSearch) []interface{} {
	var params []interface{}

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		params = append(params, info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.WithdrawType != "" {
		params = append(params, info.WithdrawType)
	}
	if info.OrderId != "" {
		params = append(params, info.OrderId)
	}
	if info.MemberPhone != "" {
		params = append(params, info.MemberPhone)
	}
	if info.MemberEmail != "" {
		params = append(params, info.MemberEmail)
	}
	if info.Currency != "" {
		params = append(params, info.Currency)
	}
	if info.ReviewStatus != "" {
		params = append(params, info.ReviewStatus)
	}
	if info.OrderStatus != "" {
		params = append(params, info.OrderStatus)
	}
	if info.FromAddress != "" {
		params = append(params, info.FromAddress)
	}
	if info.ToAddress != "" {
		params = append(params, info.ToAddress)
	}

	return params
}
