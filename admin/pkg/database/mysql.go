package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"monica-admin/pkg/config"
	"time"
)

// NewGormDB 创建数据库连接
func NewGormDB() *gorm.DB {
	// 加载配置
	dbCfg := config.GlobalConfig().Db
	// 创建实例
	db, err := gorm.Open(mysql.Open(dbCfg.Dsn))
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnMaxLifetime) * time.Second)
	return db
}
