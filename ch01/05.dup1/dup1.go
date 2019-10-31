//Dup1 exibe o texto de toda linha que aparece mais de
//uma vez na entrada-padrão precedida por sua contagem

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "" {
			break
		}
	}

	//NOTA: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}
