package mocks

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func MakeJSONRequestAndGetGinContext(method, target, body string) *gin.Context {
	request := httptest.NewRequest(
		method,
		target,
		bytes.NewBuffer([]byte(body)),
	)
	request.Header.Add("Content-Type", "application/json")

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = request

	return context
}
