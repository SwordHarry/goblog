package errcode

import (
	"fmt"
	"net/http"
)

// errcode 包
// 用于错误码管理组件

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var codes = map[int]string{}

// new error
func NewError(code int, msg string) *Error {
	// 排重校验
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在一个，请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	e.Details = []string{}
	for _, d := range details {
		e.Details = append(e.Details, d)
	}
	return e
}

// 错误码到状态码的转换
func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
