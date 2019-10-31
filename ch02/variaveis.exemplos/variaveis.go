package main

import "fmt"

func main() {

	var i, j, k int                 //declarando multiplas variaveis do mesmo tipo
	var b, f, s = true, 2.3, "four" //declaracao de diversas variaveis de diferentes tipos
	a := "olá!"                     //declaracao curta de variavel
	i, j = 10, 20
	i, j = j, i //swap
	fmt.Println(i, j, k, b, f, s, a)

	//ponteiros
	x := 1
	p := &x           //p recebe o endereco de x
	*p++              //incrementa o valor de x
	fmt.Println(x, p) //imprime o valor de x e o endereco
	incr(&x)
	fmt.Println(incr(&x), *p, x)
	/*
		for count := 0; count <= 10; count++ {
			fmt.Println(count)
		}
	*/
	//funcao new(T): cria uma variavel sem nome do tipo T e retorna seu endereço
	c := new(int)
	fmt.Println(*c)
	*c = 2
	fmt.Println(*c)
	teste := newInt()
	fmt.Println(*teste, teste)

	//cada chamada de new devolve uma variavel distinta com endereco unico
	z := new(int)
	y := new(int)
	fmt.Println(z == y) //'false
}

func incr(p *int) int {
	*p++ //incrmenta o valor para o qual p aponta; nao altera p
	return *p
}

//cria uma variavel inteira e retorna seu endereco
func newInt() *int {
	return new(int)
}
