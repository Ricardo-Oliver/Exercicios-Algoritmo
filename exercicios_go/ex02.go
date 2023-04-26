package main

import "fmt"

//Faça um programa que leia o nome de uma pessoa e mostre uma mensagem de boas-vindas para ela

func main() {
	var nome string 

	fmt.Println("Qual é seu nome ?")
	fmt.Scanln(&nome)
	fmt.Printf("Olá %s, é um prazer te conhecer!", nome)
}

// EXEMPLO DE SAÍDA 
// Ricardo
// Olá Ricardo, é um prazer te conhecer!