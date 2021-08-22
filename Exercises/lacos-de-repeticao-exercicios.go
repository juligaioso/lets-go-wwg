package main

import "fmt"

func main() {
	// ex 1
	// Utilizando a instrução for, printe na tela os números inteiros de 13 a 27 (inclusive).

	for i := 13; i < 28; i++ {
		fmt.Println(i)
	}

	// ex 2
	// Utilizando uma sintaxe de loop diferente da usada no exercício acima, printe na tela todas as horas do dia (usando o formato de 24 horas).

	hora := 00
	for hora < 24 {
		fmt.Printf("%.2d:00\n", hora)
		hora++
	}

	// ex 3
	// Utilizando a resolução anterior, demonstre também todos os minutos de um dia (formato 24 horas).

	hora := 00
	minuto := 00

	for hora < 24 {
		for minuto < 60 {
			fmt.Printf("%.2d:%.2d\n", hora, minuto)
			minuto++
		}
		hora++
		minuto = 00
	}

	// ex 4
	// É possível adicionar mais um loop à solução do exercício anterior? Vamos tentar: utilizando a instrução for, printe todos os segundos das 3 primeiras horas do dia (de 00:00:00 até 02:59:59).

	hora := 00
	minuto := 00

	for hora < 3 {
		for minuto < 60 {
			for segundo := 00; segundo < 60; segundo++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundo)
			}
			minuto++
		}
		hora++
		minuto = 00
	}

	// ex 5
	// Declare uma slice de string (entre 5 e 10 elementos) na qual o valor de cada elemento seja o número do seu índice, escrito por extenso. Usando for range, printe os valores de cada elemento e seu índice em uma nova linha.

	var slice = []string{"zero", "um", "dois", "tres", "quatro", "cinco"}
	for indice, valor := range slice {
		fmt.Printf("%d - %s\n", indice, valor)
	}

	// ex 6
	// Declare uma slice de string que representa uma lista de itens a comprar no mercado. Use o for range para percorrê-la e printe cada um dos itens de compra ao lado da sua ordem de inserção na lista.

	var listaDeCompra = []string{"Arroz", "Feijão", "Batata", "Óleo", "Sal", "Café"}
	for indice, valor := range listaDeCompra {
		fmt.Printf("%d - %s\n", indice+1, valor)
	}
}
