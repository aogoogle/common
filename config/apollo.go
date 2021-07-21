package config

type Apollo struct {
	AppID string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Cluster string `mapstructure:"cluster" json:"cluster" yaml:"cluster"`
	Ip string `mapstructure:"ip" json:"ip" db:"ip"`
	NameSpace string `mapstructure:"namespace" json:"namespace" yaml:"namespace"`
	IsBackup bool `mapstructure:"isbackup" json:"isbackup" yaml:"isbackup"`
}
