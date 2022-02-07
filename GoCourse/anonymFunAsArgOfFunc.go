package main

/*
Ejemplo de Funcion Anónima como argumento de otra funciòn.
*/

import (
	"fmt"
)

func SplitValues(f func(sum int) (int, int)) {
	x, y := f(10)
	fmt.Println(x, y)

	x, y = f(20)
	fmt.Println(x, y)
}

func main() {
	a, b := 2, 5

	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x

		return x, y
	}

	// Paso del valor de una funcion como argumento para otra funcion
	SplitValues(fn)

	// Obteniendo el valor de la funcion dado el valor del argumento
	x, y := fn(40)
	fmt.Println(x, y)

}
