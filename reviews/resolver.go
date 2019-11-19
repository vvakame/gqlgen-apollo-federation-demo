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

	return &resolver{
		Products: []*Product{
			{
				Upc: "1",
			},
			{
				Upc: "2",
			},
			{
				Upc: "3",
			},
		},
		Reviews: []*Review{
			{
				ID:         "1",
				AuthorID:   "1",
				ProductUPC: "1",
				Body:       "Love it!",
			},
			{
				ID:         "2",
				AuthorID:   "1",
				ProductUPC: "2",
				Body:       "Too expensive.",
			},
			{
				ID:         "3",
				AuthorID:   "2",
				ProductUPC: "3",
				Body:       "Could be better.",
			},
			{
				ID:         "4",
				AuthorID:   "2",
				ProductUPC: "1",
				Body:       "Prefer something else.",
			},
		},
		Usernames: []*User{
			{
				ID:       "1",
				Username: ptrString("@ada"),
			},
			{
				ID:       "2",
				Username: ptrString("@complete"),
			},
		},
	}
}

type resolver struct {
	Products  []*Product
	Reviews   []*Review
	Usernames []*User
}

func (r *resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

func (r *resolver) Product() ProductResolver {
	return &productResolver{r}
}

func (r *resolver) Review() ReviewResolver {
	return &reviewResolver{r}
}

func (r *resolver) User() UserResolver {
	return &userResolver{r}
}

type entityResolver struct{ *resolver }

func (r *entityResolver) FindReviewByID(ctx context.Context, id string) (*Review, error) {
	for _, review := range r.resolver.Reviews {
		if review.ID == id {
			return review, nil
		}
	}

	return nil, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*User, error) {
	for _, user := range r.Usernames {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*Product, error) {
	for _, product := range r.Products {
		if product.Upc == upc {
			return product, nil
		}
	}

	return nil, nil
}

type productResolver struct{ *resolver }

func (r *productResolver) Reviews(ctx context.Context, obj *Product) ([]*Review, error) {
	var resp []*Review
	for _, review := range r.resolver.Reviews {
		if review.ProductUPC == obj.Upc {
			resp = append(resp, review)
		}
	}

	return resp, nil
}

type reviewResolver struct{ *resolver }

func (r *reviewResolver) Author(ctx context.Context, obj *Review) (*User, error) {
	// Tips don't returns &User{ID: obj.AuthorID} simply, it caused username: null

	for _, user := range r.Usernames {
		if user.ID == obj.AuthorID {
			return user, nil
		}
	}
	return nil, nil
}

func (r *reviewResolver) Product(ctx context.Context, obj *Review) (*Product, error) {
	for _, product := range r.Products {
		if product.Upc == obj.ProductUPC {
			return product, nil
		}
	}
	return nil, nil
}

type userResolver struct{ *resolver }

func (r *userResolver) Reviews(ctx context.Context, obj *User) ([]*Review, error) {
	var resp []*Review
	for _, review := range r.resolver.Reviews {
		if review.AuthorID == obj.ID {
			resp = append(resp, review)
		}
	}

	return resp, nil
}
