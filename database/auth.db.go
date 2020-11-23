package database

import "gemini/status"

//CheckAuth ...
func CheckAuth(username, password string) int {
	var auth User
	db.Select("id").Where(User{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return status.Success
	}
	db.Select("id").Where(User{Username: username}).First(&auth)
	if auth.ID > 0 {
		return status.ErrorPassword
	}
	return status.ErrorNotExistUser
}
