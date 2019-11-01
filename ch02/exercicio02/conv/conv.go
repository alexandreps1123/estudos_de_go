//esse pacote esta responsavel por fazer diversas conversoes de medidas
package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Metros float64
type Pes float64
type Quilogramas float64
type Libras float64

const (
	pesConst   = 3.28084
	libraConst = 1 / 0.453592
)

//converte de Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//converte de Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//converte de Metros para Pes
func MToP(m Metros) Pes { return Pes(m * pesConst) }

//converte de Pes para Metros
func PToM(p Pes) Metros { return Metros(p / pesConst) }

//converte de Quilogramas para Libras
func QToL(q Quilogramas) Libras { return Libras(q / libraConst) }

//converte de Libras para Quilogramas
func LToQ(l Libras) Quilogramas { return Quilogramas(l * libraConst) }

//imprime o simbolo de Celsius ao lado de variaveis do tipo Celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

//imprime o simbolo de Fahrenheit ao lado de variaveis do tipo Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

//imprime o simbolo de Metros ao lado de variaveis do tipo metros
func (m Metros) String() string { return fmt.Sprintf("%gm", m) }

//imprime o simbolo de Pes ao lado de variaveis do tipo Pes
func (p Pes) String() string { return fmt.Sprintf("%gpes", p) }

//imprime o simbolo de Quilogramas ao lado de vaiaveisdo tipo Quilogramas
func (q Quilogramas) String() string { return fmt.Sprintf("%gkg", q) }

//imprime o simbolo de Libras ao lado de variaveis do tipo Libras
func (l Libras) String() string { return fmt.Sprintf("%glb", l) }
