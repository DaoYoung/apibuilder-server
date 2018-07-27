package rest

import "net/http"

var defaultEnvelope Envelope = &envelope{}

func SetDefaultEnvelope(envelope Envelope) {
	defaultEnvelope = envelope
}

type Envelope interface {
	Error(w ResponseWriter, code int, msg string, httpCode int) error
	WriteData(w ResponseWriter, data interface{}, meta map[string]interface{}, httpCode int) error
}

type envelope struct{}

type envelopeStructure struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

func (e *envelope) Error(w ResponseWriter, code int, error string, httpCode int) error {
	w.WriteHeader(httpCode)
	return w.WriteJson(envelopeStructure{
		Code:    code,
		Message: error,
	})
}

func (e *envelope) WriteData(w ResponseWriter, data interface{}, meta map[string]interface{}, httpCode int) error {
	w.WriteHeader(httpCode)
	return w.WriteJson(envelopeStructure{
		Data: data,
		Meta: meta,
	})
}

func Data(w ResponseWriter, data interface{}, meta map[string]interface{}, httpCode int) error {
	return defaultEnvelope.WriteData(w, data, meta, httpCode)
}

func Ok(w ResponseWriter, data interface{}, meta map[string]interface{}) error {
	return defaultEnvelope.WriteData(w, data, meta, http.StatusOK)
}

func Created(w ResponseWriter, data interface{}, meta map[string]interface{}) error {
	return defaultEnvelope.WriteData(w, data, meta, http.StatusCreated)
}

func NoContent(w ResponseWriter) error {
	return defaultEnvelope.WriteData(w, nil, nil, http.StatusNoContent)
}

func Error(w ResponseWriter, code int, message string, httpCode int) error {
	return defaultEnvelope.Error(w, code, message, httpCode)
}

func BadRequest(w ResponseWriter) error {
	return defaultEnvelope.Error(w, http.StatusBadRequest, "bad request", http.StatusBadRequest)
}

func Unauthorized(w ResponseWriter) error {
	return defaultEnvelope.Error(w, http.StatusUnauthorized, "unauthorized", http.StatusUnauthorized)
}

func NotFound(w ResponseWriter) error {
	return defaultEnvelope.Error(w, http.StatusNotFound, "resource not found", http.StatusNotFound)
}

func ServerError(w ResponseWriter) error {
	return defaultEnvelope.Error(w, http.StatusInternalServerError, "resource not found", http.StatusInternalServerError)
}

func ServerUnavailabe(w ResponseWriter) error {
	return defaultEnvelope.Error(w, http.StatusServiceUnavailable, "service unavailable", http.StatusServiceUnavailable)
}
