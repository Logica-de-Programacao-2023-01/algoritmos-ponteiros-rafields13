// Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar. A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

package main

import "fmt"

func refreshPairOdd(verify *int) {
	if *verify%2 == 0 {
		*verify = 0
	} else {
		*verify = 1
	}
}

func main() {
	verify := 18
	refreshPairOdd(&verify)
	fmt.Print("O número se enquadra na opção: ", verify)
}
