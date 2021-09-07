package mainpackage main

import (
"fmt"
)

func main() {
	// Exercício #01
	// Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.
	a := 9
	b := 7
	c := 3

	if a > b {
		if a > c {
			fmt.Println(a)
		} else {
			fmt.Println(c)
		}
	} else {
		if b > c {
			fmt.Println(b)
		} else {
			fmt.Println(c)
		}
	}

	// Exercício #02
	// Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.

	d := 2

	if d == 0 {
		fmt.Printf("O número é zero")
		return
	}
	if d%2 == 0 && d%3 == 0 {
		fmt.Printf("O número é múltiplo de 2 e de 3")
		return
	}
	if d%2 == 0 {
		fmt.Printf("O número é múltiplo de 2")
		return
	}
	if d%3 == 0 {
		fmt.Printf("O número é múltiplo de 3")
		return
	}
	fmt.Printf("O número não é igual a zero, múltiplo de 2 ou de 3")

}

