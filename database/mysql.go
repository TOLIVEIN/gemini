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
		config.CONFIG.Mysql.Username,
		config.CONFIG.Mysql.Password,
		config.CONFIG.Mysql.URL,
		config.CONFIG.Mysql.Database)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Fail to connect to database -- %s: %s\n",
			config.CONFIG.Mysql.Database, err)
	} else {
		fmt.Printf("Connected to database -- %s\n",
			config.CONFIG.Mysql.Database)
		sqlDB, _ := db.DB()

		sqlDB.SetMaxIdleConns(10)

		sqlDB.SetMaxOpenConns(100)

		sqlDB.SetConnMaxLifetime(time.Hour)

	}

}

// //GetDB ...
// func GetDB() *gorm.DB {
// 	return db
// }

//CreateUser ...
func CreateUser(user User) {
	result := db.Create(&user)

	fmt.Println(result)
}

//FindUserByID ...
func FindUserByID(id uint) (user User) {

	db.Where("id = ?", id).First(&user)
	return
}

//FindUserByName ...
func FindUserByName(name string) (user User) {

	db.Where("name = ?", name).First(&user)
	return
}
