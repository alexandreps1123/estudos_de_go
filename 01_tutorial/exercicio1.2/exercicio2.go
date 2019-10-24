// modificar o programa echo para exibir o indice e o valor de cada um de seus argumentos
package main

import (
	"fmt"
	"os"
)

func main() {
	//echo1
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(i, s)
	}

	s, sep = "", ""
	for i, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		println(i, s)
	}
}
