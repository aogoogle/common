package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}
