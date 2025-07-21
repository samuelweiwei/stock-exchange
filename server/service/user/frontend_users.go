package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	userRes "github.com/flipped-aurora/gin-vue-admin/server/model/user/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FrontendUsersService struct{}

// CreateFrontendUsers 创建frontendUsers表记录
// Author [yourname](https://github.com/yourname)
func (frontendUsersService *FrontendUsersService) CreateFrontendUsers(frontendUsers *user.FrontendUsers) (err error) {
	frontendUsers.Password = utils.BcryptHash(frontendUsers.Password)
	frontendUsers.PaymentPassword = utils.BcryptHash(frontendUsers.PaymentPassword)
	err = global.GVA_DB.Create(frontendUsers).Error
	return err
}

// DeleteFrontendUsers 删除frontendUsers表记录
// Author [yourname](https://github.com/yourname)
func (frontendUsersService *FrontendUsersService) DeleteFrontendUsers(ID string) (err error) {
	err = global.GVA_DB.Delete(&user.FrontendUsers{}, "id = ?", ID).Error
	return err
}

// DeleteFrontendUsersByIds 批量删除frontendUsers表记录
// Author [yourname](https://github.com/yourname)
func (frontendUsersService *FrontendUsersService) DeleteFrontendUsersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]user.FrontendUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateFrontendUsers 更新frontendUsers表记录
// Author [yourname](https://github.com/yourname)
func (frontendUsersService *FrontendUsersService) UpdateFrontendUsers(frontendUsers user.FrontendUsers) (err error) {
	if frontendUsers.Password != "" {
		frontendUsers.Password = utils.BcryptHash(frontendUsers.Password)
	}
	if frontendUsers.PaymentPassword != "" {
		frontendUsers.PaymentPassword = utils.BcryptHash(frontendUsers.PaymentPassword)
	}
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", frontendUsers.ID).Updates(&frontendUsers).Error
	return err
}

// GetFrontendUsers 根据ID获取frontendUsers表记录
// Author [yourname](https://github.com/yourname)
func (frontendUsersService *FrontendUsersService) GetFrontendUsers(ID string) (frontendUsers user.FrontendUsers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&frontendUsers).Error
	return
}

func (frontendUsersService *FrontendUsersService) DelUserCache(ID string) (err error) {
	// 空的context
	userKey := fmt.Sprintf(constants.UserInfoKEY, ID)
	err = global.GVA_REDIS.Del(context.Background(), userKey).Err()
	return
}

// GetFrontendUsersInfoList 分页获取frontendUsers表记录
// Author [yourname](https://github.com/yourname)
func (frontendUsersService *FrontendUsersService) GetFrontendUsersInfoList(info userReq.FrontendUsersSearch) (list []user.FrontendUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.FrontendUsers{})
	var frontendUserss []user.FrontendUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserIDList != "" {
		// split ,cast to int
		userIDStrList := strings.Split(info.UserIDList, ",")
		// 转换成 uint
		userIDList := make([]uint, len(userIDStrList))
		for _, v := range userIDStrList {
			userIDList = append(userIDList, cast.ToUint(v))
		}
		db = db.Where("id in (?)", userIDList)
	}

	if info.ID != 0 {
		db = db.Where("id = ?", info.ID)
	}

	if info.UserType != 0 {
		db = db.Where("user_type = ?", info.UserType)
	}

	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.Email != "" {
		db = db.Where("email = ?", info.Email)
	}
	if info.CountryId != 0 {
		db = db.Where("country_id = ?", info.CountryId)
	}

	if info.RootId != 0 {
		db = db.Where("RootUserid = ?", info.RootId)
	}
	// 根用户
	if info.RootUserid != 0 {
		db = db.Where("root_userid = ?", info.RootUserid)
	}

	// 上级代理
	if info.ParentId != 0 {
		db = db.Where("parent_id = ?", info.ParentId)
	}

	// 注册时间
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	// 会员状态
	if info.Enable != 0 {
		db = db.Where("enable = ?", info.Enable)
	}

	// 最近登录IP
	if info.LastLoginIp != "" {
		db = db.Where("last_login_ip = ?", info.LastLoginIp)
	}

	if info.AuthenticationStatus != nil {
		db = db.Where("authentication_status = ?", info.AuthenticationStatus)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("updated_at desc").Find(&frontendUserss).Error
	return frontendUserss, total, err
}
func (frontendUsersService *FrontendUsersService) GetFrontendUsersPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// BindEmail 绑定邮箱
func (frontendUsersService *FrontendUsersService) BindEmail(userID uint, email string) (err error) {
	return global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", userID).Update("email", email).Error
}

func (frontendUsersService *FrontendUsersService) BindPhone(userID uint, req *userReq.BindPhoneReq) (err error) {
	return global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"country_id": req.CountryId,
		"phone":      req.Phone,
	}).Error
}

// RealNameAuthentication 实名审核
func (frontendUsersService *FrontendUsersService) RealNameAuthentication(userID uint, status int) (err error) {
	return global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", userID).Update("authentication_status", status).Error
}

// ChangeParent 修改上级代理
func (frontendUsersService *FrontendUsersService) ChangeParent(req *userReq.ChangeParentReq) (err error) {
	userInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(req.ID))
	if err != nil {
		return
	}
	userParentInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(req.ParentID))
	if err != nil {
		return
	}

	return frontendUsersService.changeRelation(userInfo, userParentInfo)
}

func (frontendUsersService *FrontendUsersService) changeRelation(users user.FrontendUsers, parentUser user.FrontendUsers) (err error) {
	tx := global.GVA_DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	rootUserID := parentUser.RootUserid
	if rootUserID == 0 {
		rootUserID = parentUser.ID
	}
	// 1. 修改当用户的上级、 上上级、 上上上级
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", users.ID).Updates(map[string]interface{}{
		"parent_id":            parentUser.ID,
		"grandparent_id":       parentUser.ParentId,
		"great_grandparent_id": parentUser.GrandparentId,
		"root_userid":          rootUserID,
	}).Error

	if err != nil {
		return
	}

	// 2. 修改当前用户的下级的 上上级、上上上级
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("parent_id = ?", users.ID).Updates(map[string]interface{}{
		"grandparent_id":       parentUser.ID,
		"great_grandparent_id": parentUser.ParentId,
		"root_userid":          rootUserID,
	}).Error

	if err != nil {
		return
	}

	// 3. 修改当前用户的下下级的 上上上级
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("grandparent_id = ?", users.ID).Updates(map[string]interface{}{
		"great_grandparent_id": parentUser.ID,
		"root_userid":          rootUserID,
	}).Error

	if err != nil {
		return
	}

	// 提交事务
	err = tx.Commit().Error

	return
}

// GetAncestors 获取上三级代理
func (frontendUsersService *FrontendUsersService) GetAncestors(req *userReq.GetAncestorsReq) (res []uint, err error) {
	res = make([]uint, 0)
	userInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(req.ID))
	if err != nil {
		return
	}
	if userInfo.ParentId == 0 {
		return
	}
	res = append(res, userInfo.ParentId)
	pParentInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(userInfo.ParentId))
	if err != nil {
		return
	}

	if pParentInfo.ParentId == 0 {
		return
	}
	res = append(res, pParentInfo.ParentId)

	ppParentInfo, err := frontendUsersService.GetFrontendUsers(cast.ToString(pParentInfo.ParentId))
	if err != nil {
		return
	}
	if ppParentInfo.ParentId == 0 {
		return
	}
	res = append(res, ppParentInfo.ParentId)
	return
}

func (frontendUsersService *FrontendUsersService) ChangePassword(userID uint, password, newPassword string) (userInter *user.FrontendUsers, err error) {
	var userInfo user.FrontendUsers
	if err = global.GVA_DB.Where("id = ?", userID).First(&userInfo).Error; err != nil {
		return nil, err
	}
	if password != "" {
		if ok := utils.BcryptCheck(password, userInfo.Password); !ok {
			return nil, errors.New("原密码错误")
		}
	}
	userInfo.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&userInfo).Error
	return &userInfo, err
}
func (frontendUsersService *FrontendUsersService) ChangePaymentPassword(userID uint, password, newPassword string) (userInter *user.FrontendUsers, err error) {
	var userInfo user.FrontendUsers
	if err = global.GVA_DB.Where("id = ?", userID).First(&userInfo).Error; err != nil {
		return nil, err
	}
	if password != "" {
		if ok := utils.BcryptCheck(password, userInfo.PaymentPassword); !ok {
			return nil, errors.New("原密码错误")
		}
	}
	userInfo.PaymentPassword = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&userInfo).Error
	return &userInfo, err
}

func (frontendUsersService *FrontendUsersService) Register(u user.FrontendUsers) (userInter user.FrontendUsers, err error) {
	var user user.FrontendUsers
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	if u.PaymentPassword != "" {
		u.PaymentPassword = utils.BcryptHash(u.PaymentPassword)
	}
	u.UUid = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

func (frontendUsersService *FrontendUsersService) UserIdentity(userId uint, req *userReq.UserIdentityReq) (err error) {
	return global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"country_id":            req.CountryId,
		"id_number":             req.IdNumber,
		"id_type":               req.IdType,
		"id_images":             req.IdImages,
		"real_name":             req.RealName,
		"authentication_status": constants.AuthenticationStatusPending,
	}).Error
}

func (frontendUsersService *FrontendUsersService) Login(u *user.FrontendUsers, c *gin.Context) (userInter *user.FrontendUsers, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	clientIP := c.ClientIP()
	global.GVA_LOG.Info("Login clientIP", zap.String("ip", clientIP))
	var userInfo user.FrontendUsers
	if u.Username != "" {
		err = global.GVA_DB.Where("username = ?", u.Username).First(&userInfo).Error
	} else if u.CountryId != 0 && u.Phone != "" {
		err = global.GVA_DB.Where("country_id = ? and phone = ?", u.CountryId, u.Phone).First(&userInfo).Error
	} else {
		// email
		err = global.GVA_DB.Where("email = ?", u.Email).First(&userInfo).Error
	}

	if err == nil {
		if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	// update last_login_ip and last_login_time
	updates := map[string]interface{}{
		"last_login_ip":   clientIP,
		"last_login_time": time.Now().UnixMilli(),
	}
	global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", userInfo.ID).Updates(updates)
	return &userInfo, err
}

func (frontendUsersService *FrontendUsersService) GetTeamCount(info *userReq.TeamReq) (total int64, err error) {
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("parent_id = ? or grandparent_id = ? or great_grandparent_id = ?", info.TeamOwner, info.TeamOwner, info.TeamOwner).Count(&total).Error
	return
}
func (frontendUsersService *FrontendUsersService) GetSubUserList(info *userReq.SubUserReq) (list []userRes.SubUserResponse, total int64, err error) {
	list = make([]userRes.SubUserResponse, 0)

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&user.FrontendUsers{})
	var frontendUsers []user.FrontendUsers

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserID != 0 {
		db = db.Where("parent_id = (?)", info.UserID).Where("grandparent_id = ? or great_grandparent_id = ?", info.TeamOwner, info.TeamOwner)
	} else {
		db = db.Where("parent_id = (?)", info.TeamOwner)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Preload("UserFund").Find(&frontendUsers).Error

	userIDList := make([]uint, 0)

	for _, v := range frontendUsers {
		nickName := v.NickName
		if nickName == "" {
			nickName = "user_" + cast.ToString(v.ID)
		}
		userIDList = append(userIDList, v.ID)
		isCharge := 0
		if v.UserFund != nil && v.UserFund.FirstChargeTime != nil {
			isCharge = 1
		}
		list = append(list, userRes.SubUserResponse{
			ID:         v.ID,
			NickName:   nickName,
			IsRecharge: isCharge,
		})
	}

	userTeamCountMap := make(map[uint]int64)
	if len(userIDList) > 0 {
		userTeamCountMap, err = frontendUsersService.GetTeamUserCount(userIDList, info.TeamOwner)
	}

	for i, v := range list {
		list[i].SubUserCount = userTeamCountMap[v.ID]
	}
	return
}

func (frontendUsersService *FrontendUsersService) GetTeamUserCount(users []uint, teamOwner uint) (res map[uint]int64, err error) {
	res = make(map[uint]int64)
	for _, v := range users {
		count := int64(0)
		err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("parent_id = (?) or grandparent_id = ? or great_grandparent_id = ?", v, v, v).Where("parent_id = (?) or grandparent_id = ? or great_grandparent_id = ?", teamOwner, teamOwner, teamOwner).Count(&count).Error
		res[v] = count
	}
	return
}

func (frontendUsersService *FrontendUsersService) UpdateUserInfo(userID uint, req *userReq.UpdateUserInfoReq) (err error) {
	updates := make(map[string]interface{})
	if req.NickName != "" {
		updates["nick_name"] = req.NickName
	}
	if req.HeaderImg != "" {
		updates["header_img"] = req.HeaderImg
	}

	return global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", userID).Updates(updates).Error
}

func (frontendUsersService *FrontendUsersService) UpdateUserPassword(req *userReq.UpdateUserPassword) (err error) {
	updates := make(map[string]interface{})
	if req.Password != "" {
		updates["password"] = utils.BcryptHash(req.Password)
	}
	if req.PaymentPassword != "" {
		updates["payment_password"] = utils.BcryptHash(req.PaymentPassword)
	}

	return global.GVA_DB.Model(&user.FrontendUsers{}).Where("id = ?", req.UserID).Updates(updates).Error
}

func (frontendUsersService *FrontendUsersService) VerifyPaymentPassword(userID uint, password string) (ok bool, err error) {
	var userInfo user.FrontendUsers
	if err = global.GVA_DB.Where("id = ?", userID).First(&userInfo).Error; err != nil {
		return
	}
	ok = utils.BcryptCheck(password, userInfo.PaymentPassword)
	return
}

func (frontendUsersService *FrontendUsersService) GetUserIdByPhoneOrEmail(countryId uint, phone string, email string) (uid uint, err error) {
	var userInfo user.FrontendUsers
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("(country_id = ? and phone = ?) or email = ?", countryId, phone, email).First(&userInfo).Error
	if err != nil {
		return
	}
	uid = userInfo.ID
	return
}

func (frontendUsersService *FrontendUsersService) TeamCount(userID uint) (count int64, err error) {
	err = global.GVA_DB.Model(&user.FrontendUsers{}).Where("parent_id = ? or grandparent_id = ? or great_grandparent_id = ?", userID, userID, userID).Count(&count).Error
	return
}

func (frontendUsersService *FrontendUsersService) GetUsers(id string, phone string, email string, parentId string, ancestorsId string, userType uint) (userInfo []*user.FrontendUsers, err error) {
	db := global.GVA_DB.Model(&user.FrontendUsers{})
	if id != "" {
		db = db.Where("id = ?", id)
	}
	if phone != "" {
		db = db.Where("phone = ?", phone)
	}
	if email != "" {
		db = db.Where("email = ?", email)
	}
	if parentId != "" {
		db = db.Where("parent_id = ?", parentId)
	}
	if ancestorsId != "" {
		db = db.Where("root_userid in (?)", ancestorsId)
	}

	if userType > 0 {
		db = db.Where("user_type = ?", userType)
	}
	err = db.Find(&userInfo).Error
	if err != nil {
		return
	}
	return
}

func (frontendUsersService *FrontendUsersService) GetUsersByIds(ids []uint) (userInfo []*user.FrontendUsers, err error) {
	db := global.GVA_DB.Model(&user.FrontendUsers{}).Where("id in (?)", ids)
	err = db.Find(&userInfo).Error
	if err != nil {
		return
	}
	return
}

// CheckPhoneOrEmailExists 检查手机号或邮箱是否已存在
func (frontendUsersService *FrontendUsersService) CheckPhoneOrEmailExists(countryId uint, phone string, email string) (bool, error) {
	if phone != "" {
		var count int64
		err := global.GVA_DB.Model(&user.FrontendUsers{}).Where("country_id = ? AND phone = ?", countryId, phone).Count(&count).Error
		if err != nil {
			return false, err
		}
		if count > 0 {
			return true, nil
		}
	}

	if email != "" {
		var count int64
		err := global.GVA_DB.Model(&user.FrontendUsers{}).Where("email = ?", email).Count(&count).Error
		if err != nil {
			return false, err
		}
		if count > 0 {
			return true, nil
		}
	}

	return false, nil
}
