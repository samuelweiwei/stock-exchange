package userfund

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"github.com/shopspring/decimal"
	"time"
)

type RechargeRecordsService struct{}

// CreateRechargeRecords 创建rechargeRecords表记录
// Author [yourname](https://github.com/yourname)
func (rechargeRecordsService *RechargeRecordsService) CreateRechargeRecords(rechargeRecords *userfund.RechargeRecords) (err error) {
	err = global.GVA_DB.Create(rechargeRecords).Error
	return err
}

// DeleteRechargeRecords 删除rechargeRecords表记录
// Author [yourname](https://github.com/yourname)
func (rechargeRecordsService *RechargeRecordsService) DeleteRechargeRecords(ID string) (err error) {
	err = global.GVA_DB.Delete(&userfund.RechargeRecords{}, "id = ?", ID).Error
	return err
}

// DeleteRechargeRecordsByIds 批量删除rechargeRecords表记录
// Author [yourname](https://github.com/yourname)
func (rechargeRecordsService *RechargeRecordsService) DeleteRechargeRecordsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.RechargeRecords{}, "id in ?", IDs).Error
	return err
}

// UpdateRechargeRecords 更新rechargeRecords表记录
// Author [yourname](https://github.com/yourname)
func (rechargeRecordsService *RechargeRecordsService) UpdateRechargeRecords(rechargeRecords userfund.RechargeRecords) (err error) {
	err = global.GVA_DB.Model(&userfund.RechargeRecords{}).Where("id = ?", rechargeRecords.ID).Updates(&rechargeRecords).Error
	return err
}

// GetRechargeRecords 根据ID获取rechargeRecords表记录
// Author [yourname](https://github.com/yourname)
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecords(ID string) (rechargeRecords userfund.RechargeRecords, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&rechargeRecords).Error
	return
}

// GetRechargeRecords 根据ID获取rechargeRecords表记录
// Author [yourname](https://github.com/yourname)
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordsByOrderId(ID string) (rechargeRecords userfund.RechargeRecords, err error) {
	err = global.GVA_DB.Where("order_id = ?", ID).First(&rechargeRecords).Error
	return
}

// GetRechargeRecordsInfoList 分页获取RechargeRecords表记录
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordsInfoList(info userfundReq.RechargeRecordsSearch) (list []userfund.RechargeRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userfund.RechargeRecords{})
	var rechargeRecordss []userfund.RechargeRecords
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	// 时间范围查询
	if info.StartRechargeTime > 0 && info.EndRechargeTime > 0 {
		startTime := time.UnixMilli(info.StartRechargeTime)
		endTime := time.UnixMilli(info.EndRechargeTime)
		db = db.Where("recharge_time BETWEEN ? AND ?", startTime, endTime)
	}
	// 时间范围查询
	if info.StartUpdateTime > 0 && info.EndUpdateTime > 0 {
		startTime := time.UnixMilli(info.StartUpdateTime)
		endTime := time.UnixMilli(info.EndUpdateTime)
		db = db.Where("updated_at BETWEEN ? AND ?", startTime, endTime)
	}

	if info.RechargeType != "" {
		db = db.Where("recharge_type = ?", info.RechargeType)
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
	if info.UserType != 0 {
		db = db.Where("user_type = ?", info.UserType)
	}
	if info.OrderStatus != "" {
		db = db.Where("order_status = ?", info.OrderStatus)
	}
	// 添加地址搜索条件
	if info.FromAddress != "" {
		db = db.Where("from_address = ?", info.FromAddress)
	}
	if info.ToAddress != "" {
		db = db.Where("to_address = ?", info.ToAddress)
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

	// 添加默认的创建时间倒序排序
	err = db.Order("created_at desc").Find(&rechargeRecordss).Error
	return rechargeRecordss, total, err
}
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordsPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (rechargeRecordsService *RechargeRecordsService) CountAllByUserId(userId int64) decimal.Decimal {
	// 创建db
	db := global.GVA_DB.Model(&userfund.RechargeRecords{})
	db = db.Where("member_id = ?", userId).Where("order_status = ?", "2")
	//var sum struct{ TotalAmount decimal.Decimal }
	var totalAmount decimal.Decimal
	var totalAmountStr string

	rows, err := db.Select("IFNULL(SUM(exchanged_amount_usdt), 0) as total_amount").Rows()
	if err != nil {
		return decimal.Zero // 处理错误
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&totalAmountStr)
		if err != nil {
			return decimal.Zero // 处理扫描错误
		}
	} else {
		// 处理没有查询到数据的情况
		totalAmountStr = "0" // 或者其他默认值
	}
	totalAmount, err2 := decimal.NewFromString(totalAmountStr)
	if err2 != nil {
		fmt.Println("Decimal conversion error:", err)
		return decimal.Zero
	}
	return totalAmount
}

// GetRechargeRecordsByUserId 根据用户ID获取充值记录
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordsByUserId(userId int64) ([]userfund.RechargeRecords, error) {
	var records []userfund.RechargeRecords
	// 添加默认的创建时间倒序排序
	err := global.GVA_DB.Where("member_id = ?", userId).Order("created_at desc").Find(&records).Error
	return records, err
}

// GetRechargeRecordsByUserIdWithPagination 根据用户ID获取充值记录（带分页）
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordsByUserIdWithPagination(userId int64, page int, pageSize int) (list []userfund.RechargeRecords, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)

	// 查询总数
	db := global.GVA_DB.Model(&userfund.RechargeRecords{})
	db = db.Where("member_id = ?", userId)
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询分页数据
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

// GetRechargeRecordDetail 获取用户充值记录详情
func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordDetail(userId int64, recordId string) (rechargeRecords userfund.RechargeRecords, err error) {
	// 查询记录并确保属于当前用户
	err = global.GVA_DB.Where("id = ? AND member_id = ?", recordId, userId).First(&rechargeRecords).Error
	if err != nil {
		return rechargeRecords, err
	}
	return rechargeRecords, nil
}

func (rechargeRecordsService *RechargeRecordsService) GetRechargeRecordsInfoListByRootUser(info userfundReq.RechargeRecordsSearch) ([]userfund.RechargeRecords, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&userfund.RechargeRecords{})
	var rechargeRecordss []userfund.RechargeRecords
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
    FROM recharge_records o
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
    FROM recharge_records o
    INNER JOIN user_hierarchy uh ON o.member_id = uh.id
	WHERE 1=1
    `
	// 添加其他查询条件
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		recursiveQuery += " AND o.created_at BETWEEN ? AND ?"
		recursiveQueryCount += " AND o.created_at BETWEEN ? AND ?"
	}
	if info.RechargeType != "" {
		recursiveQuery += " AND o.recharge_type = ?"
		recursiveQueryCount += " AND o.recharge_type = ?"
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
	err := db.Raw(recursiveQueryCount, append([]interface{}{info.MemberId}, getQueryParams(info)...)...).Scan(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果 limit 不为 0，则设置分页
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 获取具体的充值记录
	err = db.Raw(recursiveQuery, append([]interface{}{info.MemberId}, getQueryParams(info)...)...).Limit(limit).Offset(offset).Scan(&rechargeRecordss).Error
	return rechargeRecordss, total, err
}

func getQueryParams(info userfundReq.RechargeRecordsSearch) []interface{} {
	var params []interface{}

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		params = append(params, info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RechargeType != "" {
		params = append(params, info.RechargeType)
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
