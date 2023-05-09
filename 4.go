// Escreva uma função que receba um ponteiro para uma variável inteira e atualize o valor da variável para a soma dos valores dos seus dois últimos dígitos (por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).

package main

import "fmt"

func sumLastTwoDigits(number *int) {
	lastDigit := *number % 10
	penultimateDigit := (*number / 10) % 10
	*number = lastDigit + penultimateDigit
}

func main() {
	number := 12345
	fmt.Println(number)

	sumLastTwoDigits(&number)
	fmt.Print(number)
}
