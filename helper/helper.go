package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

type ResponseTable struct {
	Meta      MetaTable   `json:"meta"`
	DataTable interface{} `json:"datatable"`
	Data      interface{} `json:"data"`
}

type MetaTable struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Query   string `json:"query"`
}

func APIResponseTable(message string, code int, status string, query string, datatable interface{}, data interface{}) ResponseTable {
	meta := MetaTable{
		Message: message,
		Code:    code,
		Status:  status,
		Query:   query,
	}

	jsonResponse := ResponseTable{
		Meta:      meta,
		DataTable: datatable,
		Data:      data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
