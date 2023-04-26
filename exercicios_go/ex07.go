package main

import "fmt"

//Crie um algoritmo que leia um número real e mostre na tela o seu dobro e a sua terça parte

func main() {
	var numero float64
	
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	dobro := numero * 2
	terco := numero / 3

	fmt.Printf("O dobro de %.1f é %.1f\n", numero, dobro)
	fmt.Printf("A terça parte de %.1f é %.4f", numero, terco)
}

// EXEMPLO DE SAÍDA
// Digite um número: 3.5
// O dobro de 3.5 é 7.0
// A terça parte de 3.5 é 1.1667