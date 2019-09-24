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

func (_ HealthzGetHandler) Handle(context *gin.Context) (*Response, *ResponseError) {
	return newResponse(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"type":       "service",
				"attributes": gin.H{"status": "OK"},
			},
		},
	), nil
}