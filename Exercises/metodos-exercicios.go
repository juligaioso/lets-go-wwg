package main

import (
	"fmt"
	"math"
)

func main() {
	// ex 1
	// Construa dois métodos: um que retorna a área e outro que retorna o perímetro de uma estrutura que representa um círculo. Printe a área e o perímetro de um círculo.
	circulo := Circulo{
		raio: 3,
	}

	areaCirculo := circulo.AreaCirculo()
	perimetroCirculo := circulo.PerimetroCirculo()

	fmt.Printf("círculo:\n\traio: %.2f\n\tárea: %.2f\n\tperímetro: %.2f", circulo.raio, areaCirculo, perimetroCirculo)
	// ex 2
	// Considere um cenário em que você precise regularmente trabalhar com slices de inteiros e extrair a soma e média dos números dessa lista. Usando o que você aprendeu sobre métodos, qual seria sua solução para esse problema em Go?

	var c = conjunto{1, 2, 3, 4}

	soma := c.some()
	fmt.Println("soma:", soma)

	media := c.calculeMedia()
	fmt.Println("média:", media)

}

type Circulo struct {
	raio float64
}

func (c Circulo) AreaCirculo() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c Circulo) PerimetroCirculo() float64 {
	return 2 * math.Pi * c.raio
}

type conjunto []int

func (c conjunto) some() int {
	var soma int
	for _, numero := range c {
		soma += numero
	}
	return soma
}

func (c conjunto) calculeMedia() float64 {
	soma := c.some()
	media := float64(soma) / float64(len(c))
	return media
