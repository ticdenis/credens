package controller

import (
	"credens/apps/http/controller"
	"credens/libs/shared/application/serializer/json_iterator"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthzGetController(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.NewHealthzGetController(*json_iterator.NewJSONIteratorJSONSerializer()))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `"status":"OK"`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
