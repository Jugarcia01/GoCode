// Programa que realiza el calculo de la n'esima serie de Fibonacci mediante el uso iterativo de tuplas de dos variables x y y.
// https://es.m.wikipedia.org/wiki/Sucesi%C3%B3n_de_Fibonacci

package main

import ("fmt"
"strconv")

func main(){
  fmt.Print("PROGRAMA PARA CALCULAR LA SERIE DE FIBONACCI\nEs decir dado un num. N la serie Fibonacci es: 0,1,1,3,5,8,13,...")

  isNum := false
		for !isNum {
      fmt.Println("\nIngrese un número entero mayor de 0 para calcular su serie de Fibonacci: ")
			var dato string
			fmt.Scanf("%s\n", &dato)
			numIn, err := strconv.Atoi(dato)
			if err != nil {
				fmt.Println("ERROR: Debe ingresar un número entero válido.")
				isNum = false
			} else {
        fib := fibo(numIn)
				fmt.Printf("Fibonacci de %d es %d.\n", numIn, fib)
				isNum = true
			}
    }
}

func fibo(n int) int {
  x, y := 0, 1

  fmt.Printf("La serie Fibonacci es: \n")
  for i := 0; i <  n; {
    fmt.Printf("%d ", x)
    x, y = y, x+y
    i++
  }
	fmt.Printf("\n")

  return x
}
