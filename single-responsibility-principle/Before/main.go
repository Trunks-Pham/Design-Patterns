package main

import "fmt"

type Order struct {
    OrderID string
    Amount  float64
}

func (o *Order) ProcessPayment() string { 
    fmt.Printf("Đang xử lý thanh toán cho đơn hàng %s với số tiền %.2f\n", o.OrderID, o.Amount) 
    paymentStatus := "Success"
    if paymentStatus == "Success" {
        o.SendConfirmationEmail()
    }
    return paymentStatus
}

func (o *Order) SendConfirmationEmail() { 
    fmt.Printf("Đang gửi email xác nhận cho đơn hàng %s\n", o.OrderID)
}

func main() {
    order := Order{OrderID: "123", Amount: 100.0}
    order.ProcessPayment()
}