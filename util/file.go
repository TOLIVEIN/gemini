package util

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//GetSize ...
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

//GetFileExtension ...
func GetFileExtension(fileName string) string {
	return path.Ext(fileName)
}

//CheckExist ...
func CheckExist(source string) bool {
	_, err := os.Stat(source)
	// fmt.Println("source:", source, "exist info: ", fileInfo, "err: ", err)
	// fmt.Println(os.IsNotExist(err))
	return os.IsNotExist(err)
}

//CheckPermission ...
func CheckPermission(source string) bool {
	_, err := os.Stat(source)
	// fmt.Println("source:", source, "permission info:", fileInfo, "err: ", err)
	return os.IsPermission(err)
}

//MakeDirectories ...
func MakeDirectories(source string) error {
	if notExist := CheckExist(source); notExist == true {
		if err := os.MkdirAll(source, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	return nil
}

//OpenFile ...
func OpenFile(name string, flag int, permission os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, permission)
	if err != nil {
		return nil, err
	}
	return f, nil
}
