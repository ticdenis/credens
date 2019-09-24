package contracts

import (
	"credens/libs/shared/application/serializer"
	"net/http"
)

type JSONResponder struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	jsonSerializer serializer.JSONSerializer
}

func NewJSONResponder(responseWriter http.ResponseWriter, request *http.Request, jsonSerializer serializer.JSONSerializer) *JSONResponder {
	return &JSONResponder{responseWriter, request, jsonSerializer}
}

func (obj JSONResponder) DataResponse(statusCode int, dataType string, dataId interface{}, dataAttributes interface{}) {
	jsonResponseObject := NewJSONAPIDataObject(dataType, dataId, dataAttributes).ToJSONResponse()

	obj.response(statusCode, obj.serialize(jsonResponseObject))
}

func (obj JSONResponder) DataResponses(statusCode int, data []struct {
	dataType       string
	dataId         interface{}
	dataAttributes interface{}
}) {
	var dataObjects []JSONAPIDataObject
	for _, dataObject := range data {
		dataObjects = append(dataObjects, *NewJSONAPIDataObject(dataObject.dataType, dataObject.dataId, dataObject.dataAttributes))
	}

	jsonResponseObject := *NewJSONResponse(dataObjects, nil)

	obj.response(statusCode, obj.serialize(jsonResponseObject))
}

func (obj JSONResponder) ErrorsResponse(statusCode int, errors ...interface{}) {
	var errorsObjects []JSONAPIErrorObject
	for _, err := range errors {
		if errorObject, ok := err.(JSONAPIErrorObject); ok {
			errorsObjects = append(errorsObjects, errorObject)
		} else if stdError, ok := err.(error); ok {
			errorsObjects = append(errorsObjects, *NewJSONAPIErrorObject(stdError, statusCode))
		}
	}

	jsonResponseObject := *NewJSONResponse(nil, errorsObjects)

	obj.response(statusCode, obj.serialize(jsonResponseObject))
}

func (obj JSONResponder) serialize(input interface{}) []byte {
	data, _ := obj.jsonSerializer.Serialize(input)
	return data
}

func (obj JSONResponder) response(statusCode int, content []byte) {
	obj.responseWriter.Header().Set("Content-Type", "application/json")

	obj.responseWriter.WriteHeader(statusCode)

	_, _ = obj.responseWriter.Write(content)
}
