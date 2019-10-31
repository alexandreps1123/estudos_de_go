//faz conversao de temperaturas nas escalas Celsius, Fahrenheit e Kelvin
package tempconv

import "fmt"

/*
declaracao dos tipos Celsius, Fahrenheit e Kelvin
para que nao se possa fazer comparacoes e nem operacoes
geometricas entre eles
*/
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

//converte de Celsiu para Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

//converte de Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//converte de Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//converte de Fahrenheit para Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(CToK(FToC(f)))
}

//converte de Kelvin para Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

//converte de Kelvin para Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(CToF(KToC(k)))
}

//imprime °C em variaveis do tipo Celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

//imprime °F em variaveis do tipo Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
