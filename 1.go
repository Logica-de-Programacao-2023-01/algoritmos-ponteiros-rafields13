// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro ``n``. A função deve atualizar o valor do inteiro para a soma dos ``n`` primeiros números naturais.

package main

import "fmt"

func refreshValue(n int, soma *int) {
	*soma = 0
	for i := 1; i <= n; i++ {
		*soma += i
	}
}

func main() {
	var soma int
	refreshValue(10, &soma)
	fmt.Print(soma)
}
