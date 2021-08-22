// Server1echo es un miniServer web que retorna el path/ruta actual

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    
    // Cada solicitud llama al manejador/handler
    http.HandlerFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// El manejador retorna la ruta/Path del componente de la URL en la peticion
func handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)

}

//output



