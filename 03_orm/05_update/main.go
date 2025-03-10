package main

import (
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
	Code  string
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
	fmt.Println("connect database success")

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to auto migrate database")
	}

	db.Model(&Product{}).Where("id = ?", 1).Update("price", 2)

}
