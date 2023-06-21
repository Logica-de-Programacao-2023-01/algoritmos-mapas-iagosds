package main

import "fmt"

//Escreva uma função que receba um slice de inteiros e retorne um mapa onde as chaves são os pares de números
//encontrados no slice e os valores são as frequências de cada par.

func pair(n []int) map[[2]int]int {
	freq := make(map[[2]int]int)
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			par := [2]int{n[i], n[j]}
			freq[par]++
		}
	}
	return freq
}
func main() {
	num := []int{1, 2, 3, 4, 4, 4, 3, 5, 2}
	res := pair(num)
	for pair, freq := range res {
		fmt.Printf("Par: %v - Frequência: %d\n", pair, freq)
	}
}
