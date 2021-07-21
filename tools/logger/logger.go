package logger

import (
	"go.uber.org/zap"
	"sync"
)

//异常等级
const (
	Normal  = "普通异常"
	Serious = "严重异常"
	Fatal   = "致命异常"
	Information	= ""
)

type JSSLogger struct {
	*zap.Logger
}

var instance *JSSLogger
var once sync.Once

// Logger
// @Description: 单例logger
// @return *JSSLog
func Logger(prefix, logLevel string) *JSSLogger {
	once.Do(func(){
		instance = &JSSLogger{
			Zap(prefix, logLevel),
		}
	})
	return instance
}

// SeriousError
// @Description: 严重异常
// @param tag
// @param key
// @param detail
func (l *JSSLogger)SeriousError(key string, detail interface{}) {
	l.Error(Serious, zap.Any(key, detail))
}

// NormalError
// @Description: 普通异常
// @param key
// @param detail
func (l *JSSLogger)NormalError(key string, detail interface{}) {
	l.Error(Normal, zap.Any(key, detail))
}

// FatalError
// @Description: 致命异常
// @param key
// @param detail
func (l *JSSLogger)FatalError(key string, detail interface{}) {
	l.Error(Fatal, zap.Any(key, detail))
}

func (l *JSSLogger)NormalInfo(key string, detail interface{}) {
	l.Info(Information, zap.Any(key, detail))
}