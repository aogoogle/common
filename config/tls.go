package config

/**
 * @Description https配置
 **/

type Tls struct {
	Addr string `mapstructure:"addr" yaml:"addr" json:"addr"`
	Cert string `mapstructure:"cert" yaml:"cert" json:"cert"`
	Key  string `mapstructure:"key" yaml:"key" json:"key"`
}
