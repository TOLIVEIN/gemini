package database

import "fmt"

//CreateTag ...
func CreateTag(tag Tag) {
	result := db.Create(&tag)

	fmt.Println(result)
}

//ExistTagByName ...
func ExistTagByName(name string) bool {
	var tag Tag

	db.Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

//ExistTagsByIDs ...
func ExistTagsByIDs(ids []uint) bool {
	var tags []Tag

	// db.Where("id = ?", id).First(&tag)
	db.Find(&tags, ids)

	if len(tags) == len(ids) {
		return true
	}
	return false
}

//FindTagsByIDs ...
func FindTagsByIDs(ids []uint) (tags []*Tag) {

	// db.Where("id = ?", id).First(&tag)
	db.Find(&tags, ids)
	return
}

//FindTagByName ...
func FindTagByName(name string) (tag Tag) {

	db.Where("name = ?", name).First(&tag)
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

//EditTag ...
func EditTag(id uint, data map[string]interface{}) {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
}

//DeleteTag ...
func DeleteTag(id uint) {
	db.Where("id = ?", id).Delete(&Tag{})
}
