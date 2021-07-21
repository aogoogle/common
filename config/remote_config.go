package config
/*
	远程配置管理
*/

import "sync"

type RemoteConfig struct {
	mutex sync.Mutex
	cfg   map[string]interface{}
}

func (r *RemoteConfig) Init() *RemoteConfig {
	r.cfg = make(map[string]interface{})
	return r
}

func (r *RemoteConfig) SetKey(key string, value interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.cfg == nil {
		r.Init()
	}
	r.cfg[key] = value
}

func (r *RemoteConfig) GetKey(key string) (value interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.cfg == nil {
		r.Init()
	}

	value, ok := r.cfg[key]
	if !ok {
		value = ""
	}
	return value
}

func (r *RemoteConfig) DefaultGetKey(key, defaultValue string) (value interface{}) {
	v, flag := r.Exists(key)
	if flag {
		return v
	}
	return defaultValue
}

func (r *RemoteConfig) Exists(key string) (interface{}, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	value, flag := r.cfg[key]
	return value, flag
}
