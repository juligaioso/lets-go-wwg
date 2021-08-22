package main

import (
	"fmt"
)

func main() {
	// ex 01
	// Escreva um programa que contenha um array com 7 elementos do tipo string.
	// O valor de cada elemento deve ser o número do índice escrito por extenso.
	// Printe na tela o tipo do seu array e os valores de seus elementos.

	indice := [7]string{"zero", "Um", "Dois", "Três", "Quatro", "Cinco", "Seis"}

	fmt.Println(indice)

	// ex 02
	// Crie 2 arrays e atribua o valor do segundo ao primeiro.

	var novoIndice [7]string

	novoIndice = indice

	fmt.Println(novoIndice)

	// ex 03
	// Declare duas slices de int e preencha 6 posições de cada uma.
	// Some as slices, formando uma 3ª slice.
	// Printe na tela o valor das três.

	var sliceUm = []int{1, 2, 3, 4, 5, 6}
	var sliceDois = []int{9, 8, 7, 6, 5, 4}
	sliceTres := append(sliceUm, sliceDois...)

	fmt.Println(sliceUm)
	fmt.Println(sliceDois)
	fmt.Println(sliceTres)

	// ex 04
	// Represente uma lista de compras usando uma slice de string.
	// Declare-a usando literal de slice.
	// Posteriormente, adicione mais itens à lista.
	// Printe todas as etapas.

	var listaDeCompras = []string{"Arroz", "Feijão", "Carne"}
	fmt.Println(listaDeCompras)

	listaDeCompras = append(listaDeCompras, "Tomate", "Alface")
	fmt.Println(listaDeCompras)

	// ex 05
	// Usando um literal, crie um mapa que tem como chaves nomes de cores e como valores seu código hexadecimal.
	// Delete uma entrada do mapa.
	// Printe todas as etapas.

	cores := map[string]string{"Vermelho": "#FF0000", "Roxo": "#A020F0", "Azul": "#0000FF", "Verde": "#00FF00"}
	fmt.Println(cores)

	delete(cores, "Verde")
	fmt.Println(cores)

	// ex 06
	// Crie um mapa chamado ano onde as chaves (keys) são os números dos meses do ano (ex: 1, 2) e os valores (values) são os nomes dos meses.
	// Printe na tela os nomes do meses 1 e 12.
	// Printe na tela o tamanho do mapa ano.

	meses := make(map[int]string)
	meses[1] = "Janeiro"
	meses[2] = "Fevereiro"
	meses[3] = "Março"
	meses[4] = "Abril"
	meses[5] = "Maio"
	meses[6] = "Junho"
	meses[7] = "Julho"
	meses[8] = "Agosto"
	meses[9] = "Setembro"
	meses[10] = "Outubro"
	meses[11] = "Novembro"
	meses[12] = "Dezembro"
	fmt.Printf("%v, %v\n", meses[1], meses[12])
	fmt.Println(len(meses))

	// Ex 7
	// Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
	// Crie três variáveis do tipo pessoa.
	// Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas

	type Pessoa struct {
		nome         string
		idade        int
		corPreferida string
	}

	func main() {
		pessoaUm := Pessoa{
			nome:         "Lena",
			idade:        55,
			corPreferida: "Rosa",
		}

		pessoaDois := Pessoa{
			nome:         "Juliana",
			idade:        30,
			corPreferida: "Azul",
		}

		pessoaTres := Pessoa{
			nome:         "Manuela",
			idade:        7,
			corPreferida: "Roxo",
		}

		fmt.Printf("%s tem %d anos de idade e sua cor preferida é %s\n", pessoaUm.nome, pessoaUm.idade, pessoaUm.corPreferida)
		fmt.Printf("%s tem %d anos de idade e sua cor preferida é %s\n", pessoaDois.nome, pessoaDois.idade, pessoaDois.corPreferida)
		fmt.Printf("%s tem %d anos de idade e sua cor preferida é %s\n", pessoaTres.nome, pessoaTres.idade, pessoaTres.corPreferida)
}
