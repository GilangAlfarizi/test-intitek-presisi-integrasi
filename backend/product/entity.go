package product

import "time"

type Product struct {
	ID        int
	Name      string
	Sku       string
	Quantity  int
	Location  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
