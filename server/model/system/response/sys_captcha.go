package response

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
}
