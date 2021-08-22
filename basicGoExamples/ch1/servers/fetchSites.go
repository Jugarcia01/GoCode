// Este programa realiza la busqueda/fetch de URLs en paralelo y reporta su tamaño y tiempo de respuesta
// NOTA IMPORTANTE: Para iniciar este programa digite por ejemplo,
// go run fetchSites.go ./fetchSites https://golang.org http://gopl.io https://godoc.org

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	//"net/url"
)

func main() {
  fmt.Print("Ejecutando analisis de WebSites!\n") 
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetch(url, ch) // Inicia la gorutina
  }
  for range os.Args[1:] {
    fmt.Println(<-ch)  // Recibe data del canal ch
  }
  fmt.Printf("%.2fs transcurridos\n", time.Since(start).Seconds())
  
}

func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // Envia al canal ch
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close() // No filtra recursos
  if err != nil {
    ch <- fmt.Sprintf("espera mientra %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}


