// Utilizando los paquetes importados tempConv y funConv
package main

import(
  "fmt"
  "tempConv"
  "funConv"
)


func main(){
	
	fmt.Printf("Esto esta muy frio!!! %v\n", tempConv.AbsoluteZeroC)

  tempFahEbullicH2O := funConv.C2f(tempConv.EbulliC)
	fmt.Printf("Temp. de ebullición del agua en Fahrenheit = %v\n", tempFahEbullicH2O)

}
