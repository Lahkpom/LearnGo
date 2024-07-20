package invoice

import (
	"fmt"
	"main/modules/customer"
	"main/modules/item"
)

// Invoice contains the information of an invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  *customer.Customer
	items   item.Items // type Items contains a list of items
}

// New() returns a new invoice
func New(country, city string) *Invoice {
	return &Invoice{country, city, 0, &customer.Customer{}, item.Items{}}
}

// SetCountry sets the country of the invoice
func (i *Invoice) SetCountry(country string) { i.country = country }

// Country() returns the country of the invoice
func (i *Invoice) Country() string { return i.country }

// SetCity sets the city of the invoice
func (i *Invoice) SetCity(city string) { i.city = city }

// City() returns the city of the invoice
func (i *Invoice) City() string { return i.city }

// Total() returns the total of the invoice
func (i *Invoice) Total() float64 { return i.total }

// SetClient sets the client of the invoice
func (i *Invoice) SetClient(client *customer.Customer) { i.client = client }

// Client() returns the client's pointer of the invoice
func (i *Invoice) Client() *customer.Customer { return i.client }

// Items() returns the items of the invoice
func (i *Invoice) Items() item.Items { return i.items }

// AddItems() adds an item to the invoice
func (i *Invoice) AddItems(itms ...item.Item) {
	i.items.AddItems(itms)
	// No pude poner esta lógica en item.Items porque no me modifica el puntero original
	i.items = append(i.items, itms...)
	i.setTotal()
}

// setTotal() returns the total of the invoice at the moment
func (i *Invoice) setTotal() {
	i.total = i.items.Total()
}

// TODO RemoveItem() removes an item from the invoice
/*
func (i *Invoice) RemoveItem(items ...invoiceitem.Item) {
	for _, v := range items {
		for _, i := range i.items {
			if v.ID() == i.ID() {
				i.items = (LÓGICA PARA REMOVER)
				i.total -= v.Value()
			}
		}
	}
}
*/

// ShowItems() shows the items of the invoice
func (i *Invoice) ShowItems() {
	fmt.Printf(("- Customer: %v\n"), i.client.Name())
	i.items.ShowItems()
}
