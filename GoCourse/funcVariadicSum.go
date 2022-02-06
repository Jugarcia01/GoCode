package main

/*
Funciones Variadic
Una función que tiene un número variable de argumentos se conoce como función variádica.
O en otras palabras, un usuario se permite pasar cero o más argumentos en la función variadic.

    func funcName(vars ...type) type {
        code of function
    }

*/

import (
	"fmt"
)

// Funcion variadic Suma, en este caso recibe un numero variable de argumentos de tipo entero.
func Suma(nums ...int) int {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	valores := []int{1, 2, 3, 4, 5}
	total, x, y, z := 0, 4, 6, 8

	total = Suma(x, y, z)
	fmt.Printf("Suma de: %d, %d, %d es: %d\n", x, y, z, total)

	// Se pasa a la funcion Suma como argumento el slice de "valores"
	total = Suma(valores...)
	fmt.Printf("Suma de: %d es: %d\n", valores, total)

}
