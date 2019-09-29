package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthzGetHandler struct {
}

func NewHealthzGetHandler() *HealthzGetHandler {
	return &HealthzGetHandler{}
}

func (_ HealthzGetHandler) validate(context *gin.Context) *ResponseError {
	return nil
}

func (_ HealthzGetHandler) Handle(context *gin.Context) (*Response, *ResponseError) {
	return NewResponse(
		http.StatusOK,
		map[string]interface{}{
			"data": map[string]interface{}{
				"type": "service",
				"attributes": map[string]interface{}{
					"status": "OK",
				},
			},
		},
	), nil
}
