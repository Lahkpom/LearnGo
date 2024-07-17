package invoiceitem

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
