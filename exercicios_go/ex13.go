package main

//Faça um algoritmo que leia o salário de um funcionario, calcule e mostre o seu novo salário, com 15% de aumento
import "fmt"

func main() {
	var salario float64

	fmt.Print("Informe o salário do funcionário: ")
	fmt.Scanln(&salario)

	aumento := salario * 0.15
	novoSalario := salario + aumento

	fmt.Printf("O novo salário do funcionário é: R$%.2f", novoSalario)
}
