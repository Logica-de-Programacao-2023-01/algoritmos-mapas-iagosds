package main

import "fmt"

//Escreva uma função que receba um mapa onde as chaves são nomes de pessoas e os valores são as quantias de dinheiro
//que cada pessoa gastou em uma conta compartilhada. A função deve calcular o valor que cada pessoa deve receber ou
//pagar para igualar as despesas

func despesas(gastos map[string]int) map[string]int {
	retorno := make(map[string]int)
	soma := 0
	for i := range gastos {
		soma += gastos[i]
	}
	soma /= len(gastos)
	for i := range gastos {
		retorno[i] = gastos[i] - soma
	}
	return retorno
}
func main() {
	contas := map[string]int{
		"iago":    300,
		"jut":     400,
		"sua mae": 550,
	}
	res := despesas(contas)
	fmt.Println(res)
}
