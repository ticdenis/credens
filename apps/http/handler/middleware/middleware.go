package middleware

import (
	"credens/apps/http/handler"
	"github.com/gin-gonic/gin"
)

type HandlerMiddleware interface {
	Handle(handler handler.Handler) gin.HandlerFunc
}

type Middleware interface {
	Handle(handler gin.HandlerFunc) gin.HandlerFunc
}
