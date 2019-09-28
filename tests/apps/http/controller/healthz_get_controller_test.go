package controller

import (
	"credens/apps/http/handler"
	"testing"
)

func TestHealthzGetController(t *testing.T) {
	sut := handler.NewHealthzGetHandler()

	_, err := sut.Handle(nil)

	if err != nil {
		t.Errorf("Ups!")
	}
}
