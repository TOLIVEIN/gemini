package util

import (
	"fmt"
	"gemini/config"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
)

//GetImageFullURL ...
func GetImageFullURL(name string) string {
	return config.GetConfig().File.ImagePrefixURL + config.GetConfig().Port + config.GetConfig().PathSeparator + GetImagePath() + name
}

//GetImageName ...
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = EncodeMD5(fileName)
	return string(fileName) + ext
}

//GetImagePath ...
func GetImagePath() string {
	return config.GetConfig().File.ImageSavePath
}

//GetImageFullPath ...
func GetImageFullPath() string {
	// fmt.Println("image savepath: ", GetImagePath())
	return config.GetConfig().RootPath + GetImagePath()
}

//CheckImageExtension ...
func CheckImageExtension(fileName string) bool {
	ext := GetFileExtension(fileName)

	formats := strings.Split(config.GetConfig().File.ImageAllowExts, ",")
	// fmt.Println(formats)
	for _, allowExt := range formats {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

//CheckImageSize ...
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)

	if err != nil {
		// log.Println(err)
		// logging.Warn(err)
		fmt.Println(err)
		return false
	}
	imageSize, _ := strconv.Atoi(config.GetConfig().File.ImageMaxSize)
	imageSize = imageSize << 20
	// fmt.Println(size, imageSize)
	return size <= imageSize
}

//CheckImage ...
func CheckImage(source string) error {
	dir, err := os.Getwd()
	// fmt.Println(dir)
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = MakeDirectories(dir + source)
	// fmt.Println("source: ", source)
	if err != nil {
		return fmt.Errorf("MakeDirectories err: %v", err)
	}
	permission := CheckPermission(dir + source)
	// fmt.Println(permission)
	if permission == true {
		return fmt.Errorf("CheckPermission Permission denied source: %s", source)
	}
	return nil
}
