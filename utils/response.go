package utils

import "github.com/mhmmdrivaldhi/go-social-media-api/model/dto"

type ResponseWithData struct {
	Code     int           `json:"code"`
	Status   string        `json:"status"`
	Message  string        `json:"message"`
	Paginate *dto.Paginate `json:"paginate,omitempty"`
	Data     interface{}   `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(params dto.ResponseParam) interface{} {
	var response interface{}
	var status string
	
	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "success"
	} else {
		status = "failed"
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Paginate: params.Paginate,
			Data:   params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return response
}