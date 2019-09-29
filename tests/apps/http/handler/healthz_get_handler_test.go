package handler_test

import (
	"credens/apps/http/handler"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthzGetHandler(t *testing.T) {
	sut := handler.NewHealthzGetHandler()

	actualRes, actualErr := sut.Handle(nil)

	expectedRes := handler.NewResponse(200, map[string]interface{}{
		"data": map[string]interface{}{
			"type": "service",
			"attributes": map[string]interface{}{
				"status": "OK",
			},
		},
	})

	assert.Nil(t, actualErr)
	assert.Equal(t, expectedRes.Status, actualRes.Status)
	assert.Equal(t, fmt.Sprintf("%+v", expectedRes.Content), fmt.Sprintf("%+v", actualRes.Content))
}
