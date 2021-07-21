package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var instance *JSSGorm
var once sync.Once

type JSSGorm struct {
	*gorm.DB
}

// Gorm
// @Description:
// @param dsn
// @param maxIdelConns
// @param maxOpenConns
// @return *JSSGorm
func Gorm(dsn, logZap string, maxIdelConns, maxOpenConns int) *JSSGorm {
	once.Do(func() {
		instance = &JSSGorm{
			initMySql(dsn, logZap, maxIdelConns, maxOpenConns),
		}
	})
	return instance
}

// GormMysql
// @Description: 初始化Mysql数据库
// @param dsn
// @param logZap
// @param idelCounts
// @param openCounts
// @return *gorm.DB
func initMySql(dsn, logZap string, idelCounts, openCounts int) *gorm.DB {
	mysqlConfig := mysql.Config {
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(logZap, false)); err != nil {
		fmt.Println("MySQL启动异常", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(idelCounts)
		sqlDB.SetMaxOpenConns(openCounts)
		return db
	}
}

// gormConfig
// @Description: 根据配置决定是否开启日志
// @param logZap
// @param mod
// @return *gorm.Config
func gormConfig(logZap string, mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch logZap {
	case "silent", "Silent":
		config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = Default.LogMode(logger.Info)
			break
		}
		config.Logger = Default.LogMode(logger.Silent)
	}
	return config
}
