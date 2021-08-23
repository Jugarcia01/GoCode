// Programa que imprime la temperatura de Ebullicion del Agua, empleando funcion de otro archivo.
package main

import "fmt"
//import "github.com/Jugarcia01/GoCode/basicGoExamples/ch2-temperature/tempConvert"
//import "basicGoExamples/ch2-temperature/tempConvert"


func main() {
    //var tempEbullicAguaCent = tempConvert.F2c(ebullicF)
    const ebullicF = 212.0
    tempEbullicAguaCent := F2c(ebullicF)
    fmt.Printf("Punto de ebullición del Agua = %g°C\n",  tempEbullicAguaCent)

}

func F2c(fahrenheit float64) float64 {
    return (fahrenheit-32.0)*(5.0/9.0)
}

func C2f(celsius float64) float64 {
    return (celsius*(9.0/5.0))+32.0
}
