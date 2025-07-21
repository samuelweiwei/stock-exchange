package config

type Server struct {
	JWT       JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	JWTUser   JWT     `mapstructure:"jwt-user" json:"jwt-user" yaml:"jwt-user"`
	Zap       Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	Mongo     Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Email     Email   `mapstructure:"email" json:"email" yaml:"email"`
	System    System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Tg        Tg      `mapstructure:"tg" json:"tg" yaml:"tg"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  Mssql           `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle Oracle          `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Sqlite Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local        Local        `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu        Qiniu        `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS    AliyunOSS    `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs    HuaWeiObs    `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS   TencentCOS   `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3        AwsS3        `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	CloudflareR2 CloudflareR2 `mapstructure:"cloudflare-r2" json:"cloudflare-r2" yaml:"cloudflare-r2"`

	// third pay
	ThirdPay ThirdPayConfig `mapstructure:"third-pay" json:"third-pay" yaml:"third-pay"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	EODHD EODHD `mapstructure:"eodhd" json:"eodhd" yaml:"eodhd"`

	// 添加 Polygon 配置
	Polygon struct {
		BaseURL string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
		APIKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`
	} `mapstructure:"polygon" json:"polygon" yaml:"polygon"`

	Service Service `mapstructure:"service" json:"service" yaml:"service"`
}

type Service struct {
	Symbol string `mapstructure:"symbol" json:"symbol" yaml:"symbol"`
}
