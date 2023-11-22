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
				Name:     ptrString("Jad"),
				Username: ptrString("@jad"),
			},
			{
				ID:       "2",
				Name:     ptrString("Marouan"),
				Username: ptrString("@marouan"),
			},
			{
				ID:       "3",
				Name:     ptrString("Anas"),
				Username: ptrString("@anaselhajjaji"),
			},
			{
				ID:       "4",
				Name:     ptrString("Chaymae"),
				Username: ptrString("@chaymae"),
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

func (r *queryResolver) AllUsers(ctx context.Context) ([]*User, error) {
	return r.users, nil
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
