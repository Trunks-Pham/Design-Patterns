package main

import (
	"fmt"
	"time"
)

type PaymentCommand interface {
	Execute() string
}

type PaymentReceiver struct{}

func (pr *PaymentReceiver) processCreditCard(amount float64) string {
	return fmt.Sprintf("Đã xử lý thanh toán thẻ tín dụng: %.2f VND", amount)
}

func (pr *PaymentReceiver) processBankTransfer(amount float64) string {
	return fmt.Sprintf("Đã xử lý chuyển khoản ngân hàng: %.2f VND", amount)
}

type CreditCardPayment struct {
	receiver *PaymentReceiver
	amount   float64
}

func (cc *CreditCardPayment) Execute() string {
	return cc.receiver.processCreditCard(cc.amount)
}

type BankTransferPayment struct {
	receiver *PaymentReceiver
	amount   float64
}

func (bt *BankTransferPayment) Execute() string {
	return bt.receiver.processBankTransfer(bt.amount)
}

type PaymentInvoker struct {
	commands []PaymentCommand
}

func (pi *PaymentInvoker) AddCommand(command PaymentCommand) {
	pi.commands = append(pi.commands, command)
}

func (pi *PaymentInvoker) ProcessPayments() []string {
	var results []string
	for _, command := range pi.commands {
		result := command.Execute()
		results = append(results, result)
		time.Sleep(500 * time.Millisecond)
	}
	return results
}

func main() {
	receiver := &PaymentReceiver{}

	creditCardPayment := &CreditCardPayment{
		receiver: receiver,
		amount:   1500000.00,
	}

	bankTransferPayment := &BankTransferPayment{
		receiver: receiver,
		amount:   2500000.00,
	}

	invoker := &PaymentInvoker{}
	invoker.AddCommand(creditCardPayment)
	invoker.AddCommand(bankTransferPayment)

	fmt.Println("Bắt đầu xử lý các thanh toán:")
	results := invoker.ProcessPayments()

	for i, result := range results {
		fmt.Printf("Giao dịch %d: %s\n", i+1, result)
	}
}
