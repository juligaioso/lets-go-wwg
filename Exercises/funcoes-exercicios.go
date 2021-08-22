package main

import "fmt"

func main() {
	// ex 1
	// Declare e chame uma função que printa um cumprimento.

	Cumprimento()

	// ex 2
	// Continuando na solução do exercício anterior, agora faça com que a função que cumprimenta seja capaz de receber o nome da pessoa que está sendo cumprimentada, e printe o nome junto do cumprimento.

	Cumprimento2("Juli")

	// ex 3
	// Ainda na função que cumprimenta, agora faça com que ela seja capaz de receber o valor da hora e decidir o cumprimento com base no valor passado, de acordo com a tabela abaixo:
	Cumprimento3("Juli", 19)

	// ex 4
	// Agora vamos mudar a abordagem: em vez da função printar o cumprimento, ela vai gerar uma string com o cumprimento e retorná-la. Quem a chamar é que vai decidir o que fazer com o cumprimento.
	saudacao := GerandoSaudacao("Juli", 19)
	fmt.Println(saudacao)
}

func Cumprimento() {
	fmt.Println("Hello, world!")
}

func Cumprimento2(nome string) {
	fmt.Printf("Hello, world! %s!\n", nome)
}

func Cumprimento3(nome string, horario int) {
	saudacao := ""
	switch {
	case horario >= 18 && horario < 24:
		saudacao = "Boa noite"
	case horario >= 12 && horario < 18:
		saudacao = "Boa tarde"
	case horario >= 6 && horario < 12:
		saudacao = "Bom dia"
	case horario < 6:
		saudacao = "Boa madrugada"
	default:
		saudacao = "Olá"
	}
	fmt.Printf("%s!, %s!\n", saudacao, nome)
}


func GerandoSaudacao(nome string, horario int) string {
	saudacao := ""
	switch {
	case horario >= 18 && horario < 24:
		saudacao = "Boa noite"
	case horario >= 12 && horario < 18:
		saudacao = "Boa tarde"
	case horario >= 6 && horario < 12:
		saudacao = "Bom dia"
	case horario < 6:
		saudacao = "Boa madrugada"
	default:
		saudacao = "Olá"
	}

	frase := fmt.Sprintf("%s, %s!", saudacao, nome)
	return frase
}
