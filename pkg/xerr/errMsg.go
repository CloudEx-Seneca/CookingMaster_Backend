package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "success"
	message[SERVER_COMMON_ERROR] = "server failed"
	message[REQUEST_PARAM_ERROR] = "request parameter error"
	message[TOKEN_EXPIRE_ERROR] = "token expired"
	message[TOKEN_GENERATE_ERROR] = "token generation failed"
	message[DB_ERROR] = "database failed"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "update affected zero"
	message[ONETIME_TOKEN_OVERUSED_ERROR] = "onetime token overused"
	message[TOKEN_REVOKED_ERROR] = "token revoked"
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
