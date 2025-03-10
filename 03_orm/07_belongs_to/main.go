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

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
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

	/** 创建表 */
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic("failed to auto create table")
	//}

	/** 新增一条数据 */
	//db.Create(&User{
	//	Name: "shane",
	//	Company: Company{
	//		ID:   1,
	//		Name: "min",
	//	},
	//})

	/** 关联查询 */
	var user User
	db.Preload("Company").First(&user)

	fmt.Println(user.Name, user.Company.Name)
}
