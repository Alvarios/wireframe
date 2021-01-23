package response

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

// Response is the default response format of the server after a succeded request
type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Body   string `json:"body"`
}

//NewResponse - create a new response instance
func NewResponse(ctx iris.Context, code int, body string) *Response {
	ctx.StatusCode(code)
	return &Response{
		Code: code,
		Status: http.StatusText(code),
		Body: body,
	}
}