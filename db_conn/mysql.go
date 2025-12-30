package db_conn

import (
	"fmt"
	"graduation_project/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB
var mysqlOnce sync.Once

func InitMySQL(cfg config.MySQLConfig) (*gorm.DB, error) {
	var err error
	mysqlOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
			cfg.Charset,
			cfg.ParseTime,
		)

		mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}

		// 连接池配置
		sqlDB, _ := mysqlDB.DB()
		sqlDB.SetMaxIdleConns(10)  // 空闲连接数
		sqlDB.SetMaxOpenConns(100) // 最大打开连接数
	})
	return mysqlDB, err
}

func GetMySQL() *gorm.DB {
	return mysqlDB
}
