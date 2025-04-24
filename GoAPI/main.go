// filepath: main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Message define la estructura para nuestra respuesta JSON
type Message struct {
    Text string `json:"message"`
}

// helloHandler maneja las solicitudes a la ruta /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") // Indicamos que la respuesta es JSON
    response := Message{Text: "¡Hola desde la API en Go!"}
    // Codificamos la respuesta a JSON y la enviamos
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        // Si hay un error al codificar, enviamos un error HTTP 500
        http.Error(w, "Error al generar la respuesta JSON", http.StatusInternalServerError)
        log.Printf("Error al codificar JSON: %v", err) // Logueamos el error en el servidor
    }
}

	// concatHandler maneja las solicitudes a la ruta /concat
	func concatHandler(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		cadena := r.URL.Query().Get("cadena")
		if cadena == "" {
			http.Error(w, "Se requiere el parámetro 'cadena'", http.StatusBadRequest)
			return
		}

		resultado := "Jelow " + cadena

		w.Header().Set("Content-Type", "application/json")
		response := Message{Text: resultado}
		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			http.Error(w, "Error al generar la respuesta JSON", http.StatusInternalServerError)
			log.Printf("Error al codificar JSON: %v", err)
			return
		}
	}


// rootHandler maneja las solicitudes a la ruta raíz /
func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Servidor API en Go funcionando. Prueba /hello") // Respuesta simple de texto
}

func main() {
    // Creamos un nuevo "mux" (manejador de rutas)
    // Usar http.NewServeMux() es generalmente mejor práctica que usar el DefaultServeMux (http.HandleFunc directamente)
    mux := http.NewServeMux()

    // Registramos nuestros handlers en el mux
    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/hello", helloHandler) // Asocia la ruta "/hello" con la función helloHandler
	mux.HandleFunc("/concat", concatHandler) // Asocia la ruta "/concat" con la función concatHandler

    // Definimos el puerto
    port := ":8080"
    fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

    // Iniciamos el servidor HTTP en el puerto especificado, usando nuestro mux
    // log.Fatal se usa para que si ListenAndServe devuelve un error (que no sea nil),
    // el programa termine y loguee el error.
    err := http.ListenAndServe(port, mux)
    if err != nil {
        log.Fatal("Error al iniciar el servidor: ", err)
    }
}