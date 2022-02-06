package main

/*
Programa para el manejo de funciones en Go.
Una funcion se crea con la palabra clave func y se le asigna un nombre.
Se puede asignar un tipo de retorno, una lista de parámetros y en el
cuerpo de la funcion se escribe la logica y lista de instrucciones de la funcion.

Estructura de una funcion:

	func nombreFuncion(parametrosEntrada) tipoRetornoSalida {
		cuerpoFuncion
	}

*/

import (
	"fmt"
	"strconv"
	"strings"
)

// Funcion suma de dos enteros
func suma(x, y int) int {
	return x + y
}

// Funcion resta de dos enteros
func resta(x, y int) int {
	return x - y
}

// Funcion que organiza los datos y el resultado en forma de texto
func textoResultado(operac string, nums ...int) string {
	texto := ""

	switch strings.ToLower(operac) {
	case "suma":
		texto = operac + ": " + strconv.Itoa(nums[0]) + " + " + strconv.Itoa(nums[1]) + " = " + strconv.Itoa(suma(nums[0], nums[1])) + "\n"
	case "resta":
		texto = operac + ": " + strconv.Itoa(nums[0]) + " - " + strconv.Itoa(nums[1]) + " = " + strconv.Itoa(resta(nums[0], nums[1])) + "\n"
	}

	return texto
}

func main() {

	x := 5
	y := 3
	msg := ""

	// Se ejecuta la funcion suma,
	//msg = textoResultado("SUMA", x, y)
	//fmt.Printf(msg)
	fmt.Printf("SUMA: %d + %d = %d\n", x, y, suma(x, y))

	// Se ejecuta la funcion resta
	//fmt.Printf("RESTA: %d - %d = %d\n", x, y, resta(x, y))
	msg = textoResultado("RESTA", x, y)
	fmt.Printf(msg)

}
