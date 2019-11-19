//go:generate go run github.com/99designs/gqlgen

package accounts

import (
	"context"
)

var _ ResolverRoot = (*resolver)(nil)

func NewResolver() ResolverRoot {
	ptrString := func(s string) *string {
		return &s
	}
	ptrInt := func(i int) *int {
		return &i
	}

	return &resolver{
		products: []*Product{
			{
				Upc:    "1",
				Name:   ptrString("Table"),
				Price:  ptrInt(899),
				Weight: ptrInt(100),
			},
			{
				Upc:    "2",
				Name:   ptrString("Couch"),
				Price:  ptrInt(1299),
				Weight: ptrInt(1000),
			},
			{
				Upc:    "3",
				Name:   ptrString("Chair"),
				Price:  ptrInt(54),
				Weight: ptrInt(50),
			},
		},
	}
}

type resolver struct {
	products []*Product
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

type queryResolver struct{ *resolver }

func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*Product, error) {
	if first == nil {
		return r.products, nil
	}

	bottom := *first
	if bottom >= len(r.products) {
		bottom = len(r.products)
	}
	return r.products[0:bottom], nil
}

type entityResolver struct{ *resolver }

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*Product, error) {
	for _, product := range r.products {
		if product.Upc == upc {
			return product, nil
		}
	}

	return nil, nil
}
