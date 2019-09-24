package handler

import (
	"credens/libs/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Errors []error
}

func newResponseError(errs ...error) *ResponseError {
	return &ResponseError{Errors: errs}
}

type Response struct {
	Status  int
	Content map[string]interface{}
}

func newResponse(status int, content map[string]interface{}) *Response {
	return &Response{Status: status, Content: content}
}

type Handler interface {
	Handle(context *gin.Context) (*Response, *ResponseError)
}

func JSONHandler(handler Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		res, resErr := handler.Handle(context)
		if resErr != nil && len(resErr.Errors) > 0 {
			res = errorsToResponse(resErr.Errors)
		}

		context.JSON(res.Status, res.Content)
	}
}

func errorToResponse(err error) *Response {
	if customError, ok := err.(shared.Error); ok {
		status, err := strconv.Atoi(customError.Code())
		if err != nil || http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		}

		errFormatted := map[string]interface{}{
			"code":   customError.Code(),
			"detail": customError.Msg(),
			"title":  "Custom error",
		}

		meta := map[string]interface{}{}

		if customError.Data() != nil {
			meta["data"] = customError.Data()
			errFormatted["meta"] = meta
		}

		if customError.Err() != nil {
			meta["err"] = customError.Err().Error()
			errFormatted["meta"] = meta
		}

		return newResponse(
			status,
			map[string]interface{}{"errors": []map[string]interface{}{errFormatted}},
		)
	}

	return newResponse(
		http.StatusInternalServerError,
		map[string]interface{}{
			"errors": []map[string]interface{}{{
				"code": "0",
				"meta": map[string]interface{}{
					"err": err.Error(),
				},
				"title": "Unknown error",
			}},
		},
	)
}

func errorsToResponse(errs []error) *Response {
	if len(errs) == 1 {
		return errorToResponse(errs[0])
	}

	var content []map[string]interface{}

	for _, err := range errs {
		res := errorToResponse(err)
		content = append(content, res.Content["errors"].([]map[string]interface{})[0])
	}

	return newResponse(http.StatusBadRequest, map[string]interface{}{"errors": content})
}
