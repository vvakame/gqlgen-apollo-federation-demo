//go:generate go run github.com/99designs/gqlgen

package demo

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var _ ResolverRoot = (*Resolver)(nil)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	panic("not implemented")
}

type entityResolver struct{ *Resolver }

func (e *entityResolver) FindTodoByID(ctx context.Context, id string) (*Todo, error) {
	panic("implement me")
}

func (e *entityResolver) FindUserByID(ctx context.Context, id string) (*User, error) {
	panic("implement me")
}
