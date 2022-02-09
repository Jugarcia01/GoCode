// Programa que convierte texto en diferentes formatos de minisculas a MAYUSCULAS, etc.
package main

import (
	"fmt"
  "strings"
  "unicode"
)

func main() {
  fmt.Println("Convierte texto en diferentes formatos de minisculas a MAYUSCULAS y viceversa\n")

  var txt1 = "Este ES UN texto de prueba"
  fmt.Printf(txt1+"\n---- TEXTO PROCESADO -----\n")

  fmt.Printf(Mayus(txt1)+"\n")
  fmt.Printf(Minus(txt1)+"\n")
  fmt.Printf(CapitWords(txt1)+"\n")
  fmt.Printf(MayusIni(txt1)+"\n")

}

// Metodo que convierte en MAYÚSCULAS toda la cadena de texto
func Mayus(txt string) string {
  return strings.ToUpper(txt)
}

// Metodo que convierte en minúsculas toda la cadena de texto
func Minus(txt string) string {
  return strings.ToLower(txt)
}

// Metodo que convierte en Mayúsculas las iniciales de las palabras en la  cadena de texto
func CapitWords(txt string) string {
  return strings.Title(strings.ToLower(txt))
}

// Metodo que convierte en Mayúsculas solo la letra inicial de la  cadena de texto
func MayusIni(txt string) string {
  if len(txt) == 0 {
    return ""
  }
  temp := []rune(strings.ToLower(txt))
  temp[0] = unicode.ToUpper(temp[0])
  return string(temp)
}


