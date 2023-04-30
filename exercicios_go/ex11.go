package main

import (
	"fmt"
)

// Desenvolva uma lógica que leia os valores de A, B, e C de uma equação do segundo grau e mostre o valor de Delta.

func main() {
	var (
		a, b, c int
	)

	fmt.Print("Digite o valor de a: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o valor de b: ")
	fmt.Scanln(&b)

	fmt.Print("Digite o valor de c: ")
	fmt.Scanln(&c)

	delta := (b * b) - 4*a*c

	fmt.Printf("O valor de Delta é: %v", delta)
}

// EXEMPLO DE SAÍDA
// Digite o valor de a: 1
// Digite o valor de b: -3
// Digite o valor de c: -10
// O valor de Delta é: 49
