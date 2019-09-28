package middleware

import (
	"credens/apps/http/handler"
	"credens/libs/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type JSONHandlerMiddleware struct {
}

func NewJSONHandlerMiddleware() *JSONHandlerMiddleware {
	return &JSONHandlerMiddleware{}
}

func (handlerMiddleware JSONHandlerMiddleware) Handle(handler handler.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		res, resErr := handler.Handle(context)
		if resErr != nil && len(resErr.Errors) > 0 {
			res = handlerMiddleware.errorsToResponse(resErr.Errors)
		}

		context.JSON(res.Status, res.Content)
	}
}

func (_ JSONHandlerMiddleware) errorToResponse(err error) *handler.Response {
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

		return handler.NewResponse(
			status,
			map[string]interface{}{"errors": []map[string]interface{}{errFormatted}},
		)
	}

	return handler.NewResponse(
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

func (handlerMiddleware JSONHandlerMiddleware) errorsToResponse(errs []error) *handler.Response {
	if len(errs) == 1 {
		return handlerMiddleware.errorToResponse(errs[0])
	}

	var content []map[string]interface{}

	for _, err := range errs {
		res := handlerMiddleware.errorToResponse(err)
		content = append(content, res.Content["errors"].([]map[string]interface{})[0])
	}

	return handler.NewResponse(http.StatusBadRequest, map[string]interface{}{"errors": content})
}
