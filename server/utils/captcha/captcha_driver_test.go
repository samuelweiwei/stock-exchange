package captcha

import (
	"fmt"
	"testing"
)

func TestGenerateIdCaptcha(t *testing.T) {
	// 实例化DigitalCaptchaDrive
	driver := DigitalCaptchaDriver{length: 6} // 假设我们想要生成6位数字的验证码

	// 调用GenerateIdCaptcha方法
	id, captcha := driver.GenerateIdCaptcha("")
	fmt.Printf("id: %s, captcha: %s \n", id, captcha)
	// 测试id的长度是否正确
	if len(id) != idLen {
		t.Errorf("Expected id length to be %d, but got %d", idLen, len(id))
	}

	// 测试captcha的长度是否正确
	if len(captcha) != driver.length {
		t.Errorf("Expected captcha length to be %d, but got %d", driver.length, len(captcha))
	}

	// 测试captcha是否只包含数字
	for _, digit := range captcha {
		if digit < '0' || digit > '9' {
			t.Errorf("Captcha contains a non-digit character: %c", digit)
		}
	}
}
