/*
Escreva um programa de conversao de unidades de proposito geral, analogo ao cf, que leia numeros de seus argumentos
de linha de comando ou da entrada-padrao se nao houver argumentos, e converta cada numero em unidades como temperatura
em Celsius e em Fahrenheit, comprimento em pes e metros, peso em libras e quilogramas e operacoes semelhantes
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"../exercicio02/conv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := conv.Fahrenheit(t)
		c := conv.Celsius(t)
		m := conv.Metros(t)
		p := conv.Pes(t)
		l := conv.Libras(t)
		q := conv.Quilogramas(t)
		fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", m, conv.MToP(m), p, conv.PToM(p))
		fmt.Printf("%s = %s, %s = %s\n", l, conv.LToQ(l), q, conv.QToL(q))
	}
}
