package database

import "gorm.io/gorm/clause"

//CreateArticle ...
func CreateArticle(article Article) {

	db.Create(&article)

	// fmt.Println(result)
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
	db.Preload(clause.Associations).Where(conditions).Offset(page).Limit(size).Find(&articles)

	return
}

//GetArticlesCount ...
func GetArticlesCount(conditions interface{}) (count int64) {

	db.Model(&Article{}).Where(conditions).Count(&count)
	return
}

//GetArticle ...
func GetArticle(id uint) (article Article) {
	db.Preload(clause.Associations).Where("id = ?", id).First(&article)
	return
}

//SearchArticles ...
func SearchArticles(page int, size int, title string) (articles []Article) {
	db.Preload(clause.Associations).Where("title like ?", title).Offset(page).Limit(size).Find(&articles)

	return
}

//SearchArticlesCount ...
func SearchArticlesCount(title string) (count int64) {
	db.Model(&Article{}).Where("title like ?", title).Count(&count)
	return
}

//EditArticle ...
func EditArticle(id uint, data interface{}) {
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
