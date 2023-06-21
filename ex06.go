package main

import "fmt"

//Escreva uma função que receba uma lista de mapas, onde cada mapa contém a contagem de palavras de um texto, e retorne
//um único mapa contendo a soma de todas as contagens

func main() {
	var lista = []map[string]int{
		{"bla": 4, "alb": 1},
		{"alb": 1, "bla": 0, "bal": 1},
	}
	res := cont(lista)
	fmt.Print(res)
}
func cont(lis []map[string]int) map[string]int {
	retorno := make(map[string]int)
	for i := range lis {
		for b, c := range lis[i] {
			retorno[b] += c
		}
	}
	return retorno
}
