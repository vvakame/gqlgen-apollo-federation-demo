// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package demo

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (ec *executionContext) __resolve__service(ctx context.Context) (introspection.Service, error) {
	if ec.DisableIntrospection {
		return introspection.Service{}, errors.New("federated introspection disabled")
	}
	return introspection.Service{
		SDL: `type Mutation {
	createTodo(input: NewTodo!): Todo!
}
input NewTodo {
	text: String!
	userId: String!
}
type Query {
	todos: [Todo!]!
}
type Todo @key(fields: "id") {
	id: ID!
	text: String!
	done: Boolean!
	user: User!
}
type User @key(fields: "id") {
	id: ID!
	name: String!
}
`,
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) ([]_Entity, error) {
	list := []_Entity{}
	for _, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			return nil, errors.New("__typename must be an existing string")
		}
		switch typeName {

		case "Todo":
			id, ok := rep["id"].(string)
			if !ok {
				return nil, errors.New("opsies")
			}
			resp, err := ec.resolvers.Entity().FindTodoByID(ctx, id)
			if err != nil {
				return nil, err
			}

			list = append(list, resp)

		case "User":
			id, ok := rep["id"].(string)
			if !ok {
				return nil, errors.New("opsies")
			}
			resp, err := ec.resolvers.Entity().FindUserByID(ctx, id)
			if err != nil {
				return nil, err
			}

			list = append(list, resp)

		default:
			return nil, errors.New("unknown type: " + typeName)
		}
	}
	return list, nil
}