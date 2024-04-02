package response

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	resp := ResponseData{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	ctx.SetBody(resp.ToJson())
}

func Error(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	resp := ResponseData{
		Code: 1,
		Msg:  err.Error(),
		Data: nil,
	}
	ctx.SetBody(resp.ToJson())
}

func (resp *ResponseData) ToJson() []byte {
	jsonData, _ := json.Marshal(resp)
	return jsonData
}
