package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const dsn = "root:511518nibubda%@tcp(175.178.26.250:3307)/dousheng_db?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`love`表为`t_love`
			SingularTable: true,  // 使用单数表名，启用该选项后，`love` 表将是`love`
		},
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao/dal",
		ModelPkgPath: "./model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	// 设置目标 db
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
