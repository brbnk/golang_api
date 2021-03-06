package response

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func New() *HttpResponse {
	return &HttpResponse{
		Status:  http.StatusOK,
		Success: true,
	}
}

// Set
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

// Writers
func (h *HttpResponse) InternalServerError(w http.ResponseWriter, r *http.Request) {
	h.SetStatus(http.StatusInternalServerError).
		SetSuccess(false).
		write(w, r)
}

func (h *HttpResponse) BadRequest(w http.ResponseWriter, r *http.Request) {
	h.SetStatus(http.StatusBadRequest).
		SetSuccess(false).
		write(w, r)
}

func (h *HttpResponse) NotFound(w http.ResponseWriter, r *http.Request) {
	h.SetStatus(http.StatusNotFound).
		write(w, r)
}

func (h *HttpResponse) NoContent(w http.ResponseWriter, r *http.Request) {
	h.SetStatus(http.StatusNoContent).
		write(w, r)
}

func (h *HttpResponse) Ok(w http.ResponseWriter, r *http.Request) {
	h.SetStatus(http.StatusOK).
		write(w, r)
}

func (h *HttpResponse) write(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(h.Status)
	response, _ := json.Marshal(h)
	w.Write(response)
}
