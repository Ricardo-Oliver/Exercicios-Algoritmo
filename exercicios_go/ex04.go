package main

//Desenvolva um algoritmo que leia dois números inteiros e mostre o somatório entre eles
import "fmt"

func main() {
	var num1 int
	var num2 int
	var resultado int

	fmt.Print("Digite um valor: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite outro valor: ")
	fmt.Scanln(&num2)

	resultado = num1 + num2

	fmt.Printf("A soma entre %d e %d é igual a %d.", num1, num2, resultado)
}

// EXEMPLO DE SAÍDA
// Digite um valor: 30
// Digite outro valor: 20
// A soma entre 30 e 20 é igual a 50.