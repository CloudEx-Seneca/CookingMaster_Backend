package xerr

import "fmt"

type CodeError struct {
	errCode uint32
	errMsg  string
}

func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d, ErrMsg:%s", e.errCode, e.errMsg)
}

func NewCodeMsgError(errCode uint32, errmsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errmsg}
}

func NewCodeFullError(errCode uint32, err error) *CodeError {
	errFullMsg := fmt.Sprintf("%s:%s", MapErrMsg(errCode), err.Error())
	return &CodeError{errCode: errCode, errMsg: errFullMsg}
}

func NewCodeError(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
