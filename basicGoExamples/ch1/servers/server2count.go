// Server2count mimi servidor que emite una reepuesta "echo" y un contador

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
  log.Println("\nIniciado servidor que cuenta peticiones...\n"+"Ahora desde un browser intente peticiones con las siguientes URLs así:\n"+"http://localhost:8000\n"+"http://localhost:8000/count\n"+"Termine la ejecucion del server oprimiendo: Ctrl+c")
  
  http.HandleFunc("/", handler)
  http.HandleFunc("/count", counter)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// El manejador emite la ruta/Path del componente de la URL de peticion
func handler(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// El contador emite el num. de llamadas que se han realizado
func counter(w http.ResponseWriter, r * http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Contador = %d\n", count)
  mu.Unlock()
}

