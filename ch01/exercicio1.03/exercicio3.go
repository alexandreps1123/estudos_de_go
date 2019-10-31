//Medir a diferenca de tempo de execucao entre as versoes echo e a versao que usa strings.Join
//-----ainda nao foi implementado a modificacao-----
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs\n", secs)
	fmt.Println(s)

	fmt.Printf("%.2fs\n", secs)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
