package database

import (
	"fmt"
	"gemini/status"

	"golang.org/x/crypto/bcrypt"
)

//CheckAuth ...
func CheckAuth(username, password string) int {
	var auth User
	db.Select("id, password").Where("username = ?", username).First(&auth)

	fmt.Println(auth)
	if auth.ID > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
		if err == nil {
			return status.Success
		}
		fmt.Println(err)
		return status.ErrorPassword
	}
	return status.ErrorNotExistUser
}
