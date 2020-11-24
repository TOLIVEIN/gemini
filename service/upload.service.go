package service

import (
	"fmt"
	"gemini/status"
	"gemini/util"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//UploadImage ...
func UploadImage(c *gin.Context) {
	code := status.Success
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Println(err)
		code = status.Error
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": status.GetMessage(code),
			"data":    data,
		})
	}
	if image == nil {
		code = status.InvalidParams
	} else {
		imageName := util.GetImageName(image.Filename)
		fullPath := util.GetImageFullPath()
		savePath := util.GetImagePath()

		// fmt.Println("imageName: ", imageName, "fullPath: ", fullPath, "savePath: ", fullPath)
		dir, _ := os.Getwd()
		src := dir + fullPath + imageName
		if !util.CheckImageExtension(imageName) || !util.CheckImageSize(file) {
			code = status.ErrorUploadCheckImageFormat
		} else {
			err := util.CheckImage(fullPath)
			if err != nil {
				// logging.Warn(err)
				fmt.Println(err)
				code = status.ErrorUploadCheckImageFail
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				// logging.Warn(err)
				fmt.Println(err)

				code = status.ErrorUploadSaveImageFail
			} else {
				data["image_url"] = strings.ReplaceAll(util.GetImageFullURL(imageName), "\\", "/")
				data["image_save_url"] = strings.ReplaceAll((savePath + imageName), "\\", "/")
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})
}
