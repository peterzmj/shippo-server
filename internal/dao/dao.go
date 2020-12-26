package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shippo-server/configs"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"time"
)

type Dao struct {
	db *gorm.DB
}

func New() (d *Dao) {

	var conf configs.DB
	utils.ReadConfigFromFile("configs/db.json", &conf)

	fmt.Printf("ReadConfigFromFile:%+v\n", conf)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.DSN, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	var users []model.User
	db.Find(&users)
	fmt.Printf("%+v\n", users)

	d = &Dao{
		db: db,
	}

	return d
}
