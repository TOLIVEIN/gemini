package database

import (
	"fmt"
	"gemini/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//Init ...
func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().Mysql.Username,
		config.GetConfig().Mysql.Password,
		config.GetConfig().Mysql.URL,
		config.GetConfig().Mysql.Database)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: 0,
	// 		LogLevel:      logger.Info,
	// 		Colorful:      true,
	// 	},
	// )

	var err error
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Fail to connect to database ---- %s: %s\n",
			config.GetConfig().Mysql.Database, err)
	} else {
		fmt.Printf("Connected to database ---- %s\n",
			config.GetConfig().Mysql.Database)
		sqlDB, _ := db.DB()

		sqlDB.SetMaxIdleConns(10)

		sqlDB.SetMaxOpenConns(100)

		sqlDB.SetConnMaxLifetime(time.Hour)

	}

	CreateTables()

}

//CreateTables ...
func CreateTables() {
	if !db.Migrator().HasTable(&Article{}) {
		db.Migrator().CreateTable(&Article{})
		fmt.Println("Created table ---- articles")
	} else {
		fmt.Println("Exist table ---- articles")
	}
	if !db.Migrator().HasTable(&Tag{}) {
		db.Migrator().CreateTable(&Tag{})
		fmt.Println("Created table ---- tags")
	} else {
		fmt.Println("Exist table ---- tags")
	}
	if !db.Migrator().HasTable(&User{}) {
		db.Migrator().CreateTable(&User{})
		fmt.Println("Created table ---- users")
	} else {
		fmt.Println("Exist table ---- users")
	}
}
