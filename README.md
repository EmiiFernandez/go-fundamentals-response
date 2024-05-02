# Fundamentos de Go: Usuarios Web - Paquete de Respuestas

Este paquete proporciona un conjunto de interfaces y funciones para crear y manejar respuestas HTTP en aplicaciones Go, tanto las respuestas de éxito como los errores, simplificando el desarrollo de servicios web y APIs.

## Características Clave:

- **Interfaz de Respuesta:** Define una interfaz común para todos los tipos de respuestas, asegurando un comportamiento consistente y permitiendo un manejo genérico.

- **Respuesta de Éxito:** Representa una respuesta exitosa con un mensaje, código de estado y datos opcionales.

- **Respuesta de Error:** Representa una respuesta de error con un mensaje, código de estado y datos vacíos.

- **Funciones Predefinidas:** Proporciona funciones convenientes para crear tipos de respuesta comunes, incluyendo `OK`, `Created`, `InternalServerError`, `NotFound`, y más.

- **Serialización JSON:** Convierte automáticamente los objetos de respuesta en formato JSON para una transmisión fácil sobre HTTP.

- **Manejo de Errores:** Maneja posibles errores durante la serialización JSON, proporcionando un mecanismo de respuesta más robusto y fiable.

## Integración con go-fundamentals-web-users:

Este paquete de respuestas fue integrado en el proyecto `go-fundamentals-web-users` importándolo y utilizando sus funciones para crear y devolver respuestas.

## Ejemplo de Uso:

```go
import (
  "github.com/tu-nombre-de-usuario-en-github/go-fundamentals-web-users/response"
)

func makeGetAllEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}
		return response.OK("success", users), nil
	}
}
