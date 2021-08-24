// Uso basico de Variables y Punteros
// En este programa se crea una variable int x, cuyo valor se modifica indirectamente por medio del puntero p.
// Al asignar el valor 2 al interior de la  posicion de memoria del puntero *p entonces el valor de x indirectamente pasa a ser 2. Lo anterior debido a que la variable x se encuentra almacenada en aquella direccion de memoria.
package main

import (
  "fmt"
)

func main () {
  x := 1          // var int x is created and asigned value 1
  p := &x         // p, of type *int, points to x
  fmt.Printf(msg(1)+"%v\n", p) // Location of p = 0x40000020123
  fmt.Printf(msg(2)+"%v\n", *p) // value of *p = "1"
  *p = 2          // equivalent to x = 2
  fmt.Printf(msg(3)+"%v\n"+msg(4), x)  // "2"
}

func msg (num int) string {
  switch num {
  case 1:
    return "Dirección del  puntero p = "
  case 2:
    return "Valor almacenado en x donde apunta puntero p = "
  case 3:
    return "Se asigna 2 al valor donde apunta el puntero *p entonces x = "
  case 4:
    return "Se observa que indirectamente x cambió.\n"
  default:
    return "opción de mensaje invalido!."
  }
}

