package Extra_Exercises

import (
	"fmt"
	"time"
)

package main

import (
"fmt"
"time"
)

func main() {
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

		var currentYear, birthYear, age int
		currentYear = time.Now().Year()
		fmt.Println("Em qual ano você nasceu?")
		fmt.Scan(&birthYear)
		age = currentYear - birthYear
		fmt.Printf("Sua idade é %d anos!", age)

	// 2) Como podemos pegar a informação do ano diretamente do console?
}

import (
	"fmt"
	"time"
)

package main

import (
"fmt"
"time"
)

func main() {
	var anoAtual int
	var anoNasci int
	var idade int
	anoAtual = time.Now().Year()
	fmt.Println("Quando você nasceu?")
	fmt.Scan(&anoNasci)
	idade = anoAtual - anoNasci

	fmt.Printf("Você tem %d anos", idade)
}