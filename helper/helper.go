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

// type ResponseN struct {
// 	Meta MetaN `json:"meta"`
// 	Data DataN `json:"data"`
// }

// type MetaN struct {
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// 	Status  string `json:"status"`
// }

// type DataN struct {
// 	Kd_periode_spp int         `json:"kd_periode_spp"`
// 	Tahun          int         `json:"tahun"`
// 	Tahun_akademik string      `json:"tahun_akademik"`
// 	Detail         interface{} `json:"detail"`
// }

// func APIResponseN(message string, code int, status string, kd_periode_spp int, tahun int, tahun_akademik string, datax interface{}) ResponseN {
// 	meta := MetaN{
// 		Message: message,
// 		Code:    code,
// 		Status:  status,
// 	}

// 	data := DataN{
// 		Kd_periode_spp: kd_periode_spp,
// 		Tahun:          tahun,
// 		Tahun_akademik: tahun_akademik,
// 		Detail:         datax,
// 	}

// 	jsonResponse := ResponseN{
// 		Meta: meta,
// 		Data: data,
// 	}

// 	return jsonResponse
// }
