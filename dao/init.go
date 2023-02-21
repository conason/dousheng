package dao

import (
	"dousheng/dao/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

const dsn = "root:511518nibubda%@tcp(175.178.26.250:3307)/dousheng_db?charset=utf8mb4&parseTime=True&loc=Local"

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`love`表为`t_love`
			SingularTable: true,  // 使用单数表名，启用该选项后，`love` 表将是`love`
		},
	})
	if err != nil {
		return
	}
	db, err := DB.DB()
	if err != nil {
		return
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数
	db.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间
	db.SetConnMaxLifetime(10 * time.Second)

	dal.SetDefault(DB)
}
