package entity

import "net/http"

type BaseResult struct {
	Code int
	Msg  string
	Data any
}

func Success(data any) BaseResult {
	result := BaseResult{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: data,
	}
	return result
}

func Fail(msg string) *BaseResult {
	return &BaseResult{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: nil,
	}
}
