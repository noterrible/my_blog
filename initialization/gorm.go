package initialization

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"my_blog/global"
	"time"
)

func InitMysql() {
	if global.Config.Mysql.Host == "" {
		panic("Mysql配置错误")
	}
	dsn := global.Config.Mysql.DSN()
	var mysqlLogger logger.Interface
	if global.Config.Mysql.LogLevel == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else if global.Config.Mysql.LogLevel == "release" {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		panic("Mysql连接失败")
	}
	//预检
	var testDb int
	err = db.Debug().Raw("SELECT 1").Scan(&testDb).Error
	if err != nil {
		panic("预检失败")

	}
	fmt.Println("mysql连接成功")

	myDB, _ := db.DB()
	myDB.SetMaxIdleConns(10)               //设置最大空闲连接数
	myDB.SetMaxOpenConns(100)              //最多可打开连接数
	myDB.SetConnMaxLifetime(4 * time.Hour) //打开的连接的持续时长
	global.DB = db
}
