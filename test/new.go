

package main

type Order struct {
	ID     int
	Amount float64
	Status string
}
func (o *Order) changeStatus(newStatus string) {
	o.Status = newStatus
}
func main() {
	order := Order{
		ID:     1,
		Amount: 99.99,
		Status: "Pending",
	}
	println(order.ID, order.Amount, order.Status)
	order.changeStatus("Shipped")
	println(order.ID, order.Amount, order.Status)
	
}
