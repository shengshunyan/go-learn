package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			//IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			//ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect database success", db)
	// 会执行建表语句
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to auto migrate database")
	}

	db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	fmt.Println(product)

	/** 用sql.NullString类型，让orm可以更新string为零值（空字符串） */
	db.Model(&product).Updates(Product{Code: sql.NullString{Valid: true, String: ""}})
}
