package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {

	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoTales := contas.ContaPoupanca{}
	contaDoTales.Depositar(100)
	PagarBoleto(&contaDoTales, 60)

	fmt.Println(contaDoTales.ObterSaldo(), contaDoTales)

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(1000)
	PagarBoleto(&contaDaLuisa, 500)

	fmt.Println(contaDaLuisa.ObterSaldo(), contaDaLuisa)

}
