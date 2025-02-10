package serctx

import (
	"fmt"
	"log"
	"nasspider/config"
	"nasspider/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initMysql() (*gorm.DB, error) {
	dbConfig := config.Conf.DB
	username := config.GetConf(dbConfig.Username, constants.ENV_MYSQL_USER)
	password := config.GetConf(dbConfig.Password, constants.ENV_MYSQL_PASSWORD)
	host := config.GetConf(dbConfig.Host, constants.ENV_MYSQL_HOST)
	port := config.GetConf(dbConfig.Port, constants.ENV_MYSQL_PORT)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, dbConfig.Database, dbConfig.Charset)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,                                // 禁用自动创建外键约束
		Logger:                                   logger.Default.LogMode(logger.Info), // 日志级别
	}); err != nil {
		log.Fatalf("mysql connect error %v", err)
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		return db, nil
	}
}
