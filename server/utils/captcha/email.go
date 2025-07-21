package captcha

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Email interface {
	Send(target, msg string) error
}

type aokSend struct {
}

func NewAokSend() Email {
	return &aokSend{}
}

func (s *aokSend) Send(target, code string) (err error) {
	// Set POST request data
	formData := url.Values{}
	formData.Set("app_key", "e87633e7aabfd335339dbf060af9f5bb") // Set the app key here
	formData.Set("to", target)                                  // Set the recipient here
	formData.Set("template_id", "E_100062763940")               // Set the template id here
	formData.Set("data", fmt.Sprintf(`{"code":"%s"}`, code))    // JSON data

	//global.GVA_LOG.Info("aokSend formData %v", formData.``)
	// Create a new POST request
	resp, err := http.PostForm("https://www.aoksend.com/index/api/send_email", formData)
	if err != nil {
		global.GVA_LOG.Info("aokSend Error:", zap.Any("err", err))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Info("aokSend Error ReadAll:", zap.Any("err", err))
		return
	}

	// Check if the request was successful
	global.GVA_LOG.Info("aokSend res StatusCode :", zap.Any("target", target), zap.Any("code", code), zap.Any("status", resp.StatusCode), zap.Any("Response", string(body)))
	return nil
}
