package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse representa una respuesta exitosa que incluye un mensaje, un código de estado y datos opcionales.
type SuccessResponse struct {
	Message string      `json:"message"` // El mensaje asociado a la respuesta.
	Status  int         `json:"status"`  // El código de estado HTTP de la respuesta.
	Data    interface{} `json:"data"`    // Los datos asociados a la respuesta.
}

// success es una función interna que crea una nueva instancia de SuccessResponse con el mensaje, los datos y el código de estado proporcionados.
func success(msg string, data interface{}, code int) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    data,
	}
}

// OK crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 200 (OK) predeterminado.
func OK(msg string, data interface{}) Response {
	return success(msg, data, http.StatusOK)
}

// Created crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 201 (Created) predeterminado.
func Created(msg string, data interface{}) Response {
	return success(msg, data, http.StatusCreated)
}

// Accepted crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 202 (Accepted) predeterminado.
func Accepted(msg string, data interface{}) Response {
	return success(msg, data, http.StatusAccepted)
}

// NonAuthoritativeInfo crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 203 (Non-Authoritative Information) predeterminado.
func NonAuthoritativeInfo(msg string, data interface{}) Response {
	return success(msg, data, http.StatusNonAuthoritativeInfo)
}

// NoContent crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 204 (No Content) predeterminado.
func NoContent(msg string, data interface{}) Response {
	return success(msg, data, http.StatusNoContent)
}

// ResetContent crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 205 (Reset Content) predeterminado.
func ResetContent(msg string, data interface{}) Response {
	return success(msg, data, http.StatusResetContent)
}

// PartialContent crea una respuesta de éxito con un mensaje y datos proporcionados, y un código de estado 206 (Partial Content) predeterminado.
func PartialContent(msg string, data interface{}) Response {
	return success(msg, data, http.StatusPartialContent)
}

// Métodos de SuccessResponse

// Error devuelve una cadena vacía ya que las respuestas exitosas no tienen errores asociados.
func (s *SuccessResponse) Error() string {
	return ""
}

// StatusCode devuelve el código de estado HTTP de la respuesta exitosa.
func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

// GetBody serializa la respuesta exitosa a JSON y devuelve el cuerpo como un array de bytes.
func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

// GetData devuelve los datos asociados a la respuesta exitosa.
func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
