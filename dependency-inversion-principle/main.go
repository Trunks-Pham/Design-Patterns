package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) error
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Thanh toán %.2f VND qua PayPal.\n", amount)
	return nil
}

type VNPay struct{}

func (v *VNPay) Pay(amount float64) error {
	fmt.Printf("Thanh toán %.2f VND qua VNPay.\n", amount)
	return nil
}

type OrderService struct {
	Processor PaymentProcessor
}

func (o *OrderService) Checkout(amount float64) error {
	fmt.Println("Tiến hành thanh toán đơn hàng...")
	return o.Processor.Pay(amount)
}

func main() {
	var processor PaymentProcessor
	processor = &PayPal{}

	orderService := OrderService{
		Processor: processor,
	}

	err := orderService.Checkout(250000)
	if err != nil {
		fmt.Println("Lỗi thanh toán:", err)
	}
}
