package constants

var (
	UserRegisterCaptchaKEY = "user:register:captcha:{captchaId}"
	UserLoginCaptchaKEY    = "user:login:captcha:{captchaId}"
	UserInfoKEY            = "user:info:%s"
)

// REDIS KEY定义
const (
	RedisKeyPlatformCommissionRate  = "system:platform-commission-rate"  //平台佣金比例
	RedisKeyFirstGradeShareRate     = "system:first-grade-share-rate"    //一级分成比例
	RedisKeySecondGradeShareRate    = "system:second-grade-share-rate"   //二级分成比例
	RedisKeyThirdGradeShareRate     = "system:third-grade-share-rate"    //三级分成比例
	RedisKeyWithdrawCommissionType  = "system:withdraw-commission-type"  //手续费方式 1 固定 2按比例
	RedisKeyWithdrawCommissionQuota = "system:withdraw-commission-quota" //固定收费金额
	RedisKeyWithdrawCommissionRate  = "system:withdraw-commission-rate"  //手续费收费比例

	RedisKeySystemConfigInfo = "system:info" //系统业务相关配置信息

	RedisKeyIpToRegionList = "ip:to-region-list"
	RedisKeyIpRegionInfo   = "ip:region-info"
)
