package database

import "fmt"

//CreateArticle ...
func CreateArticle(article Article) {

	result := db.Create(&article)

	fmt.Println(result)
}

//ExistArticleByName ...
func ExistArticleByName(name string) bool {
	var article Article

	db.Where("name = ?", name).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

//ExistArticleByID ...
func ExistArticleByID(id uint) bool {
	var article Article

	db.Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

//GetArticles ...
func GetArticles(page int, size int, conditions interface{}) (articles []Article) {
	db.Preload("Tag").Where(conditions).Offset(page).Limit(size).Find(&articles)

	return
}

//GetArticle ...
func GetArticle(id uint) (article Article) {
	db.Where("id = ?", id).First(&article)
	return
}

//EditArticle ...
func EditArticle(id uint, data map[string]interface{}) {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
}

//DeleteArticle ...
func DeleteArticle(id uint) {
	db.Where("id = ?", id).Delete(Article{})
}

//FindArticleByID ...
func FindArticleByID(id uint) (article Article) {

	db.Where("id = ?", id).First(&article)
	return
}

//FindArticleByTitle ...
func FindArticleByTitle(title string) (article Article) {

	db.Where("title like ?", title).First(&article)
	return
}
