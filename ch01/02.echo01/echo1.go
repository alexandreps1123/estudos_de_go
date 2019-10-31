/*
indexação em go utiliza intervalos abertos que incluim o primeiro indice e excluem o ultimo
exemplo: s[m:n] contem n-m elementos
*/
//echo1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	//por default string sao inicializadas por ""
	var s, sep string
	for i := 1; i < len(os.Args); i++ {

		//por default os.Args recebe como argumento os.Args[0:len(os.Args)]
		// o sinal de + com strings concatena os valores
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
