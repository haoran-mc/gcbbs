package db

import (
	"fmt"
	"log"
	"time"

	"github.com/haoran-mc/gcbbs/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormDefaultLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// 初始化 MySQL 数据库
func init() {
	c := config.Conf.DB
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Name, c.Pass, c.Host, c.Port, c.DB)
	gormDB, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:                       dsn,   // DSN data source name
				DefaultStringSize:         255,   // string 类型字段的默认长度
				DisableDatetimePrecision:  true,  // 禁用 datatime 精度，MySQL 5.6 之前的数据库不支持
				DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库不支持
				DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库不支持
				SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
			},
		),
		&gorm.Config{
			Logger:                 gormDefaultLogger.Default.LogMode(gormDefaultLogger.Info),
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)           // 设置空闲的最大连接数
	db.SetMaxOpenConns(40)           // 设置与数据库的最大打开连接数
	db.SetConnMaxLifetime(time.Hour) // 设置可以重用连接的最长时间
	DB = gormDB
}
