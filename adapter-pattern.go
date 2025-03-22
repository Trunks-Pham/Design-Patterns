package main

import "fmt"

// PaymentGateway là giao diện chung cho các phương thức thanh toán
type PaymentGateway interface {
	Pay(amount float64)
}

// BankPayment là hệ thống thanh toán của ngân hàng
type BankPayment struct{}

func (b *BankPayment) Pay(amount float64) {
	fmt.Printf("Thanh toán qua ngân hàng: %.2f VND\n", amount)
}

// MoMoPayment là hệ thống thanh toán MoMo
type MoMoPayment struct{}

func (m *MoMoPayment) MoMoPay(amount float64) {
	fmt.Printf("Thanh toán qua MoMo: %.2f VND\n", amount)
}

// MoMoAdapter là adapter giúp MoMo tương thích với hệ thống ngân hàng
type MoMoAdapter struct {
	momo *MoMoPayment
}

func (a *MoMoAdapter) Pay(amount float64) {
	a.momo.MoMoPay(amount)
}

func main() {
	// Sử dụng hệ thống thanh toán qua ngân hàng
	bankPayment := &BankPayment{}
	bankPayment.Pay(100000)

	// Sử dụng hệ thống thanh toán MoMo thông qua adapter
	momoPayment := &MoMoPayment{}
	momoAdapter := &MoMoAdapter{momo: momoPayment}
	momoAdapter.Pay(150000)
}
