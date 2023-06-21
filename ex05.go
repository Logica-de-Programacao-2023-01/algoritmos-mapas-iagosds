package main

import "fmt"

//Escreva uma função que receba uma string e retorne um mapa onde as chaves são os caracteres presentes na string e os
//valores são a frequência de cada caractere.

func main() {
	str := "anagrama"
	frq := letras(str)
	fmt.Print(frq)
}
func letras(s string) map[string]int {
	freq := make(map[string]int)
	for _, i := range s {
		freq[string(i)]++
	}
	return freq
}
