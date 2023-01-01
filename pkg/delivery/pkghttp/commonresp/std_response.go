package commonresp

import (
	"fmt"
	http2 "net/http"
)

type ResponseData struct {
	ResponseCode    string      `json:"responseCode,omitempty"`
	ResponseMessage string      `json:"responseMessage,omitempty"`
	Data            interface{} `json:"data,omitempty"`
	HTTPCode        int         `json:"-"`
}

func NewResponseData(responseCode Response, result interface{}) *ResponseData {

	return &ResponseData{
		ResponseCode:    fmt.Sprintf("%s", responseCode.GetCaseCode()),
		ResponseMessage: responseCode.Error(),
		Data:            result,
		HTTPCode:        getHttp(responseCode),
	}
}

func getHttp(code Response) int {
	http := http2.StatusInternalServerError

	if code != nil {
		http = code.GetHttp()
	}

	return http
}

func (e *ResponseData) StatusCode() int {
	return e.HTTPCode
}

func (e *ResponseData) Error() string {
	return e.ResponseCode
}

func (e *ResponseData) GetErrorResponse() error {
	return e
}
