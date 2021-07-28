package main

import "fmt"

func main() {
	// ex 01
	// Utilizando a palavra reservada var declare uma variável numérica do tipo int.
	// Atribua um valor a essa variável.
	// Printe na tela o seu valor.
	// Repita para 3 variáveis diferentes.
	var a int = 1
	var b int = 2
	var c int = 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	// ex 02
	// Utilizando a marmota :=, declare duas variáveis: a e b.
	// Atribua os seguintes valores a elas, respectivamente: 230 e 27.
	// Crie variáveis para representar os resultados das operações matemáticas soma (a + b) e subtração (a - b).
	// Printe na tela os valores de todas as variáveis, um em cada linha.
	{
		a := 230
		b := 27
		sum := a + b
		sub := a - b
		fmt.Println(sum)
		fmt.Println(sub)
	}

	// ex 03
	// Declare variáveis para representar o preço dos itens do mercado conforme os valores da tabela abaixo:
	// Item
	// Preço
	// Banana
	// 1.25
	// Cerveja
	// 2.98
	// Abacate
	// 4.59
	// Salgadinho
	// 7.29
	// Declare variáveis para representar a quantidade ou peso dos itens do mercado conforme os valores da tabela abaixo. A banana e o abacate serão calculados pelo peso, a cerveja e o salgadinho serão calculados pela quantidade de unidades.
	//  Item
	// Quantidade/Peso
	// Banana
	// 2.170 kg
	// Cerveja
	// 6 unidades
	// Abacate
	// 5.650 kg
	// Salgadinho
	// 3 unidades
	// Printe na tela os valores totais de cada um dos itens e o preço total da compra.
	{
		banana := 1.25
		pesoBanana := 2.17
		cerveja := 2.98
		unCerveja := 6.00
		abacate := 4.59
		pesoAbacate := 5.65
		salgadinho :=7.29
		unSalgadinho := 3.00
		valorTotal := (banana*pesoBanana) + (cerveja*unCerveja) + (abacate*pesoAbacate) + (salgadinho*unSalgadinho)
		fmt.Printf("A banana custou R$ %v.", (banana*pesoBanana))
		fmt.Printf(" A cerveja custou R$ %v.", (cerveja*unCerveja))
		fmt.Printf(" O abacate custou R$ %v.", (abacate*pesoAbacate))
		fmt.Printf(" O salgadinho custou R$ %v.", (salgadinho*unSalgadinho))
		fmt.Printf(" O valor total da compra foi de R$ %v.", valorTotal)

	}

	// ex 04
	// Utilizando a palavra reservada var, declare uma variável do tipo string, e atribua a ela o valor que corresponde a seu nome usando uma literal de string interpretada.
	// Utilizando a marmota, declare uma variável do tipo string, e atribua a ela o valor que corresponde à sua cor favorita usando uma literal de string bruta.
	// Printe na tela os valores de todas as variáveis formando uma frase que faça sentido.
	{
		var nome string = "Juliana"
		cor := "azul"
		fmt.Printf("Oi! Meu nome é %v e minha cor preferida é %v! Até mais!", nome, cor)
	}

	// ex 05
	// Declare 5 variáveis e atribua operações relacionais a elas.
	// Usando uma linha por variável, printe na tela o valor e o tipo de cada uma das variáveis.
	{
		a := 10 == 2 // false
		b := -30 < 0 // true
		c := 42 != 3 // true
		d := 666 > 0 && -666 < 0 // true
		e := 5 >= 1 // true
		fmt.Printf("A variável A é do tipo %T e possui o valor %v\n", a, a)
		fmt.Printf("A variável B é do tipo %T e possui o valor %v\n", b, b)
		fmt.Printf("A variável C é do tipo %T e possui o valor %v\n", c, c)
		fmt.Printf("A variável D é do tipo %T e possui o valor %v\n", d, d)
		fmt.Printf("A variável E é do tipo %T e possui o valor %v\n", e, e)
	}

	// ex 06
	// Estabeleça uma comparação que utilize o operador lógico E (&&) duas vezes
	// Crie variáveis que representem condições para serem comparadas
	// Printe na tela o resultado da comparação (utilizando o operador lógico E) das variáveis - antes de rodar seu programa, tente prever qual será o resultado
	{
		fimDeSemana := true
		estaFrio := true
		estudosEmDia := true

		vouDormirMais := fimDeSemana && estaFrio && estudosEmDia

		fmt.Println(vouDormirMais) // true
	}

	// ex extra 01
	// A aplicação abaixo não compila.
	// package main
	//
	// import (
	// 	"fmt"
	// )
	//
	// func main() {
	// 	var quilometros int8
	// 	quilometros = 150
	//
	// 	fmt.Println(quilometros)
	//}
	// 1) Descubra por que não compila
	// O tipo int8 só engloba valores entre -128 e 127
	//
	// 2) Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
	// Ele indica que na linha 9, no caracter 14, o valor de uma variável está excedendo o valor máximo que ela pode armazenar
	//
	// 3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela
	{
		var quilometros int16
		quilometros = 150

		fmt.Println(quilometros)
	}

	// ex extra 02
	// 1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
	// 2) Como podemos pegar a informação do ano diretamente do console?
}
