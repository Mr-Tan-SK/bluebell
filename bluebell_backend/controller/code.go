package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy
	CodeUnknownError

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExit:     "用户不存在",
	CodeInvalidPassword: "用户名或者密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeUnknownError:    "未知的错误",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
