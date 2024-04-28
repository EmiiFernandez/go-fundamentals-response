package response

// Response es una interfaz que define el contrato para todos los tipos de respuestas.
type Response interface {
	StatusCode() int          // Método para obtener el código de estado HTTP de la respuesta.
	GetBody() ([]byte, error) // Método para obtener el cuerpo de la respuesta como un array de bytes.
	Error() string            // Método para obtener el mensaje de error asociado a la respuesta, si lo hay.
	GetData() interface{}     // Método para obtener los datos asociados a la respuesta.
}
