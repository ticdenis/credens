package contracts

import (
	"credens/libs/shared/application"
	"credens/libs/shared/domain"
	"credens/libs/shared/infrastructure"
	"fmt"
)

type JSONAPIDataObject struct {
	Type       string      `json:"type"`
	Id         string      `json:"id,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

func NewJSONAPIDataObject(dataType string, dataId interface{}, dataAttributes interface{}) *JSONAPIDataObject {
	idOmitted := dataId == nil || dataId == ""
	attributesOmitted := dataAttributes == nil || dataAttributes == ""

	if idOmitted {
		if attributesOmitted {
			return &JSONAPIDataObject{Type: dataType}
		}

		return &JSONAPIDataObject{Type: dataType, Attributes: dataAttributes}
	}

	id := dataId.(string)

	if attributesOmitted {
		return &JSONAPIDataObject{Type: dataType, Id: id}
	}

	return &JSONAPIDataObject{Type: dataType, Id: id, Attributes: dataAttributes}
}

func (dataObject JSONAPIDataObject) ToJSONResponse() JSONResponse {
	return JSONResponse{
		Data: dataObject,
	}
}

type JSONAPIErrorObject struct {
	HttpStatus int
	Code       string                 `json:"code,omitempty"`
	Detail     string                 `json:"detail,omitempty"`
	Meta       map[string]interface{} `json:"meta,omitempty"`
	Status     string                 `json:"status,omitempty"`
	Title      string                 `json:"title,omitempty"`
}

func NewJSONAPIErrorObject(err error, httpStatus int) *JSONAPIErrorObject {
	status := fmt.Sprintf("%d", httpStatus)

	if domainError, ok := err.(domain.DomainError); ok {
		return &JSONAPIErrorObject{
			HttpStatus: httpStatus,
			Code:   domainError.Code,
			Detail: domainError.Message,
			Meta: map[string]interface{}{
				"data": domainError.Data,
			},
			Status: status,
			Title:  "Domain error",
		}
	}

	if appError, ok := err.(application.ApplicationError); ok {
		return &JSONAPIErrorObject{
			HttpStatus: httpStatus,
			Code:   appError.Code,
			Detail: appError.Message,
			Meta: map[string]interface{}{
				"data": appError.Data,
			},
			Status: status,
			Title:  "Application error",
		}
	}

	if infraError, ok := err.(infrastructure.InfrastructureError); ok {
		return &JSONAPIErrorObject{
			HttpStatus: httpStatus,
			Code:   infraError.Code,
			Detail: infraError.Message,
			Meta: map[string]interface{}{
				"data": infraError.Data,
				"err":  infraError.Err.Error(),
			},
			Status: status,
			Title:  "Infrastructure error",
		}
	}

	return &JSONAPIErrorObject{
		HttpStatus: httpStatus,
		Detail: err.Error(),
		Meta: map[string]interface{}{
			"err": err.Error(),
		},
		Status: status,
		Title:  "UserInterface error",
	}
}

func (errorObject JSONAPIErrorObject) ToJSONResponse() JSONResponse {
	return JSONResponse{
		Errors: []JSONAPIErrorObject{errorObject},
	}
}

type JSONResponse struct {
	Data   interface{}          `json:"data,omitempty"`
	Errors []JSONAPIErrorObject `json:"errors,omitempty"`
}

func NewJSONResponse(data interface{}, errors []JSONAPIErrorObject) *JSONResponse {
	if data == nil && errors != nil {
		return &JSONResponse{Errors: errors}
	}

	if data != nil && errors == nil {
		return &JSONResponse{Data: data}
	}

	return &JSONResponse{}
}
