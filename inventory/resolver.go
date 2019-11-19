//go:generate go run github.com/99designs/gqlgen

package inventory

import (
	"context"
	"errors"
)

var _ ResolverRoot = (*resolver)(nil)

func NewResolver() ResolverRoot {
	ptrBool := func(b bool) *bool {
		return &b
	}

	return &resolver{
		inventory: []*Product{
			{
				Upc:     "1",
				InStock: ptrBool(true),
			},
			{
				Upc:     "2",
				InStock: ptrBool(false),
			},
			{
				Upc:     "3",
				InStock: ptrBool(true),
			},
		},
	}
}

type resolver struct {
	inventory []*Product
}

func (r *resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

func (r *resolver) Product() ProductResolver {
	return &productResolver{}
}

type entityResolver struct{ *resolver }

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*Product, error) {
	for _, product := range r.inventory {
		if product.Upc == upc {
			return product, nil
		}
	}

	return nil, nil
}

type productResolver struct {
}

func (r *productResolver) ShippingEstimate(ctx context.Context, obj *Product) (*int, error) {
	if obj == nil {
		return nil, nil
	}
	if obj.Price == nil {
		return nil, errors.New("price is required")
	}
	if obj.Weight == nil {
		return nil, errors.New("weight is required")
	}

	if *obj.Price > 1000 {
		var zero int
		return &zero, nil
	}

	estimate := int(float64(*obj.Weight) * 0.5)
	return &estimate, nil
}
