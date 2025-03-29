package main

import "fmt"

type Payment struct {
	ID     int
	Amount float64
}

type PaymentCollection struct {
	payments []Payment
}

type Iterator interface {
	HasNext() bool
	Next() *Payment
}

type PaymentIterator struct {
	collection *PaymentCollection
	index      int
}

func (i *PaymentIterator) HasNext() bool {
	return i.index < len(i.collection.payments)
}

func (i *PaymentIterator) Next() *Payment {
	if i.HasNext() {
		payment := &i.collection.payments[i.index]
		i.index++
		return payment
	}
	return nil
}

func (c *PaymentCollection) CreateIterator() Iterator {
	return &PaymentIterator{
		collection: c,
		index:      0,
	}
}

func (c *PaymentCollection) AddPayment(payment Payment) {
	c.payments = append(c.payments, payment)
}

func main() {
	collection := &PaymentCollection{}

	collection.AddPayment(Payment{ID: 1, Amount: 100.50})
	collection.AddPayment(Payment{ID: 2, Amount: 250.75})
	collection.AddPayment(Payment{ID: 3, Amount: 75.25})
	collection.AddPayment(Payment{ID: 4, Amount: 300.00})
	collection.AddPayment(Payment{ID: 5, Amount: 150.00})
	collection.AddPayment(Payment{ID: 6, Amount: 200.00})
	collection.AddPayment(Payment{ID: 7, Amount: 50.00})
	collection.AddPayment(Payment{ID: 8, Amount: 125.00})
	collection.AddPayment(Payment{ID: 9, Amount: 175.00})

	iterator := collection.CreateIterator()

	fmt.Println("Danh sách các con nợ của Thảo")
	for iterator.HasNext() {
		payment := iterator.Next()
		fmt.Printf("Đòi nợ người dùng: %d, số tiền bị nó thiếu: %.2f\n", payment.ID, payment.Amount)
	}
}
