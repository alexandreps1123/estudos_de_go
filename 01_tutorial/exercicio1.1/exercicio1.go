//Modificar o programa echo para exibir o nome do comando que o chamou
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//echo1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//echo2
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	//echo3
	fmt.Println(strings.Join(os.Args[1:], " "))

	//extra
	fmt.Println(os.Args[1:])
}
