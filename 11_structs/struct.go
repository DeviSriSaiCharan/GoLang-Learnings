package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	updatedAt time.Time
}

// This is a constructor function, it returns a pointer to an order struct
func newOrder(id string, amount float32) *order {
	order := order{
		id:        id,
		amount:    amount,
		status:    "pending",
		createdAt: time.Now(),
	}
	return &order
}

// This is a method, it has a receiver (o order) and a name (isPending)
// Since are not modifying the order, we can use a value receiver
func (o order) isPending() bool {
	return o.status == "pending"
}

// This is a method, it has a receiver (o *order) and a name (changeStatus)
// Since we are modifying the order, we need to use a pointer receiver
func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

func main() {
	var order1 order = order{}
	order1.id = "1"
	order1.amount = 100.0
	order1.status = "pending"
	order1.createdAt = time.Now()

	order2 := order{
		id:        "2",
		amount:    200.0,
		status:    "pending",
		createdAt: time.Now(),
	}

	order2.updatedAt = time.Now()

	fmt.Println(order1)
	fmt.Println(order2)

	order3 := newOrder("3", 300.0)
	fmt.Println(order3)

	fmt.Println(order3.isPending()) // true

	order3.changeStatus("completed")
	fmt.Println(order3.isPending()) // false
}
