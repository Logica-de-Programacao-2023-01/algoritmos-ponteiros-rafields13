// Implemente uma função que receba um ponteiro para uma slice e seu tamanho e preencha-o com os n primeiros números primos.

package main

import "fmt"

func fillPrimeNumbers(list *[]int, size int) {
	for i := 2; i < size; i++ {
		primeNumber := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				primeNumber = false
				break
			}
		}
		if primeNumber {
			*list = append(*list, i)
		}
	}
}

func main() {
	var list []int
	fillPrimeNumbers(&list, 14)
	fmt.Print(list)
}
