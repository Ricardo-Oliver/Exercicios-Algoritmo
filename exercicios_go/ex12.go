package main

//Crie um programa que leia o preço de um produto, calcule e mostre o seu PREÇO PROMOCIONAL, com 5% de desconto
import "fmt"

func main() {
	var preco float64

	fmt.Print("Informe o preço do produto: ")
	fmt.Scanln(&preco)

	desconto := preco * 0.05

	precoPromocional := preco - desconto

	fmt.Printf("O preço com desconto, desse produto é: R$%.2f", precoPromocional)
}

// EXEMPLO DE SAÍDA
// Informe o preço do produto: 3400
// O preço com desconto, desse produto é: R$3230.00
