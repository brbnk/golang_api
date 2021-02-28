package response

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Status  int
	Message string
	Success bool
	Result  interface{}
}

func New() *HttpResponse {
	return &HttpResponse{
		Status:  http.StatusAccepted,
		Success: true,
	}
}

func (h *HttpResponse) SetStatus(s int) *HttpResponse {
	h.Status = s
	return h
}

func (h *HttpResponse) SetMessage(m string) *HttpResponse {
	h.Message = m
	return h
}

func (h *HttpResponse) SetSuccess(s bool) *HttpResponse {
	h.Success = s
	return h
}

func (h *HttpResponse) SetResult(r interface{}) *HttpResponse {
	h.Result = r
	return h
}

func (h *HttpResponse) Write(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(h.Status)
	response, _ := json.Marshal(h)
	w.Write(response)
}
