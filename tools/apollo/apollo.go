package apollo

/**
 * 读取远程apollo配置服务器
 **/

import (
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
	"github.com/zouyx/agollo/v4/storage"
	"strings"
)

type ReadConfig func(key string, value interface{})

type CustomChangeEvent interface {
	OnChange(key string, value interface{})
}

type CustomChangeListener struct {
	Event CustomChangeEvent
}

func (c CustomChangeListener) OnChange(event *storage.ChangeEvent) {
	for key, value := range event.Changes {
		c.Event.OnChange(key, value.NewValue)
	}
}

func (c CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {

}

func Apollo(appId, cluster, ip, nameSpace string, isBackup bool,
	read ReadConfig, listener CustomChangeEvent) *agollo.Client {
	appCfg := &config.AppConfig{
		AppID:          appId,
		Cluster:        cluster,
		IP:             ip,
		NamespaceName:  nameSpace,
		IsBackupConfig: isBackup,
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return appCfg, nil
	})

	client.AddChangeListener(&CustomChangeListener{listener})
	writeConfig(appCfg.NamespaceName, read, client)
	return client
}

func writeConfig(namespace string, read ReadConfig, client *agollo.Client) {
	names := strings.Split(namespace, ",")
	for _, nameItem := range names {
		cache := client.GetConfigCache(nameItem)
		cache.Range(func(key, value interface{}) bool {
			read(key.(string), value)
			return true
		})
	}
}
