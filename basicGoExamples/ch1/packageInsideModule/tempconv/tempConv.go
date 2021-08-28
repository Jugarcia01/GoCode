// Realiza la gestión de tipos de temperatura, de Fahrenheit, Celsius, Kelvin y define algunos valores de constantes
package tempconv

import "fmt"

type Celsius    float64
type Fahrenheit float64
type Kelvin     float64

const (
  AbsoluteZeroC Celsius = -273.15
  EbulliC       Celsius = 100.0
  CongelaC      Celsius = 0.0
)


func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

