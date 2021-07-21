package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

/*
	系统所有配置信息
*/

type Configure struct {
	JWT           JWT           `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System        System        `mapstructure:"system" json:"system" yaml:"system"`
	Zap           Zap           `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis         Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql         Mysql         `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local         Local         `mapstructure:"local" json:"local" yaml:"local"`
	Apollo        Apollo        `mapstructure:"apollo" json:"apollo" yaml:"apollo"`
	Tls			  Tls			`mapstructrue:"tls" json:"tls" yaml:"tls"`
	Remote        RemoteConfig
}

func InitConfigure(configFile string) *Configure {
	fmt.Println("正在读取本地初始配置......")
	config := new(Configure)

	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取异常," , err)
		os.Exit(0)
	}
	v.WatchConfig()
	v.Unmarshal(config)
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		v.Unmarshal(config)
	})
	return config
}
