package item

import "fmt"

// Item contains the information of an invoice item
type Item struct {
	id      uint
	product string
	value   float64
}

// New() returns a new item
func New(id uint, product string, value float64) *Item {
	return &Item{id, product, value}
}

// SetId sets the id of the item
func (i *Item) SetId(id uint) { i.id = id }

// ID() returns the id of the item
func (i *Item) ID() uint { return i.id }

// SetProduct sets the product of the item
func (i *Item) SetProduct(product string) { i.product = product }

// Product() returns the product of the item
func (i *Item) Product() string { return i.product }

// SetValue sets the value of the item
func (i *Item) SetValue(value float64) { i.value = value }

// Value() returns the value of the item
func (i *Item) Value() float64 { return i.value }

//! CREAMOS ESTE TYPO ITEMS QUE ES UN SLICE DE ITEMS PARA ASIGNARLE TODO LO QUE TENGA QUE VER CON LOS ITEMS

// Items contains a list of items
type Items []Item

// AddItem() adds an item to the list
func (is Items) AddItems(itms []Item) {
	for _, v := range itms {
		is = append(is, v)
	}
}

// Total() returns the total of the items
func (is Items) Total() (total float64) {
	for _, v := range is {
		total += v.value
	}
	return
}

// ShowItems() shows the items
func (is Items) ShowItems() {
	fmt.Println("- Items:")
	for j, v := range is {
		fmt.Printf("%v. %v\t%v\n", j+1, v.Product(), v.Value())
	}
}
