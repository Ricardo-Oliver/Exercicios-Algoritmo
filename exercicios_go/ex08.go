package main

import "fmt"

// Desenvolva um programa que leia uma distância em metros e mostre os valores relativos em outras medidas
// Ex: 185.7m, 0.18572Km, 1.8572Hm, 18.572Dam, 1857.2dm, 18572.0cm
func main() {
	var metros float64
	
	fmt.Print("Digite uma distância em metros: ")
	fmt.Scanln(&metros)

	km := metros / float64(1000)
	hm := metros / float64(100)
	dam := metros / float64(10)
	dm := metros * float64(10)
	cm := metros * float64(100)
	mm := metros * float64(1000)

	fmt.Printf("A distância de %.2fm corresponde a: \n%.5fkm\t\t\t%.1fdm\n%.4fhm\t\t\t%.1fcm\n%.3fdam\t\t\t%.1fmm", metros, km, dm, hm, cm, dam, mm)

}

// EXEMPLO DE SAÍDA
// Digite uma distância em metros: 185.72
// A distância de 185.72m corresponde a: 
// 0.18572km                       1857.2dm
// 1.8572hm                        18572.0cm
// 18.572dam                       185720.0mm