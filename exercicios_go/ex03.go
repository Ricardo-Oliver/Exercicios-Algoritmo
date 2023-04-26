package main

import (
	"fmt"
	"strings"
)

//Crie um programa que leia o nome e o salário de um funcionário, mostrando no final uma mensagem

func main() {
	var (
		nome string
		salario float64
	)

	fmt.Println("Qual o nome do(a) funcionário(a) ?")
	fmt.Scanln(&nome)
	fmt.Println("Qual o salário do(a) funcionário(a) ?")
	fmt.Scanln(&salario)

	fmt.Printf("O(a) funcionário(a) %s tem um salário de R$%.2f em Junho.", strings.Title(strings.ToLower(nome)), salario)
}

// EXEMPLO DE SAÍDA
// Qual o nome do(a) funcionário(a) ?
// Ricardo
// Qual o salário do(a) funcionário(a) ?
// 575
// O(a) funcionário(a) Ricardo tem um salário de R$575.00 em Junho.