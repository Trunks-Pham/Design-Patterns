package main

import (
	"fmt"
)

type PaymentExecutor interface {
	ThucHien(amount float64) (string, error)
}

type Refundable interface {
	HoanTien(amount float64) (string, error)
}

type CreditCard struct {
	SoThe string
}

func (cc CreditCard) ThucHien(amount float64) (string, error) {
	return fmt.Sprintf("Đã thanh toán %.2f VND bằng thẻ tín dụng %s", amount, cc.SoThe), nil
}

func (cc CreditCard) HoanTien(amount float64) (string, error) {
	return fmt.Sprintf("Đã hoàn %.2f VND cho thẻ tín dụng %s", amount, cc.SoThe), nil
}

type PayPal struct {
	Email string
}

func (pp PayPal) ThucHien(amount float64) (string, error) {
	return fmt.Sprintf("Đã thanh toán %.2f VND bằng PayPal %s", amount, pp.Email), nil
}

func (pp PayPal) HoanTien(amount float64) (string, error) {
	return fmt.Sprintf("Đã hoàn %.2f VND cho PayPal %s", amount, pp.Email), nil
}

type BankTransfer struct {
	SoTaiKhoan string
}

func (bt BankTransfer) ThucHien(amount float64) (string, error) {
	return fmt.Sprintf("Đã thanh toán %.2f VND bằng chuyển khoản %s", amount, bt.SoTaiKhoan), nil
}

func ThanhToan(executor PaymentExecutor, amount float64) (string, error) {
	return executor.ThucHien(amount)
}

func HoanTien(refund Refundable, amount float64) (string, error) {
	return refund.HoanTien(amount)
}

func main() {
	creditCard := CreditCard{SoThe: "1234-5678-9012-3456"}
	paypal := PayPal{Email: "user@example.com"}
	bankTransfer := BankTransfer{SoTaiKhoan: "987654321"}

	result, _ := ThanhToan(creditCard, 100.50)
	fmt.Println(result)

	result, _ = ThanhToan(paypal, 75.25)
	fmt.Println(result)

	result, _ = ThanhToan(bankTransfer, 200.00)
	fmt.Println(result)

	result, _ = HoanTien(creditCard, 50.00)
	fmt.Println(result)

	result, _ = HoanTien(paypal, 25.00)
	fmt.Println(result)
}