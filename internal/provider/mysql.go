package provider

import (
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// mysql 连接配置
func NewMysqlClient(conf *config.Config) *gorm.DB {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.Mysql.DSN(),
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), gormConfig)

	if err != nil {
		logger.Panicf("Mysql Connect Error: %v!", err)
		fmt.Println("Mysql Connect Error: ", err)
		return nil
	}

	if db.Error != nil {
		logger.Panicf("Database Error: %v!", err)
		fmt.Println("Database Error: ", err)
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
