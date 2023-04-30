package main

import "fmt"

//Faça um algoritmo que leia a largura e altura de uma parede, calcule e mostre a área a ser pintada, e a
//quantidade de tinta necessária para o serviço, sabendo que cada litro de tinta pinta uma área de 2 metros quadrados

func main() {
	var (
		altura, largura float64
	)

	fmt.Print("Digite a largura da parede em metros: ")
	fmt.Scanln(&largura)

	fmt.Print("Agora digite a altura da parede em metros: ")
	fmt.Scanln(&altura)

	area := largura * altura
	quantidadeDeTinta := area / 2

	fmt.Printf("A área a ser pintada tem %.2f metros e será necessário %.1f litro(s) de tinta para pintar tudo!", area, quantidadeDeTinta)
}

// EXEMPLO DE SAÍDA
// Digite a largura da parede em metros: 3
// Agora digite a altura da parede em metros: 4
// A área a ser pintada tem 12.00 metros e será necessário 6.0 litro(s) de tinta para pintar tudo!
