package logger

import (
	"fmt"
	"github.com/aogoogle/common/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	level zapcore.Level
	prefix string
)

func Zap(tagPrefix string, logLevel string) (logger *zap.Logger) {
	fmt.Println("正在初始化日志......")
	if ok, _ := utils.PathExists("log"); !ok { // 判断是否有Director文件夹
		fmt.Printf("create log directory\n")
		_ = os.Mkdir("log", os.ModePerm)
	}
	prefix = tagPrefix
	switch logLevel { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}

	logger = logger.WithOptions(zap.AddCaller())

	return logger
}

/*
getEncoderConfig 获取zapcore.EncoderConfig
 */
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	//switch global.JSS_Config.Zap.EncodeLevel {
	//case "LowercaseLevelEncoder": // 小写编码器(默认)
	//	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	//case "LowercaseColorLevelEncoder": // 小写编码器带颜色
	//config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	//case "CapitalLevelEncoder": // 大写编码器
	//	config.EncodeLevel = zapcore.CapitalLevelEncoder
	//case "CapitalColorLevelEncoder": // 大写编码器带颜色
	//	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//default:
	//	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	//}
	return config
}

/*
getEncoder 获取zapcore.Encoder
 */
func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

/*
getEncoderCore 获取Encoder的zapcore.Core
*/
func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

/*
自定义日志输出时间格式
 */
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(prefix + "15:04:05.000"))
}
