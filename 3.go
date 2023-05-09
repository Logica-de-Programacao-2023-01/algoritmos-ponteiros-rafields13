// Escreva uma função que receba um ponteiro para uma string e atualize o valor da string para seu reverso.

package main

import "fmt"

func reverse(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}

func main() {
	str := "Rafael Marliere de Oliveira"
	fmt.Println("Valor antes da inversão:", str)

	reverse(&str)
	fmt.Print("Valor após a inversão: ", str)
}
