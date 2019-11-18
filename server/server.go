package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	demo "github.com/vvakame/gqlgen-apollo-federation-demo"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(demo.NewExecutableSchema(demo.Config{Resolvers: &demo.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
