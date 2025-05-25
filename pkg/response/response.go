package response

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type APIResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data interface{}, pagination *Pagination) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIResponse{
		Data:       data,
		Pagination: pagination,
	}

	json.NewEncoder(w).Encode(resp)
}

func Error(w http.ResponseWriter, status int, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIResponse{
		Error: err,
	}

	json.NewEncoder(w).Encode(resp)
}
