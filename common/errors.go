package common

import "os"

const (
	SuccessCode  = 1000
	FailCode     = 1001
	UnknownError = 1
)

var EnglishMsg = map[int]string{
	SuccessCode:  "ok",
	FailCode:     "failed",
	UnknownError: "unknown error",
}

var ChinesMsg = map[int]string{
	SuccessCode:  "ok",
	FailCode:     "操作失败",
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
