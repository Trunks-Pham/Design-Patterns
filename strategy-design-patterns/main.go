package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type Momo struct {
	Momo string
}

func (c *Momo) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Momo %s", amount, c.Momo)
}

type Cash struct{}

func (c *Cash) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f in Cash", amount)
}

type GoodBye struct {
	running string
}

func (g *GoodBye) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f by %s", amount, g.running)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64) string {
	return pc.strategy.Pay(amount)
}

func main() {
	payment := &PaymentContext{}

	payment.SetStrategy(&Momo{Momo: "0999999999"})
	fmt.Println(payment.ExecutePayment(100.10))

	payment.SetStrategy(&Cash{})
	fmt.Println(payment.ExecutePayment(50.00))

	payment.SetStrategy(&GoodBye{running: "GoodBye"})
	fmt.Println(payment.ExecutePayment(100.00))
}
