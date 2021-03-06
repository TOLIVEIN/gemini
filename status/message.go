package status

//Messages ...
var Messages = map[int]string{
	Success:                     "ok",
	Error:                       "fail",
	InvalidParams:               "请求参数错误",
	ErrorExistTag:               "已存在该标签名称",
	ErrorNotExistTag:            "该标签不存在",
	ErrorNotExistArticle:        "该文章不存在",
	ErrorExistUser:              "用户名已存在",
	ErrorNotExistUser:           "该用户不存在",
	ErrorPassword:               "密码错误",
	ErrorAuthCheckTokenFail:     "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:  "Token已超时",
	ErrorAuthToken:              "Token生成失败",
	ErrorAuth:                   "Token错误",
	ErrorUploadSaveImageFail:    "保存图片失败",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadCheckImageFormat: "图片格式或大小错误",
}

//GetMessage ...
func GetMessage(code int) string {
	message, status := Messages[code]
	if status {
		return message
	}
	return Messages[Error]
}
