package main

import "fmt"

//Crie um programa que leia o número de dias trabalhados em um mês e mostre o salário de um funcionário,
// sabendo que ele trabalha 8 horas por dia e ganha R$25,00 por hora trabalhada

func main() {
	var (
		diasTrabalhados int
	)

	fmt.Print("Informe quantos dias o funcionário trabalhou esse mês: ")
	fmt.Scanln(&diasTrabalhados)

	valorDiaria := 8 * 25
	salario := diasTrabalhados * valorDiaria

	fmt.Printf("O salário do funcionário é de R$%v", salario)
}

// EXEMPLO DE SAÍDA
// Informe quantos dias o funcionário trabalhou esse mês: 18
// O salário do funcionário é de R$3600