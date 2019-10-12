package handler

import "github.com/gin-gonic/gin"

type (
	ResponseError struct {
		Errors []error
	}

	Response struct {
		Status  int
		Content map[string]interface{}
	}

	Handler interface {
		validate(context *gin.Context) *ResponseError
		Handle(context *gin.Context) (*Response, *ResponseError)
	}
)

func NewResponseError(errs ...error) *ResponseError {
	return &ResponseError{Errors: errs}
}

func NewResponse(status int, content map[string]interface{}) *Response {
	return &Response{Status: status, Content: content}
}
