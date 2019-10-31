package main

import (
	"fmt"

	"../exercicio01/tempconv"
)

func main() {
	fmt.Printf("AbsoluteZeroF: %v\n", tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Printf("FreezingF: %v\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("BoilingF: %v\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Println()
	fmt.Printf("AbsoluteZeroF: %v\n", tempconv.KToF(tempconv.AbsoluteZeroK))
	fmt.Printf("FreezingF: %v\n", tempconv.KToF(tempconv.FreezingK))
	fmt.Printf("BoilingF: %v\n", tempconv.KToF(tempconv.BoilingK))
}
