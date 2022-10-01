package utils

import (
	"encoding/json"
	"net/http"
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

//APIResponse is for generating template responses
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

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(APIResponse("internal server error", http.StatusInternalServerError, "error", nil))
	if err != nil {
		InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonBytes)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(APIResponse("page not found", http.StatusNotFound, "error", nil))
	if err != nil {
		InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonBytes)
}

func APIResponseSuccess(w http.ResponseWriter, r *http.Request, message string, code int, status string, data interface{}) {
	jsonBytes, err := json.Marshal(APIResponse(message, http.StatusOK, status, data))
	if err != nil {
		InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func APIResponseError(w http.ResponseWriter, r *http.Request, message string, code int) {
	jsonBytes, err := json.Marshal(APIResponse(message, code, "error", nil))
	if err != nil {
		InternalServerError(w, r)
		return
	}
	w.WriteHeader(code)
	w.Write(jsonBytes)
}
