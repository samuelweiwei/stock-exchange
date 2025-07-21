package config

type EODHD struct {
	APIKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`
	BaseURL string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
}
