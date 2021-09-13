package mainpackage
package main

import (
"fmt"
)

func main() {
	// No Exercício #06 da seção "Exercícios", usamos for range para percorrer uma slice de string que representava uma lista de itens a comprar no mercado.
	// Agora, resolva o mesmo exercício usando a sintaxe básica da instrução for (sintaxe apresentada aqui).t
	var listaDeCompra = []string{"Arroz", "Feijão", "Batata", "Óleo", "Sal", "Café"}
	for i := 0; i < len(listaDeCompra); i++ {
		fmt.Printf("%d - %s\n", i+1, listaDeCompra[i])
	}
}