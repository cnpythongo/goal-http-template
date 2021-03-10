package common

import "os"

const (
	SuccessCode  = 0
	UnknownError = 10000
)

var EnglishMsg = map[int]string{
	SuccessCode:  "ok",
	UnknownError: "unknown error",
}

var ChinesMsg = map[int]string{
	SuccessCode:  "ok",
	UnknownError: "未知错误",
}

var MsgMapping = map[string]map[int]string{
	"en":    EnglishMsg,
	"zh_cn": ChinesMsg,
}

var Language = os.Getenv("LANGUAGE")

func GetCodeMsg(code int) string {
	mapping, ok := MsgMapping[Language]
	if !ok {
		return ""
	}
	msg, ok := mapping[code]
	if !ok {
		return ""
	}
	return msg
}
