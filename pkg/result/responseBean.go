package result

import "CookingMaster_Backend/pkg/xerr"

type ResponseBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type NullJson struct{}

func Success(data interface{}) *ResponseBean {
	return &ResponseBean{
		Code: xerr.OK,
		Msg:  xerr.MapErrMsg(xerr.OK),
		Data: data,
	}
}

func Error(errCode uint32, errMsg string) *ResponseBean {
	return &ResponseBean{
		Code: errCode,
		Msg:  errMsg,
		Data: NullJson{},
	}
}
