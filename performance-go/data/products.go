package data

import (
	"encoding/json"
	"io"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

// FromJSON .
func (p *Product) FromJSON(r io.Reader) error {

	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Water",
		Description: "Mineral Water",
		Price:       1.75,
	},
	&Product{
		ID:          2,
		Name:        "French Fries",
		Description: "Medium French Fries",
		Price:       3.55,
	},
	&Product{
		ID:          3,
		Name:        "Onion",
		Description: "Onion Rings",
		Price:       5.25,
	},
	&Product{
		ID:          4,
		Name:        "Egg",
		Description: "Fresh eggs",
		Price:       2.75,
	},
	&Product{
		ID:          5,
		Name:        "Coffee",
		Description: "Black Coffee",
		Price:       3.20,
	},
}
