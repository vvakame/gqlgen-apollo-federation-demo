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
		users: []*User{
			{
				ID:       "1",
				Name:     ptrString("Ada Lovelace"),
				Username: ptrString("@ada"),
			},
			{
				ID:       "2",
				Name:     ptrString("Alan"),
				Username: ptrString("@complete"),
			},
		},
	}
}

type resolver struct {
	users []*User
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

type queryResolver struct{ *resolver }

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	return r.users[0], nil
}

type entityResolver struct{ *resolver }

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}
