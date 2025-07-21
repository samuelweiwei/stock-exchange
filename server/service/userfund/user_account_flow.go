package userfund

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	userfundReq "github.com/flipped-aurora/gin-vue-admin/server/model/userfund/request"
	"gorm.io/gorm"
)

type UserAccountFlowService struct {
	UserFundAccountsService *UserFundAccountsService
}

func NewUserAccountFlowService(tx *gorm.DB, autoCommit bool) *UserAccountFlowService {
	// 在构造函数中初始化 UserFundService
	return &UserAccountFlowService{
		UserFundAccountsService: &UserFundAccountsService{}, // 自动初始化 UserFundService
	}
}

// CreateUserAccountFlow 创建userAccountFlow表记录
// Author [yourname](https://github.com/yourname)
func (userAccountFlowService *UserAccountFlowService) CreateUserAccountFlow(userAccountFlow *userfund.UserAccountFlow) (err error) {
	err = global.GVA_DB.Create(userAccountFlow).Error
	return err
}

// DeleteUserAccountFlow 删除userAccountFlow表记录
// Author [yourname](https://github.com/yourname)
func (userAccountFlowService *UserAccountFlowService) DeleteUserAccountFlow(ID string) (err error) {
	err = global.GVA_DB.Delete(&userfund.UserAccountFlow{}, "id = ?", ID).Error
	return err
}

// DeleteUserAccountFlowByIds 批量删除userAccountFlow表记录
// Author [yourname](https://github.com/yourname)
func (userAccountFlowService *UserAccountFlowService) DeleteUserAccountFlowByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]userfund.UserAccountFlow{}, "id in ?", IDs).Error
	return err
}

// UpdateUserAccountFlow 更新userAccountFlow表记录
// Author [yourname](https://github.com/yourname)
func (userAccountFlowService *UserAccountFlowService) UpdateUserAccountFlow(userAccountFlow userfund.UserAccountFlow) (err error) {
	err = global.GVA_DB.Model(&userfund.UserAccountFlow{}).Where("id = ?", userAccountFlow.ID).Updates(&userAccountFlow).Error
	return err
}

// GetUserAccountFlow 根据ID获取userAccountFlow表记录
// Author [yourname](https://github.com/yourname)
func (userAccountFlowService *UserAccountFlowService) GetUserAccountFlow(ID string) (userAccountFlow userfund.UserAccountFlow, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&userAccountFlow).Error
	return
}

// GetUserAccountFlowInfoList 分页获取userAccountFlow表记录
// Author [yourname](https://github.com/yourname)
func (userAccountFlowService *UserAccountFlowService) GetUserAccountFlowInfoList(info userfundReq.UserAccountFlowSearch) (list []userfund.UserAccountFlowUnion, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&userfund.UserAccountFlow{})

	// 时间范围查询
	if info.StartTime > 0 && info.EndTime > 0 {
		startTime := time.UnixMilli(info.StartTime)
		endTime := time.UnixMilli(info.EndTime)
		db = db.Where("user_account_flow.transaction_date BETWEEN ? AND ?", startTime, endTime)
	}

	// 其他查询条件
	if info.UserId > 0 {
		db = db.Where("user_account_flow.user_id = ?", info.UserId)
	}
	if info.TransactionType != "" {
		db = db.Where("user_account_flow.transaction_type = ?", info.TransactionType)
	}
	if info.UserType > 0 {
		db = db.Where("user_account_flow.user_type = ?", info.UserType)
	}

	if info.OrderId != "" {
		db = db.Where("user_account_flow.order_id = ?", info.OrderId)
	}
	if info.ParentId > 0 {
		db = db.Where("user_account_flow.parent_id = ?", info.ParentId)
	}
	if info.RootUserId > 0 {
		db = db.Where("user_account_flow.root_id = ?", info.RootUserId)
	}

	// 关联用户表查询
	db = db.Joins("JOIN frontend_users ON frontend_users.id = user_account_flow.user_id").
		Select("user_account_flow.*, frontend_users.phone as phone_number, frontend_users.email, frontend_users.username")

	// 用户信息查询
	if info.PhoneNumber != "" {
		db = db.Where("frontend_users.phone = ?", info.PhoneNumber)
	}
	if info.Email != "" {
		db = db.Where("frontend_users.email = ?", info.Email)
	}
	if info.UserName != "" {
		db = db.Where("frontend_users.username like ?", "%"+info.UserName+"%")
	}

	// 按时间倒序
	db = db.Order("transaction_date desc")

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 分页
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&list).Error
	return list, total, err
}
func (userAccountFlowService *UserAccountFlowService) GetUserAccountFlowPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
