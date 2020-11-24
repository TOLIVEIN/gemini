package database

//CreateUser ...
func CreateUser(user User) {
	db.Create(&user)
}

//FindUserByID ...
func FindUserByID(id uint) (user User) {

	db.Where("id = ?", id).First(&user)
	return
}

//FindUserByName ...
func FindUserByName(username string) (user User) {

	db.Where("username = ?", username).First(&user)
	return
}

//ExistUserByUsername ...
func ExistUserByUsername(username string) bool {

	var user User

	db.Where("username = ?", username).First(&user)

	if user.ID > 0 {
		return true
	}
	return false
}

//ExistUserByID ...
func ExistUserByID(id uint) bool {

	var user User

	db.Where("id = ?", id).First(&user)

	if user.ID > 0 {
		return true
	}
	return false
}

//GetUsers ...
func GetUsers(page int, size int, conditions interface{}) (users []User) {

	db.Where(conditions).Offset(page).Limit(size).Find(&users)
	return
}

//GetUser ...
func GetUser(id uint) (users []User) {

	db.Where("id = ?", id).Find(&users)
	return
}

//GetUsersCount ...
func GetUsersCount(conditions interface{}) (count int64) {

	db.Model(&User{}).Where(conditions).Count(&count)
	return
}

//EditUser ...
func EditUser(id uint, data interface{}) {
	db.Model(&User{}).Where("id = ?", id).Updates(data)
}

//DeleteUser ...
func DeleteUser(id uint) {
	db.Where("id = ?", id).Delete(User{})
}
