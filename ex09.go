package main

import "fmt"

//Escreva uma função que gere a sequência de Fibonacci até um determinado número n e retorne um mapa onde as chaves são
//os índices da sequência e os valores são os números correspondentes

func fibonacci(n int) map[int]int {
	seq := make(map[int]int)
	v1, v2 := 0, 1
	for i := 1; i <= n; i++ {
		seq[i] = v1
		v1, v2 = v2, v1+v2
	}
	return seq
}
func main() {
	tamanho := 10
	res := fibonacci(tamanho)
	fmt.Print(res)
}
