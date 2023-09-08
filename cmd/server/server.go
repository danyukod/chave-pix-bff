package main

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands"
	"github.com/danyukod/chave-pix-bff/internal/infrastructure/client/pix-key"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/danyukod/chave-pix-bff/graph"
)

const defaultPort = "8083"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gateway := pix_key.NewPixKeyClient()
	findPixKeyByKeyUsecase := commands.NewFindPixKeyByKeyUsecase(gateway)
	findPixKeyUsecase := commands.NewFindPixKeyUsecase(gateway)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		FindPixKeyUsecase:      findPixKeyUsecase,
		FindPixKeyByKeyUsecase: findPixKeyByKeyUsecase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
