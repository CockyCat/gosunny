package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DataBase(connString string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "eb_",
		},
		//Logger: logger.Default.LogMode(logger.Info),
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	db.Logger.LogMode(1)
	sqlDb, err := db.DB()
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	sqlDb.SetMaxIdleConns(20)
	//打开
	sqlDb.SetMaxOpenConns(100)
	//超时
	sqlDb.SetConnMaxLifetime(time.Second * 30)

	DB = db
}
