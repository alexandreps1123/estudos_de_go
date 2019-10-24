//Medir a diferenca de tempo de execucao entre as versoes echo e a versao que usa strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[0:], " "))
}
