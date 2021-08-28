// Funciones de conversión de temperatura

// F2c convierte un valor de temperatura Fahrenheit a Celsius
func F2c(f Fahrenheit) Celsius {
	return Celsius( (f - 32.0) * (5.0 / 9.0) )
}

//C2f convierte un valor de temperatura Celsius a Fahrenheit
func C2f(c Celsius) Fahrenheit {
	return Fahrenheit ( (c * (9.0 / 5.0)) + 32.0 )
}


