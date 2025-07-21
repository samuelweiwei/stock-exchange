package captcha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SMS interface {
	Send(target, msg string) error
}

type sms struct {
	AppKey    string
	AppSecret string
	AppCode   string
}

func NewSMS(appKey, appSecret, appCode string) SMS {
	return &sms{
		AppKey:    appKey,
		AppSecret: appSecret,
		AppCode:   appCode,
	}
}

type SmsRequest struct {
	AppKey    string
	AppSecret string
	AppCode   string
	Phone     string
	Msg       string
	UID       string // optional
	Extend    string // optional
}

// SendSMS sends a batch SMS request
func (s *sms) Send(target, msg string) error {
	// Prepare the request URL
	baseURL := "http://47.242.85.7:9090/sms/batch/v2"
	params := url.Values{}
	params.Add("appkey", s.AppKey)
	params.Add("appsecret", s.AppSecret)
	params.Add("appcode", s.AppCode)
	params.Add("phone", target)
	params.Add("msg", msg)

	// Construct the full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	fmt.Printf("SendSMS Send %s", fullURL)

	// Send the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and return the response body
	if resp.StatusCode != http.StatusOK {
		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// 打印错误信息
		fmt.Printf("Request failed with status: %s\n", resp.Status)
		fmt.Printf("Response body: %s\n", string(body))
		return fmt.Errorf("request failed with status:%s", resp.Status)
	}
	// 打印
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	fmt.Println(string(body))
	return nil
}
