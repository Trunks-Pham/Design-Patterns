package main

import (
	"fmt"
)

type PaymentProcessor interface {
	Pay(amount float64) (string, error)
}

type CreditCard struct {
	SoThe string
}

func (cc CreditCard) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Đã thanh toán %.2f VND bằng thẻ tín dụng %s", amount, cc.SoThe), nil
}

type MTPEBank struct {
	Email string
}

func (pp MTPEBank) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Đã thanh toán %.2f VND bằng MTPEBank %s", amount, pp.Email), nil
}

type BankTransfer struct {
	SoTaiKhoan string
}

func (bt BankTransfer) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Đã thanh toán %.2f VND bằng chuyển khoản %s", amount, bt.SoTaiKhoan), nil
}

func Service(processor PaymentProcessor, amount float64) (string, error) {
	return processor.Pay(amount)
}

func main() {
	creditCard := CreditCard{SoThe: "1234-5678-9012-3456"}
	mtpebank := MTPEBank{Email: "thaopm@paypay.com"}
	bankTransfer := BankTransfer{SoTaiKhoan: "987654321"}

	result, _ := Service(creditCard, 100.50)
	fmt.Println(result)

	result, _ = Service(mtpebank, 75.25)
	fmt.Println(result)

	result, _ = Service(bankTransfer, 200.00)
	fmt.Println(result)
}
