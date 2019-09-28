package handler

import (
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Errors []error
}

func NewResponseError(errs ...error) *ResponseError {
	return &ResponseError{Errors: errs}
}

type Response struct {
	Status  int
	Content map[string]interface{}
}

func NewResponse(status int, content map[string]interface{}) *Response {
	return &Response{Status: status, Content: content}
}

type Handler interface {
	validate(context *gin.Context) *ResponseError
	Handle(context *gin.Context) (*Response, *ResponseError)
}
