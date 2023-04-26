package main

import "fmt"

//Faça um programa que leia um número inteiro e mostre o seu antecessor e seu sucessor

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	antecessor := numero - 1
	sucessor := numero + 1

	fmt.Printf("O antecessor de %d é %d\n", numero, antecessor)
	fmt.Printf("O sucessor de %d é %d", numero, sucessor)
}

// EXEMPLO DE SAÍDA
// Digite um número: 10
// O antecessor de 10 é 9
// O sucessor de 10 é 11