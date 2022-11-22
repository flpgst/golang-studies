package main

type Conta struct {
	saldo int
}

func NewConta(saldo int) *Conta {
	return &Conta{saldo}
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	println("Saldo simulado", c.saldo)
	return c.saldo
}

func (c *Conta) creditar(valor int) int {
	c.saldo += valor
	return c.saldo
}
func main() {

	conta := NewConta(100)

	conta.simular(200)
	println("Saldo após simulação", conta.saldo)

	conta.creditar(200)
	println("Saldo após crédito", conta.saldo)

}
