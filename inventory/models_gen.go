// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package inventory

type Product struct {
	Upc              string `json:"upc"`
	Weight           *int   `json:"weight"`
	Price            *int   `json:"price"`
	InStock          *bool  `json:"inStock"`
	ShippingEstimate *int   `json:"shippingEstimate"`
}

func (Product) IsEntity() {}
