/**
* @Author: Jackey
* @Date: 12/16/24 1:00 pm
 */

package config

type Tg struct {
	IsSend             bool     `mapstructure:"is-send" json:"is-send" yaml:"is-send"`                                        // 是否发送
	IgnoreRouter       []string `mapstructure:"ignore-router" json:"ignore-router" yaml:"ignore-router"`                      // 忽略的路径
	RouterSendDuration float64  `mapstructure:"router-send-duration" json:"router-send-duration" yaml:"router-send-duration"` // 控制的路由
	Host               string   `mapstructure:"host" json:"host" yaml:"host"`
	Path               string   `mapstructure:"path" json:"path" yaml:"path"`
	Token              string   `mapstructure:"token" json:"token" yaml:"token"`
	ChatId             string   `mapstructure:"chat-id" json:"chat-id" yaml:"chat-id"`
	NodeName           string   `mapstructure:"node-name" json:"node-name" yaml:"node-name"`
	NodeVersion        string   `mapstructure:"node-version" json:"node-version" yaml:"node-version"`
}
