package main

import "fmt"

//A locadora de carros precisa da sua ajuda para cobrar seus serviços. Escreva um programa que
//pergunte a quantidade de Km percorridos por um carro alugado e a quantidade de dias pelos quais ele foi alugado.
//Calcule o preço total a pagar, sabendo que o carro custa R$90 por dia e R$0,20 por Km rodado.

func main() {
	var (
		kmPercorrido, diaAlugado int 
	)

	fmt.Print("Informe a quantidade de Km percorridos: ")
	fmt.Scanln(&kmPercorrido)
	fmt.Print("Informe a quantidade de dias alugado: ")
	fmt.Scanln(&diaAlugado)

	valorDosDias := diaAlugado * 90.0
	valorDoKm := float64(kmPercorrido) * 0.20

	valorTotal := valorDoKm + float64(valorDosDias)

	fmt.Printf("O valor total do aluguel é de R$%.2f", valorTotal)
}

// EXEMPLO DE SAÍDA
// Informe a quantidade de Km percorridos: 100
// Informe a quantidade de dias alugado: 2
// O valor total do aluguel é de R$200.00