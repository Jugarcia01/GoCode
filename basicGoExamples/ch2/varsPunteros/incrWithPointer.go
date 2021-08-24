// Programa que incrementa el valor de una variable de forma indirecta a traves de una variable  puntero.

package main

import (
  "fmt"
)

func main() {
  v := 11
  fmt.Println("Variable v =", v)
  incr(&v)  // De forma indirecta el valor de v ahora pasa a ser 12
  fmt.Println("Despues de aplicar incr(&v), resulta v =", v)
}

func incr (p *int) int {
  *p++  // Incrementa el valor almacenado en la posicion a la cual apunta el puntero p. Ojo No se afecta la direcc de memoria..
  return *p
}

