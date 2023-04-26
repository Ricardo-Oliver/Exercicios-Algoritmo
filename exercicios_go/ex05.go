package main

import "fmt"

//Faça um programa que leia as duas notas de um aluno em uma matéria e mostre na tela a sua média na disciplina

func main() {
	var nota1 float64
	var nota2 float64
	var media float64

	fmt.Print("Nota 1: ")
	fmt.Scanln(&nota1)
	fmt.Print("Nota 2: ")
	fmt.Scanln(&nota2)

	media = (nota1 + nota2) / 2

	fmt.Printf("A media entre %.1f e %.1f é igual a %.1f", nota1, nota2, media)
}

// EXEMPLO DE SAÍDA
// Nota 1: 8
// Nota 2: 7
// A media entre 8.0 e 7.0 é igual a 7.5