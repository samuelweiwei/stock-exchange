package config

type ThirdPayConfig struct {
	NotifyUri      string `mapstructure:"notify-uri" json:"notify-uri" yaml:"notify-uri"`
	AAPayUri       string `mapstructure:"aa-pay-uri" json:"aa-pay-uri" yaml:"aa-pay-uri"`
	AAPayUserId    string `mapstructure:"aa-pay-user-id" json:"aa-pay-user-id" yaml:"aa-pay-user-id"`
	AAPaySecretKey string `mapstructure:"aa-pay-secret-key" json:"aa-pay-secret-key" yaml:"aa-pay-secret-key"`
}
