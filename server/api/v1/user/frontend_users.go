package user

import (
	"database/sql"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/client"
	"github.com/flipped-aurora/gin-vue-admin/server/constants"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/i18n"
	captchReq "github.com/flipped-aurora/gin-vue-admin/server/model/captcha"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userfund"
	captcha3 "github.com/flipped-aurora/gin-vue-admin/server/service/captcha"
	. "github.com/shopspring/decimal"

	//userReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"time"

	userRes "github.com/flipped-aurora/gin-vue-admin/server/model/user/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	captcha2 "github.com/flipped-aurora/gin-vue-admin/server/utils/captcha"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type FrontendUsersApi struct{}

var frontStore = captcha2.NewRedisCaptchaStore()

// CreateFrontendUsers 后台-创建frontendUsers表
// @Tags FrontendUsers
// @Summary 创建frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FrontendUsers true "创建frontendUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /users/createFrontendUsers [post]
func (frontendUsersApi *FrontendUsersApi) CreateFrontendUsers(c *gin.Context) {
	var frontendUsers user.FrontendUsers
	err := c.ShouldBindJSON(&frontendUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if frontendUsers.ParentId != 0 {
		parent, err := frontendUsersService.GetFrontendUsers(cast.ToString(frontendUsers.ParentId))
		if err != nil {
			response.FailWithMessage("上级用户不存在", c)
			return
		}
		if parent.RootUserid == 0 {
			frontendUsers.RootUserid = parent.ID
		} else {
			frontendUsers.RootUserid = parent.RootUserid
		}
	}

	err = frontendUsersService.CreateFrontendUsers(&frontendUsers)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	//默认创建用户资金表
	value := Zero
	now := time.Now()
	var myTime sql.NullTime
	myTime.Time = time.Time{}
	myTime.Valid = false // 设置为false表示时间为NULL
	userFundAccounts := &userfund.UserFundAccounts{
		UserId:            int(frontendUsers.ID),
		AssetType:         "USDT",
		Balance:           value,
		FrozenBalance:     value,
		AvailableBalance:  value,
		STATUS:            "1", //正常
		CreatedAt:         &now,
		UpdatedAt:         &now,
		DeletedAt:         nil,
		FirstChargeAmount: NullDecimal{},
		FirstChargeTime:   nil,
		UserType:          frontendUsers.UserType,
	}
	err = userfundAccountService.CreateUserFundAccounts(userFundAccounts)
	if err != nil {
		global.GVA_LOG.Error("资金账户创建失败!", zap.Error(err))
		response.FailWithMessage("资金账户创建失败", c)
		return
	}
	go func() {
		client.InitUserSymbolsFromAdmin(utils.GetToken(c), frontendUsers.ID)
	}()
	response.OkWithMessage("创建成功", c)
}

// DeleteFrontendUsers 后台-删除frontendUsers表
// @Tags FrontendUsers
// @Summary 删除frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FrontendUsers true "删除frontendUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /users/deleteFrontendUsers [delete]
func (frontendUsersApi *FrontendUsersApi) DeleteFrontendUsers(c *gin.Context) {
	ID := c.Query("ID")
	err := frontendUsersService.DeleteFrontendUsers(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFrontendUsersByIds 后台-批量删除frontendUsers表
// @Tags FrontendUsers
// @Summary 批量删除frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /users/deleteFrontendUsersByIds [delete]
func (frontendUsersApi *FrontendUsersApi) DeleteFrontendUsersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := frontendUsersService.DeleteFrontendUsersByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateFrontendUsers 后台-更新frontendUsers表
// @Tags FrontendUsers
// @Summary 更新frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FrontendUsers true "更新frontendUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /users/updateFrontendUsers [put]
func (frontendUsersApi *FrontendUsersApi) UpdateFrontendUsers(c *gin.Context) {
	var frontendUsers user.FrontendUsers
	err := c.ShouldBindJSON(&frontendUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = frontendUsersService.UpdateFrontendUsers(frontendUsers)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UpdateFailed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// FindFrontendUsers 后台-用id查询frontendUsers表
// @Tags FrontendUsers
// @Summary 用id查询frontendUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.FrontendUsers true "用id查询frontendUsers表"
// @Success 200 {object} response.Response{data=user.FrontendUsers,msg=string} "查询成功"
// @Router /users/findFrontendUsers [get]
func (frontendUsersApi *FrontendUsersApi) FindFrontendUsers(c *gin.Context) {
	ID := c.Query("ID")
	refrontendUsers, err := frontendUsersService.GetFrontendUsers(ID)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithData(refrontendUsers, c)
}

// GetFrontendUsers 前台-查询用户基本信息
// @Tags FrontendUsers
// @Summary 前台-查询用户基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.FrontendUsers true "用id查询frontendUsers表"
// @Success 200 {object} response.Response{data=user.FrontendUsers,msg=string} "查询成功"
// @Router /users/getFrontUserInfo [get]
func (frontendUsersApi *FrontendUsersApi) GetFrontendUsers(c *gin.Context) {
	claimUserID := utils.GetUserIDFrontUser(c)
	frontendUsers, err := frontendUsersService.GetFrontendUsers(cast.ToString(claimUserID))
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithData(frontendUsers, c)
}

// GetFrontendUsersList 后台-分页获取frontendUsers表列表
// @Tags FrontendUsers
// @Summary 后台-分页获取frontendUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.FrontendUsersSearch true "分页获取frontendUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /users/getFrontendUsersList [get]
func (frontendUsersApi *FrontendUsersApi) GetFrontendUsersList(c *gin.Context) {
	var pageInfo userReq.FrontendUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := frontendUsersService.GetFrontendUsersInfoList(pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.FailedObtained, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// UserLogin
// @Tags     FrontendUsers
// @Summary  前台-用户登录
// @Produce   application/json
// @Param    data  body      userReq.UserLogin                                             true  "邮箱 or 手机号, 密码"
// @Success  200   {object}  response.Response{data=userRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router /users/login [post]
func (frontendUsersApi *FrontendUsersApi) UserLogin(c *gin.Context) {
	var l userReq.UserLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证码 默认不需要，超过错误次数开启  // todo
	//if l.Captcha == "" || l.CaptchaId == "" || !frontStore.Verify(l.CaptchaId, l.Captcha, true) {
	//	response.FailWithMessage("验证码错误", c)
	//	return
	//}

	u := &user.FrontendUsers{
		Username:  l.UserName,
		CountryId: l.CountryId,
		Phone:     l.Phone,
		Email:     l.Email,
		Password:  l.Password,
	}
	userInfo, err := frontendUsersService.Login(u, c)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UsernameOrPasswordError, response.ERROR), c)
		return
	}
	if userInfo.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.UserBanLogin, response.ERROR), c)
		return
	}

	//根据ip获取地区
	var (
		clientIP  = c.ClientIP()
		region, _ = utils.GetRegionByIp(clientIP)
		UserAgent = c.Request.UserAgent()
	)
	if len(region) == 0 {
		_ = utils.AddIpStringInIpToRegionList(clientIP)
	}

	//写入登录记录
	err = frontendUserLoginLogService.CreateFrontendUserLoginLog(&user.FrontendUserLoginLog{
		Uid:         userInfo.ID,
		LoginIp:     clientIP,
		LoginRegion: region,
		LoginTime:   time.Now().UnixMilli(),
		UserAgent:   UserAgent,
	})
	if err != nil {
		utils.SendMsgToTgAsync("写入前台用户登录记录失败->" + err.Error())
	}

	frontendUsersApi.TokenNext(c, *userInfo)
	return
}

// UserRegister
// @Tags     FrontendUsers
// @Summary  前台-用户注册账号
// @Produce   application/json
// @Param    data  body      userReq.UserRegister                                            true  "用户名, 密码"
// @Success  200   {object}  response.Response{data=user.FrontendUsers, msg=string}  "用户注册账号,返回包括用户信息"
// @Router /users/register [post]
func (frontendUsersApi *FrontendUsersApi) UserRegister(c *gin.Context) {
	var r userReq.UserRegister
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if r.Phone == "" && r.Email == "" {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.PhoneAndEmailEmpty, response.ERROR), c)
		return
	}
	if r.Password == "" {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.PasswordEmpty, response.ERROR), c)
		return
	}
	if r.InviteCode == "" {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.InviteCodeEmpty, response.ERROR), c)
		return
	}

	// 检查手机号或邮箱是否已注册
	exists, err := frontendUsersService.CheckPhoneOrEmailExists(r.CountryId, r.Phone, r.Email)
	if err != nil {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	if exists {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.PhoneOrEmailRegisterd, response.ERROR), c)
		return
	}

	// 验证码
	if r.Captcha == "" || r.CaptchaId == "" || !frontStore.Verify(r.CaptchaId, r.Captcha, true) {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
		return
	}

	userName := frontendUsersApi.mergeUserName(r.CountryId, r.Phone, r.Email)
	if userName == "" {
		userName = r.Email
	}

	parentID := uint(0)
	grandParent := uint(0)
	greatGrandParent := uint(0)
	rootID := uint(0)
	if r.InviteCode != "" {
		userID, errInviteCode := utils.DecryptID(r.InviteCode)
		if errInviteCode != nil {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ParentUserNotExist, response.ERROR), c)
			return
		}

		parent, parentErr := frontendUsersService.GetFrontendUsers(cast.ToString(userID))
		if parentErr != nil {
			global.GVA_LOG.Error("Decrypt InviteCode err", zap.Error(parentErr))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ParentUserNotExist, response.ERROR), c)
			return
		}
		parentID = cast.ToUint(parent.ID)
		if parent.RootUserid == 0 {
			rootID = parent.ID
		} else {
			rootID = parent.RootUserid
		}

		grandParentInfo, grandParentErr := frontendUsersService.GetFrontendUsers(cast.ToString(parent.ParentId))
		if grandParentErr == nil && grandParentInfo.ID != 0 {
			grandParent = grandParentInfo.ID
			greatGrandParent = grandParentInfo.ParentId
		}
	}
	authenticationStatus := int(constants.AuthenticationStatusInvalid)
	userReg := &user.FrontendUsers{
		Username:             userName,
		Password:             r.Password,
		Enable:               1,
		CountryId:            r.CountryId,
		Phone:                r.Phone,
		Email:                r.Email,
		ParentId:             parentID,
		GrandparentId:        grandParent,
		GreatGrandparentId:   greatGrandParent,
		IdType:               *r.IdType,
		RealName:             r.RealName,
		IdNumber:             r.IdNumber,
		IdImages:             r.IdImages,
		AuthenticationStatus: &authenticationStatus,
		PaymentPassword:      r.PaymentPassword,
		RootUserid:           rootID,
		UserType:             1,
	}

	if r.IdType != nil && r.RealName != "" && r.IdNumber != "" && r.IdImages != "" {
		status := int(constants.AuthenticationStatusPending)
		userReg.AuthenticationStatus = &status
	}

	userReturn, err := frontendUsersService.Register(*userReg)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(userReturn, i18n.Message(request.GetLanguageTag(c), i18n.RegisterFail, response.ERROR), c)
		return
	}

	//默认创建用户资金表
	value := Zero
	now := time.Now()
	var myTime sql.NullTime
	myTime.Time = time.Time{}
	myTime.Valid = false // 设置为false表示时间为NULL
	userFundAccounts := &userfund.UserFundAccounts{
		UserId:            int(userReturn.ID),
		AssetType:         "USDT",
		Balance:           value,
		FrozenBalance:     value,
		AvailableBalance:  value,
		STATUS:            "1", //正常
		CreatedAt:         &now,
		UpdatedAt:         &now,
		DeletedAt:         nil,
		FirstChargeAmount: NullDecimal{},
		FirstChargeTime:   nil,
		UserType:          1,
	}
	err = userfundAccountService.CreateUserFundAccounts(userFundAccounts)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CreateFundAccountFail, response.ERROR), c)
		return
	}

	token, _, err := utils.LoginTokenFrontUser(&userReturn)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ErrGetTokenFail, response.ERROR), c)
		return
	}
	go client.InitUserSymbols(token)
	response.OkWithDetailed(userReturn, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

func (frontendUsersApi *FrontendUsersApi) mergeUserName(countryId uint, reqPhone, email string) string {
	if reqPhone == "" {
		return email
	} else {
		return cast.ToString(countryId) + "_" + reqPhone
	}
}

func (frontendUsersApi *FrontendUsersApi) verifyPhoneOrEmail(phone, email string) bool {
	if phone != "" {
		return true
	}

	if email != "" {
		return true
	}
	return false
}

// ChangePaymentPassword
// @Tags      FrontendUsers
// @Summary   前台-用户修改支付密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      userReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router /users/changePassword [post]
func (frontendUsersApi *FrontendUsersApi) ChangePaymentPassword(c *gin.Context) {
	var req userReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Password == req.NewPassword {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.NewPasswordSameAsOld, response.ERROR), c)
		return
	}
	if req.Password == "" {
		// 验证码
		if req.Captcha == "" || req.CaptchaId == "" || !frontStore.Verify(req.CaptchaId, req.Captcha, true) {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
			return
		}
	}
	uid := utils.GetUserID(c)
	_, err = frontendUsersService.ChangePaymentPassword(uid, req.Password, req.NewPassword)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.OperateSuccess, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// ChangeLoginPassword
// @Tags      FrontendUsers
// @Summary   前台-用户修改登录密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      userReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router /users/changePassword [post]
func (frontendUsersApi *FrontendUsersApi) ChangeLoginPassword(c *gin.Context) {
	var req userReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Password == req.NewPassword {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.NewPasswordSameAsOld, response.ERROR), c)
		return
	}
	if req.Password == "" {
		// 验证码
		if req.Captcha == "" || req.CaptchaId == "" || !frontStore.Verify(req.CaptchaId, req.Captcha, true) {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
			return
		}
	}
	uid := utils.GetUserID(c)
	_, err = frontendUsersService.ChangePassword(uid, req.Password, req.NewPassword)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// ResetLoginPassword
// @Tags      FrontendUsers
// @Summary   前台-用户重置登录密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      userReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router /users/resetLoginPassword [post]
func (frontendUsersApi *FrontendUsersApi) ResetLoginPassword(c *gin.Context) {
	var req userReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Password == req.NewPassword {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.NewPasswordSameAsOld, response.ERROR), c)
		return
	}
	if req.Password == "" {
		// 验证码
		if req.Captcha == "" || req.CaptchaId == "" || !frontStore.Verify(req.CaptchaId, req.Captcha, true) {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
			return
		}
	}
	uid := utils.GetUserIDFrontUser(c)
	if uid == 0 {
		uid, err = frontendUsersService.GetUserIdByPhoneOrEmail(req.CountryId, req.Phone, req.Email)
		if err != nil {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.NewPasswordSameAsOld, response.ERROR), c)
			return
		}
	}
	_, err = frontendUsersService.ChangePassword(uid, req.Password, req.NewPassword)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.ERROR), c)
}

// BindEmail
// @Tags     FrontendUsers
// @Summary  前台-前台用户绑定邮箱
// @Produce   application/json
// @Param    data  body      userReq.BindEmailReq true  "用户名, 密码"
// @Success  200   {object}  response.Response{msg=string}  "绑定成功"
// @Router /users/bindEmail [post]
func (frontendUsersApi *FrontendUsersApi) BindEmail(c *gin.Context) {
	var r userReq.BindEmailReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证码
	if r.Captcha == "" || r.CaptchaId == "" || !frontStore.Verify(r.CaptchaId, r.Captcha, true) {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
		return
	}
	userID := utils.GetUserIDFrontUser(c)
	err = frontendUsersService.BindEmail(userID, r.Email)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}

	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// BindPhone
// @Tags     FrontendUsers
// @Summary  前台-前台用户绑定手机
// @Produce   application/json
// @Param    data  body      userReq.BindPhoneReq true  "用户名, 密码"
// @Success  200   {object}  response.Response{msg=string}  "绑定成功"
// @Router /users/BindPhone [post]
func (frontendUsersApi *FrontendUsersApi) BindPhone(c *gin.Context) {
	var r userReq.BindPhoneReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证码
	if r.Captcha == "" || r.CaptchaId == "" || !frontStore.Verify(r.CaptchaId, r.Captcha, false) {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
		return
	}
	userID := utils.GetUserIDFrontUser(c)
	err = frontendUsersService.BindPhone(userID, &r)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// RealNameAuthentication
// @Tags     FrontendUsers
// @Summary  后台-用户实名认证审核
// @Produce   application/json
// @Param    data  body      userReq.RealNameAuthenticationReq                                         true  "用户名, 密码"
// @Success   200   {object}  response.Response{msg=string}  "审核成功"
// @Router /users/realNameAuthentication [post]
func (frontendUsersApi *FrontendUsersApi) RealNameAuthentication(c *gin.Context) {
	var r userReq.RealNameAuthenticationReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if constants.ParseRealNameAuthenticationStatus(r.Status) != constants.AuthenticationStatusPassed && constants.ParseRealNameAuthenticationStatus(r.Status) != constants.AuthenticationStatusRejected {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.CaptchaError, response.ERROR), c)
		return
	}

	err = frontendUsersService.RealNameAuthentication(r.ID, r.Status)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}

	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// ChangeParent
// @Tags     FrontendUsers
// @Summary  后台-修改用户上级代理
// @Produce   application/json
// @Param    data  body      userReq.ChangeParentReq                                          true  "用户id, 父级用户id"
// @Success   200   {object}  response.Response{msg=string}  "修改成功"
// @Router /users/changeParent [post]
func (frontendUsersApi *FrontendUsersApi) ChangeParent(c *gin.Context) {
	var r userReq.ChangeParentReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = frontendUsersService.ChangeParent(&r)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}

	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// TokenNext 登录以后签发jwt
func (frontendUsersApi *FrontendUsersApi) TokenNext(c *gin.Context, user user.FrontendUsers) {
	token, claims, err := utils.LoginTokenFrontUser(&user)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.ErrGetTokenFail, response.ERROR), c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetTokenFrontUser(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(userRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.StatusCodeError, response.ERROR), c)
			return
		}
		utils.SetTokenFrontUser(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(userRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, i18n.Message(request.GetLanguageTag(c), i18n.LoginSuccess, response.SUCCESS), c)
	} else if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.GetUsername()); err != nil {
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
			return
		}
		utils.SetTokenFrontUser(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(userRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, i18n.Message(request.GetLanguageTag(c), i18n.LoginSuccess, response.SUCCESS), c)
	}
}

// UserIdentity
// @Tags     FrontendUsers
// @Summary  前台-用户认证
// @Produce   application/json
// @Param    data  body      userReq.UserIdentityReq                                            true  "用户名, 密码"
// @Success  200   {object}  response.Response{data=user.FrontendUsers, msg=string}  "用户注册账号,返回包括用户信息"
// @Router /users/userIdentity [post]
func (frontendUsersApi *FrontendUsersApi) UserIdentity(c *gin.Context) {
	var r userReq.UserIdentityReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserIDFrontUser(c)
	err = frontendUsersService.UserIdentity(userID, &r)
	if err != nil {
		global.GVA_LOG.Error("认证失败!", zap.Error(err))
		response.FailWithDetailed(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), err.Error(), c)
		return
	}
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetSubUserList 前台-获取下级
// @Tags FrontendUsers
// @Summary 分页获取下级
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.SubUserReq true "分页获取下级"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /users/getSubUserList [get]
func (frontendUsersApi *FrontendUsersApi) GetSubUserList(c *gin.Context) {
	var pageInfo userReq.SubUserReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.TeamOwner == 0 {
		pageInfo.TeamOwner = utils.GetUserIDFrontUser(c)
	}

	list, total, err := frontendUsersService.GetSubUserList(&pageInfo)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetTeamCount 前台-获取团队总人数
// @Tags FrontendUsers
// @Summary 前台-获取团队总人数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.SubUserReq true "获取团队总人数"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /users/getTeamCount [get]
func (frontendUsersApi *FrontendUsersApi) GetTeamCount(c *gin.Context) {
	var info userReq.TeamReq
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	total, err := frontendUsersService.GetTeamCount(&info)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	response.OkWithDetailed(map[string]int64{"total": total}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// GetAncestors
// @Tags     FrontendUsers
// @Summary  前台-获取上三级代理
// @Produce   application/json
// @Param    data  body      userReq.GetAncestorsReq                                          true  "用户id"
// @Success  200   {object}  response.Response{data=[]uint}  "获取成功"
// @Router /users/getAncestors [get]
func (frontendUsersApi *FrontendUsersApi) GetAncestors(c *gin.Context) {
	var r userReq.GetAncestorsReq
	err := c.ShouldBindQuery(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if r.ID == 0 {
		r.ID = utils.GetUserIDFrontUser(c)
	}
	// 调用服务层的函数来获取祖先信息
	ancestors, err := frontendUsersService.GetAncestors(&r)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}

	// 返回成功响应
	response.OkWithDetailed(ancestors, i18n.Message(request.GetLanguageTag(c), i18n.SuccessfullyObtained, response.SUCCESS), c)
}

// UpdateUserInfo
// @Tags     FrontendUsers
// @Summary  前台-修改用户基本信息
// @Produce   application/json
// @Param    data  body      userReq.UpdateUserInfoReq                                          true  "用户头像、昵称"
// @Success  200   {object}  response.Response{data=[]uint}  "获取成功"
// @Router /users/updateUserInfo [post]
func (frontendUsersApi *FrontendUsersApi) UpdateUserInfo(c *gin.Context) {
	var r userReq.UpdateUserInfoReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if r.HeaderImg == "" && r.NickName == "" {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	userID := utils.GetUserIDFrontUser(c)
	err = frontendUsersService.UpdateUserInfo(userID, &r)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}

	// 返回成功响应
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// UpdateUserPassword
// @Tags     FrontendUsers
// @Summary  后台-修改用户密码
// @Produce   application/json
// @Param    data  body      userReq.UpdateUserPassword                                          true  "用户头像、昵称"
// @Success  200   {object}  response.Response{data=[]uint}  "修改成功"
// @Router /users/updateUserPassword [post]
func (frontendUsersApi *FrontendUsersApi) UpdateUserPassword(c *gin.Context) {
	var r userReq.UpdateUserPassword
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if r.Password == "" && r.PaymentPassword == "" {
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	err = frontendUsersService.UpdateUserPassword(&r)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}

	// 返回成功响应
	response.OkWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}

// TeamCount
// @Tags     FrontendUsers
// @Summary  前台-团队人数
// @Produce   application/json                                        true  ""
// @Success  200   {object}  response.Response{data=uint}  "成功"
// @Router /users/teamCount [get]
func (frontendUsersApi *FrontendUsersApi) TeamCount(c *gin.Context) {
	userID := utils.GetUserIDFrontUser(c)

	users, err := utils.GetClaimsFrontUser(c)

	// log user and err
	global.GVA_LOG.Info("user", zap.Any("user", users))

	count, err := frontendUsersService.TeamCount(userID)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.Failed, response.ERROR), c)
		return
	}
	// 返回成功响应
	response.OkWithData(map[string]int64{
		"count": count,
	}, c)
}

// SendCaptcha
// @Tags      Base
// @Summary   发送验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body systemReq.SendCaptcha true "发送验证码"
// @Success   200  {object}  response.Response{data=systemRes.CaptchaResponse,msg=string}  "发送验证码,返回包括随机数id"
// @Router    /base/sendCaptcha [post]
func (frontendUsersApi *FrontendUsersApi) SendCaptcha(c *gin.Context) {
	var req systemReq.SendCaptcha
	err := c.ShouldBindJSON(&req)

	// 字符,公式,验证码配置
	notificationChannel := captcha2.NewPhoneMessageChannel()
	target := req.Phone

	n := time.Now().UnixMilli()
	log := &captchReq.Captcha{
		ChannelType: uint(constants.ParseNotificationChannel(req.NotificationChannel)),
		CreatedAt:   n,
		UpdatedAt:   n,
		UserType:    1,
		AccessIp:    c.ClientIP(),
	}
	if req.CountryId != 0 {
		country, err := countriesService.GetCountries(cast.ToString(req.CountryId))
		if err != nil {
			global.GVA_LOG.Error("获取国家信息失败!", zap.Error(err))
			response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.GetCountryInfoFail, response.ERROR), c)
			return
		}
		log.NationalCode = country.PhoneCode
		log.UniqueId = target
		target = strings.Replace(country.PhoneCode, "+", "", -1) + target
	}
	if constants.ParseNotificationChannel(req.NotificationChannel) == constants.ChannelEmail {
		notificationChannel = captcha2.NewEmailMessageChannel()
		target = req.Email
		log.UniqueId = target
	}

	driver := captcha2.NewDriver(6)
	captchaStore := captcha2.NewRedisCaptchaStore()
	captchaIns := captcha2.NewCaptcha(notificationChannel, driver, captchaStore)
	id, captcha, err := captchaIns.Generate(target)
	if err != nil {
		_ = c.Error(err)
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.GetCaptchaFail, response.ERROR), c)
		return
	}
	log.CaptchaCode = captcha
	err = captcha3.NewLogService().Record(log)
	if err != nil {
		global.GVA_LOG.Error("log sending captcha err", zap.Error(err), zap.Any("req", req))
		response.FailWithMessage(i18n.Message(request.GetLanguageTag(c), i18n.GetCaptchaFail, response.ERROR), c)
		return
	}

	response.OkWithDetailed(systemRes.CaptchaResponse{
		CaptchaId: id,
		Captcha:   captcha,
	}, i18n.Message(request.GetLanguageTag(c), i18n.Success, response.SUCCESS), c)
}
