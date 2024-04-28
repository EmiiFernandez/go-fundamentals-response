package response

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse representa una respuesta de error que incluye un mensaje y un código de estado.
type ErrorResponse struct {
	Status  int    `json:"status"`  // El código de estado HTTP asociado al error.
	Message string `json:"message"` // El mensaje de error.
}

// errorR es una función interna que crea una nueva instancia de ErrorResponse con el mensaje y el código de estado proporcionados.
func errorR(msg string, code int) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  code,
	}
}

// InternalServerError devuelve una respuesta de error con el mensaje proporcionado y el código de estado 500 (Internal Server Error) predeterminado.
func InternalServerError(msg string) Response {
	return errorR(msg, http.StatusInternalServerError)
}

// NotFound devuelve una respuesta de error con el mensaje proporcionado y el código de estado 404 (Not Found) predeterminado.
func NotFound(msg string) Response {
	return errorR(msg, http.StatusNotFound)
}

// Unauthorized devuelve una respuesta de error con el mensaje proporcionado y el código de estado 401 (Unauthorized) predeterminado.
func Unauthorized(msg string) Response {
	return errorR(msg, http.StatusUnauthorized)
}

// NonAuthoritativeForbiddentiveInfo devuelve una respuesta de error con el mensaje proporcionado y el código de estado 403 (Forbidden) predeterminado.
func NonAuthoritativeForbiddentiveInfo(msg string) Response {
	return errorR(msg, http.StatusForbidden)
}

// BadRequest devuelve una respuesta de error con el mensaje proporcionado y el código de estado 400 (Bad Request) predeterminado.
func BadRequest(msg string) Response {
	return errorR(msg, http.StatusBadRequest)
}

// Métodos de ErrorResponse

// Error devuelve una cadena vacía ya que las respuestas de error no tienen errores asociados.
func (e ErrorResponse) Error() string {
	return ""
}

// StatusCode devuelve el código de estado HTTP de la respuesta de error.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// GetBody serializa la respuesta de error a JSON y devuelve el cuerpo como un array de bytes.
func (e ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

// GetData devuelve nil ya que las respuestas de error no contienen datos.
func (e ErrorResponse) GetData() interface{} {
	return nil
}
