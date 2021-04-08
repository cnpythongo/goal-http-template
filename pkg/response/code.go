package response

import "os"

const (
	SuccessCode  = 1000
	FailCode     = 1001
	UnknownError = 1

	PayloadError = 1100

	AccountUserExistError      = 2000
	AccountEmailExistsError    = 2001
	AccountCreateError         = 2002
	AccountUserIdError         = 2003
	AccountUserNotExistError   = 2004
	AccountQueryUserError      = 2005
	AccountQueryUserParamError = 2006
	AccountQueryUserListError  = 2007
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
