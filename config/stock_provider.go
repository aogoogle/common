package config

type StockProvider struct {
	WadituUrl 		string `mapstructure:"waditu" json:"waditu" yaml:"waditu"`
	HengShengUrl	string `mapstructure:"hengsheng" json:"hengsheng" yaml:"hengsheng"`
	Provider 	string `mapstructure:"provider" json:"name" yaml:"name"`
}
