package main

import "fmt"

type Order struct {
    OrderID string
    Amount  float64
}
type PaymentProcessor struct{}

func (p *PaymentProcessor) Process(order Order) string {
    fmt.Printf("Đang xử lý thanh toán cho đơn hàng %s với số tiền %.2f\n", order.OrderID, order.Amount)
    return "Success"
}
 
type EmailService struct{}

func (e *EmailService) SendConfirmation(order Order) {
    fmt.Printf("Đang gửi email xác nhận cho đơn hàng %s\n", order.OrderID)
}

func main() {
    order := Order{OrderID: "123", Amount: 100.0}
    paymentProcessor := PaymentProcessor{}
    emailService := EmailService{}

    paymentStatus := paymentProcessor.Process(order)
    if paymentStatus == "Success" {
        emailService.SendConfirmation(order)
    }
}