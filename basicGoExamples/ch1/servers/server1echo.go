// Server1echo es un miniServer web que retorna el path/ruta actual
// El prog. es solo un puñado de lineas, puesto que la mayoria del trabajo la realizan las funciones de las librerias empleadas. El prog. conecta un manejador para aquellas URLs entrantes hacia el puerto 8000 que comienzan con /
// Una peticion se representa como una estructura de tipo http.Requestla cual contiene un num. de campos relacionados, uno de los cuales es la URL de las peticiones entrantes. Cuando una petic. llega, se pasa a la func manejadora, la cual extrae la ruta del componente (/hello) de la URL de la petic. y la retorna a la respuesta, empleando fmt.Fprintf

// NOTA IMPORTANTE: Para iniciar este server en segundo plano en Linux/macOS ejecuta desde la terminal: 
//$ go run Server1echo.go &
// Posteriormente realiza peticiones desde ka linea de comandos así por ejem:
//$ go build gopl.io/ch1/fetch
//$ ./fetch http://localhost:8080
//URL.Path = "/"
//$ ./fetch http://localhost:8080/help
//URL.Path = "/help"

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    
    // Cada solicitud llama al manejador/handler
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// El manejador retorna la ruta/Path del componente de la URL en la peticion
func handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)
}

