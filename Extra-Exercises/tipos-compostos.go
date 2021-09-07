package main

import (
	"fmt"
)

func main() {
	// Existem dois times de futebol, o time amarelo e o time vermelho. O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana)
	// e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).
	// Crie um array de string para cada time e nomeie com o nome do time.
	// Printe na tela os nomes dos jogadores de cada time

	amarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := []string{"Helena", "Jonas", "José", "Juliana"}

	fmt.Println("Amarelo:", amarelo)
	fmt.Println("Vermelho:", vermelho)

	// Considerando os times do item anterior, crie uma slice para representar cada um.
	// 1. Adicione o jogador Luis Inácio no time vermelho usando a função append()

	vermelho = append(vermelho, "Luis Inácio")

	// 2. Printe na tela os nomes dos jogadores do time vermelho

	fmt.Println("Vermelho:", vermelho)

	// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.
	// Passos:
	// Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde a chave seja o nome da cidade e o valor o nome do país.
	// Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país. Esse mapa segundo mapa terá tipo map[string]int, sendo a
	// chave o nome do país e o valor a quantidade de menções.
	// Printe na tela os valores do último mapa.

	mapa := map[string]string{
		"Berlim": "Alemanha",
		"Sófia": "Bulgária",
		"Osasco": "Brasil",
		"Valparaíso": "Chile",
		"Goiânia": "Brasil",
		"Varsovia": "Polônia",
		"Moscow": "Rússia",
	}

	frequenciaDeAparicoes := make(map[string]int)

	for _, valor := range mapa{
		frequenciaDeAparicoes[valor] += 1
	}

	fmt.Println(frequenciaDeAparicoes)
}