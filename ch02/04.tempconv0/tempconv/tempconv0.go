package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsolutZeroC Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

//CToF converte uma temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converte uma temperatura em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//imprime °C em variaveis do tipo Celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

//imprime °F em variaveis do tipo Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
