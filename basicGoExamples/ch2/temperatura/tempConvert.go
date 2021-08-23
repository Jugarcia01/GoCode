// Realiza la conversion numerica entre grados de temperatura, de Fahrenheit a Celsius o Celsius a Fahrenheit.
// e imprime el resultado en consola
package main
//package basicGoExamples


import "fmt"

func main(){
	const congelacF, ebullicF = 32.0, 212.0
	const congelacC, ebullicC = 0.0, 100.0
	fmt.Printf("%v°F = %v°C\n", congelacF, F2c(congelacF))
	fmt.Printf("%g°C = %g°F\n", congelacC, C2f(congelacC))

}

func F2c(fahrenheit float64) float64 {
	return (fahrenheit-32.0)*(5.0/9.0)
}

func C2f(celsius float64) float64 {
	return (celsius*(9.0/5.0))+32.0
}
