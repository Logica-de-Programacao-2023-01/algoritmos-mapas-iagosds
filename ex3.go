package main

import "fmt"

//Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.

func main() {
	mapa := map[string]int{
		"valor1": 21,
		"valor2": 5,
		"valor3": 10,
	}
	res := soma(mapa)
	fmt.Print(res)
}

func soma(x map[string]int) int {
	soma := 0
	for i := range x {
		soma += x[i]
	}
	return soma
}
