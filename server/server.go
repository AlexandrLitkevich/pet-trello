package main

import (
	"github.com/AlexandrLitkevich/pet-trello/graph/services"
	"log"
	"net/http"
	"os"

	resolvers "github.com/AlexandrLitkevich/pet-trello/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	generated "github.com/AlexandrLitkevich/pet-trello/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userServices := services.NewUserService()

	resolver := resolvers.Resolver{
		UserService: userServices,
	}

	var srv = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
