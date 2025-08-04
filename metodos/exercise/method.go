package main

import "fmt"

type conta struct {
	saldo         float64
	numeroDaConta string
}

func (c *conta) depositar(valor float64) error {

	if valor <= 0 {
		return fmt.Errorf("Não é possível depositar valor negativo")
	}
	c.saldo += valor

	return nil
}

func (c *conta) sacar(valor float64) (error, bool) {
	if valor <= 0 {
		return fmt.Errorf("Não é possível sacar um valor negativo"), false
	} else if valor > c.saldo {
		return fmt.Errorf("Não é possível sacar um valor maior que o saldo"), false
	}

	c.saldo -= valor

	return nil, true
}

func (c *conta) obterSaldo() float64 {
	return c.saldo
}

func main() {
	count1 := conta{saldo: 2214.95}
	fmt.Println(count1.obterSaldo())
	count1.sacar(-45)
	fmt.Println(count1.obterSaldo())
	count1.sacar(10)
	fmt.Println(count1.obterSaldo())
	count1.depositar(20)
	count1.sacar(3000)
	fmt.Println(count1.obterSaldo())
	// method fica intrisicamente ligado a struct
}
