package main

import "fmt"

func main() {
	fmt.Println(fib(1), fib(2), fib(3), fib(4), fib(5), fib(6), fib(7))
}

func fib(n int) int {
	//atribuicao de tupla, permite que diversas variaveis recebam valores de uma so vez
	x, y := 0, 1
	for i := 1; i < n; i++ {
		y, x = y+x, y
	}
	return y
}
