// Programa que imprime la temperatura de Ebullicion del Agua, empleando funcion de otro archivo.
package main

import "fmt"
import "tempConvert"

const ebullicF = 212.0

func main() {
    var tempEbullicAguaCent = f2c(ebullicF)
    fmt.Printf("Punto de ebullición del Agua = %v°C\n",  tempEbullicAguaCent)

}

