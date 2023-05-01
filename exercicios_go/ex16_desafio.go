package main

import "fmt"

// [DESAFIO] Escreva um programa para calcular a redução do tempo de vida de um fumante.
// Pergunte a quantidade de cigarros fumados por dias e quantos anos ele já fumou.
// Considere que um fumante perde 10 min de vida a cada cigarro. Calcule quantos dias

func main() {
	var qtdCigarrosDia, anosDeFumante int

	fmt.Print("Informe a quantidade de cigarros fumados por dia: ")
	fmt.Scanln(&qtdCigarrosDia)
	fmt.Print("Informe a quantidade de anos de fumante: ")
	fmt.Scanln(&anosDeFumante)

	tempoDeFumante := anosDeFumante * 365
	qtdDeCigarrosFumados := qtdCigarrosDia * tempoDeFumante
	totalPerdidoEmMinutos := qtdDeCigarrosFumados * 10
	totalPerdidoEmHoras := totalPerdidoEmMinutos / 60
	totalPerdidosEmDias := totalPerdidoEmHoras / 24

	fmt.Printf("O fumante perdeu %v dias de vida.", totalPerdidosEmDias)

}
