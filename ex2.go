package main

import "fmt"

//Escreva uma função que receba dois mapas e retorne um novo mapa contendo todos os elementos dos mapas de entrada.
//Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.

func integrar(x map[string]int, y map[string]int) map[string]int {
	for chavex := range x {
		for chavey := range y {
			if chavex == chavey {
				x[chavex] = y[chavey]
			} else {
				x[chavey] = y[chavey]
			}
		}
	}
	return x
}
func main() {
	mapa := map[string]int{
		"laranja": 10,
	}
	mapinha := map[string]int{
		"laranja": 15,
		"banana":  10,
	}
	res := integrar(mapa, mapinha)
	fmt.Print(res)
}
