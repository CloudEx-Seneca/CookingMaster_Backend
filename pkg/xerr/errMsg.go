package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "success"
	message[SERVER_COMMON_ERROR] = "Server common error"
	message[REQUEST_PARAM_ERROR] = "Request parameter error"
	message[TOKEN_INVALID_ERROR] = "Token invalid"
	message[EMAIL_UNREGISTERED_ERROR] = "User unregistered"
	message[EMAIL_REGISTERD_BEFORE_ERROR] = "This email has already been registered"
	message[USER_NOT_VERIFIED_ERROR] = "User not verified"
	message[USER_PASSWORD_ERROR] = "Email or password incorrect"
	message[DB_ERROR] = "Database error"
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
