package main

import "fmt"

func main() {
	// ex 1
	// Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
	// Declare uma variável e atribua a ela o ano atual.
	// Escreva um programa que verifique se no ano atual essa pessoa já está apta a votar e que printe na tela uma mensagem informando caso positivo e caso negativo.
	anoNascimento := 2011
	anoAtual := 2021

	if anoAtual-anoNascimento > 16 {
		fmt.Println("Está apto para votar")
		return
	}
	fmt.Println("Não está apto para votar")

	// ex 2
	// Faça um programa que:
	// Declara uma variável;
	// Atribui o valor de um número a ela;
	// Informa se o número é positivo ou negativo.
	numero := 2

	if numero >= 0 {
		fmt.Println("Este número é positivo")
		return
	}
	fmt.Println("Este número é negativo")

	// ex 3
	// Faça um programa que, dada a idade de uma pessoa, verifique se ela é menor de idade, maior de idade ou idosa utilizando a instrução if.
	idade := 12

	if idade < 18 {
		fmt.Println("Menor de idade")
		return
	} else if idade < 60 {
		fmt.Println("Maior de idade")
		return
	}
	fmt.Println("Idoso")

	// ex 4
	// Faça um programa a partir das mesmas orientações do exercício acima, mas utilizando switch em vez de if.
	idade := 12

	switch {
	case idade >= 60:
		fmt.Println("Idoso")
	case idade >= 18 && idade < 60:
		fmt.Println("Maior de idade")
	default:
		fmt.Printf("Menor de idade")
	}

	// ex 5
	// Declare uma variável chamada hora e atribuir um valor int a ela.
	// Usando switch, descreva cases e printe na tela se a hora corresponde ao período da manhã, tarde, noite ou madrugada.
	hora := 12

	switch {
	case hora >= 18 && hora < 24:
		fmt.Println("Noite")
	case hora >= 12 && hora < 18:
		fmt.Println("Tarde")
	case hora >= 6 && hora < 12:
		fmt.Println("Manhã")
	default:
		fmt.Printf("Madrugada")
	}
}
