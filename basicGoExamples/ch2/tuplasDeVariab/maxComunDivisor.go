// Programa que realiza el cálculo del Maximo Común Divisor entre dos nuneros enteros.
// https://es.m.wikipedia.org/wiki/M%C3%A1ximo_com%C3%BAn_divisor

//package tuplasDeVariab
package main

import(
  "fmt"
  "strconv"
)

func main() {
  fmt.Print("PROGRAMA PARA CALCULAR EL MAXIMO COMUN DIVISOR ENTRE DOS NUMEROS\nEs decir dados dos  num. x y;  Result = MCD(x,y)\n")

  isNum1 := false
		for !isNum1 {
      fmt.Println("\nIngrese el 1er  número entero: ")
			var dato1 string
			fmt.Scanf("%s\n", &dato1)
			x, err := strconv.Atoi(dato1)
			if err != nil {
				fmt.Println("ERROR: Debe ingresar un número entero válido.")
				isNum1 = false
      } else {
        isNum1 = true
        isNum2 := false
          for !isNum2 {
            fmt.Println("\nIngrese el 2do  número entero: ")
			      var dato2 string
			      fmt.Scanf("%s\n", &dato2)
			      y, err := strconv.Atoi(dato2)
			      if err != nil {
				      fmt.Println("ERROR: Debe ingresar un número entero válido.")
				      isNum2 = false
            } else if y == 0 {
				      fmt.Println("ERROR: Debe ingresar un número distinto de cero 0.")
            } else {
              result := mcd(x,y)

				      fmt.Printf("\nEl Max.ComunDivisor de %d y %d es %d.\n", x, y, result)
				      isNum2 = true
          }
			}
    }
  }
}

// Realiza el cálculo del máximo común divisor "MCD" entre dos números enteros dados x, y.
func mcd(x, y int) int {
  for y != 0 {
    x, y = y, x%y
  }
  return x
}
