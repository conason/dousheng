package dao

import (
	"dousheng/dao/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

const dsn = "root:root@tcp(localhost:3306)/dousheng_db?charset=utf8mb4&parseTime=True&loc=Local"

//const dsn = "root:root@tcp(127.0.0.1:3306)/dousheng_db?charset=utf8mb4&parseTime=True&loc=Local"

func Init() {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`love`表为`t_love`
			SingularTable: true,  // 使用单数表名，启用该选项后，`love` 表将是`love`
		},
	})
	if err != nil {
		return
	}
	DB, err := db.DB()
	if err != nil {
		return
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数
	DB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量
	DB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间
	DB.SetConnMaxLifetime(10 * time.Second)

	dal.SetDefault(db)
}
