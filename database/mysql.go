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

//CreateUser ...
func CreateUser(user User) {
	result := db.Create(&user)

	fmt.Println(result)
}

//CreateTag ...
func CreateTag(tag Tag) {
	result := db.Create(&tag)

	fmt.Println(result)
}

//ExistTag ...
func ExistTag(name string) bool {
	var tag Tag

	db.Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

//CreateArticle ...
func CreateArticle(article Article) {

	result := db.Create(&article)

	fmt.Println(result)
}

//FindUserByID ...
func FindUserByID(id uint) (user User) {

	db.Where("id = ?", id).First(&user)
	return
}

//FindTagByID ...
func FindTagByID(id uint) (tag Tag) {

	db.Where("id = ?", id).First(&tag)
	return
}

//FindArticleByID ...
func FindArticleByID(id uint) (article Article) {

	db.Where("id = ?", id).First(&article)
	return
}

//FindUserByName ...
func FindUserByName(name string) (user User) {

	db.Where("name = ?", name).First(&user)
	return
}

//FindTagByName ...
func FindTagByName(name string) (tag Tag) {

	db.Where("name = ?", name).First(&tag)
	return
}

//FindArticleByTitle ...
func FindArticleByTitle(title string) (article Article) {

	db.Where("title like ?", title).First(&article)
	return
}

//GetTags ...
func GetTags(page int, size int, conditions interface{}) (tags []Tag) {

	db.Where(conditions).Offset(page).Limit(size).Find(&tags)
	// fmt.Println(tags[0])
	return
}

//GetTagsCount ...
func GetTagsCount(conditions interface{}) (count int64) {

	db.Model(&Tag{}).Where(conditions).Count(&count)
	return
}
