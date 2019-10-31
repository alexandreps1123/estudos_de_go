package main

import (
	"fmt"

	"../04.tempconv0/tempconv"
)

func main() {

	fmt.Printf("%g\n", tempconv.BoilingC) //"100" °C
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC)) //"180" °F
	//erro de compilacao, incompatibilidade de tipos
	//fmt.Printf("%g\n", boilingF-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0) //"true"
	fmt.Println(f >= 0) //"true"
	//erro de compilacao, incompatibilidade de tipos
	//fmt.Println(c == f)

	fmt.Println(c == tempconv.Celsius(f)) //"true"
	c = tempconv.FToC(212.0)
	fmt.Println(c.String()) //"100°C"
	fmt.Printf("%v\n", c)   //"100°C", nao tem necessidade de chamar String() explicitamente
	fmt.Printf("%s\n", c)   //"100°C"
	fmt.Println(c)          //"100°C"
	fmt.Printf("%g\n", c)   //"100", nao chama String()
	fmt.Println(float64(c)) //"100", nao chama String()
}
