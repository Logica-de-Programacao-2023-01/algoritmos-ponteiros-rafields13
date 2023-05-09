// Escreva uma função em Go que receba um ponteiro para uma variável float64 e atualize o valor da variável para a média aritmética entre o seu valor atual e o valor da constante Pi.

package main

import (
	"fmt"
	"math"
)

func mediaAritmetica(number *float64) {
	*number = (*number + math.Pi) / 2
}

func main() {
	number := 6.86
	fmt.Println(number)
	mediaAritmetica(&number)
	fmt.Print(number)
}
