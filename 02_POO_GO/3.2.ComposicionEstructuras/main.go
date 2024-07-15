package main

import (
	"fmt"
	//
	"github.com/Lahkpom/myModules/customer"
	"github.com/Lahkpom/myModules/invoice"
	"github.com/Lahkpom/myModules/invoiceitem"
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
		3000,
	)

	invoice1 := createInvoice(
		"Argentina",
		"Buenos Aires",
	)

	invoice1.SetCliente(client1)
	invoice1.AddItem(item1, item2, item3)

	fmt.Println(invoice1)
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
func createItem(id uint, product string, value float64) *invoiceitem.InvoiceItem {
	// Debería tener el mismo formato que el otro
	return invoiceitem.New(id, product, value)
}

// createInvoice() returns an invoice
func createInvoice(country, city string) *invoice.Invoice {
	// Debería tener el mismo formato que el otro
	return invoice.New(country, city)
}
