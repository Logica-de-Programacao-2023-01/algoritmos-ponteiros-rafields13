// Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular, e adicione um valor ao saldo da conta.

package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func moreCash(conta *Conta) {
	conta.Saldo = conta.Saldo + 1212
}

func main() {
	conta := Conta{
		Saldo:   3788,
		Titular: "Zézinho do Pagode",
	}
	fmt.Println(conta)

	moreCash(&conta)
	fmt.Print(conta)
}
