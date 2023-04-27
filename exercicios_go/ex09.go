package main

import "fmt"

//Faça um algoritmo que leia quanto dinheiro uma pessoa tem na carteira em R$ e mostre quantos dólares ela pode comprar.
//Considere US$1,00 = R$3,45
func main() {
	var dinheiroNaCarteira float64

	fmt.Print("Digite quanto você tem em reais: ")
	fmt.Scanln(&dinheiroNaCarteira)

	emDolares := dinheiroNaCarteira / 3.45

	fmt.Printf("Você pode comprar US$%.2f dólares", emDolares)

}