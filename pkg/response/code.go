package response

import "os"

const (
	SuccessCode  = 1000
	FailCode     = 1001
	UnknownError = 1
)

var MsgMapping = map[string]map[int]string{
	"en":    MessageEn,
	"zh_cn": MessageZHCN,
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
