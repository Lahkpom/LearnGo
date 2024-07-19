package customer

// Customer is the struct of a client
type Customer struct {
	name    string
	address string
	phone   string
}

// New() returns a new customer
func New(name, address, phone string) *Customer {
	return &Customer{name, address, phone}
}

// SetName() sets the name of the customer
func (c *Customer) SetName(name string) { c.name = name }

// Name() returns the name of the customer
func (c *Customer) Name() string { return c.name }

// SetAddress() sets the address of the customer
func (c *Customer) SetAddress(address string) { c.address = address }

// Address() returns the address of the customer
func (c *Customer) Address() string { return c.address }

// SetPhone() sets the phone of the customer
func (c *Customer) SetPhone(phone string) { c.phone = phone }

// Phone() returns the phone of the customer
func (c *Customer) Phone() string { return c.phone }
