package request

type SendCaptcha struct {
	NotificationChannel string `json:"notificationChannel"` // 手机: phoneMessage 2. 邮箱:email
	CountryId           uint   `json:"countryId" swaggertype:"integer" example:"1"`
	Phone               string `json:"phone" example:"手机号码"`
	Email               string `json:"email" example:"电子邮箱"`
}
