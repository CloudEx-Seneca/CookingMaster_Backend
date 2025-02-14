package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "success"
	message[SERVER_COMMON_ERROR] = "server common error"
	message[REQUEST_PARAM_ERROR] = "request parameter error"
	message[TOKEN_INVALID_ERROR] = "token invalid"
	message[EMAIL_UNREGISTERED_ERROR] = "email unregistered"
	message[EMAIL_REGISTERD_BEFORE_ERROR] = "email register before"
	message[USER_NOT_VERIFIED_ERROR] = "user not verified"
	message[USER_PASSWORD_ERROR] = "password error"
}

func MapErrMsg(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "unknown error"
	}
}

func IsCodeErr(code uint32) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}
