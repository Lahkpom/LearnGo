package main

import (
	"fmt"
	"main/modules/customer"
	"main/modules/invoice"
	"main/modules/invoiceitem"
)

func main() {
	//En realidad tendría que ya haber una lista de clientes y verificar si ya existe
	client1 := createCustomer()

	item1 := createItem(
		001,
		"Computer",
		100000,
	)

	item2 := createItem(
		002,
		"Mouse",
		2000,
	)

	item3 := createItem(
		003,
		"Keyboard",
		3000.97,
	)

	invoice1 := createInvoice(
		"Argentina",
		"Buenos Aires",
	)
	invoice1.SetClient(client1)

	invoice1.AddItem(*item1, *item2, *item3)

	invoice1.ShowItems()
	fmt.Println("Total:", invoice1.Total())
}

// createCustomer() returns a customer
func createCustomer() *customer.Customer {
	// Acá habría que pedir la información por consola
	name := "Leonel Alejandro HIDALGO"
	address := "118.hidalgo.leonel.1358@gmail.com"
	phone := "1151407622"
	return customer.New(name, address, phone)
}

// createItem() returns an item
func createItem(id uint, product string, value float64) *invoiceitem.Item {
	// Debería tener el mismo formato que el otro
	return invoiceitem.New(id, product, value)
}

// createInvoice() returns an invoice
func createInvoice(country, city string) *invoice.Invoice {
	// Debería tener el mismo formato que el otro
	return invoice.New(country, city)
}
