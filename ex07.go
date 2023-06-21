package main

import (
	"fmt"
	"strings"
)

//Escreva uma função que receba uma string contendo uma frase e retorne um mapa onde as chaves são as palavras
//encontradas na frase e os valores são mapas contendo a contagem de cada letra em cada palavra

func mapa(frase string) map[string]map[string]int {
	retorno := make(map[string]map[string]int)
	cont_letras := make(map[string]int)

	palavras := strings.Split(frase, " ")
	fmt.Println(palavras)
	for _, palavra := range palavras {
		for _, i := range palavra {
			cont_letras[string(i)]++
		}
		retorno[palavra] = cont_letras
	}
	return retorno
}
func main() {
	str := "duas palavras"
	fmt.Println(str)
	res := mapa(str)
	fmt.Print(res)
}
