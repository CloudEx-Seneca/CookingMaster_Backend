package result

type ResponseBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type NullJson struct{}

func Success(data interface{}) *ResponseBean {
	return &ResponseBean{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

type ResponseErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{
		Code: errCode,
		Msg:  errMsg,
	}
}
