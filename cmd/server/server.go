package main

import (
	"github.com/danyukod/chave-pix-bff/internal/application/commands"
	"github.com/danyukod/chave-pix-bff/internal/application/queries"
	"github.com/danyukod/chave-pix-bff/internal/infrastructure/client/pix-key"
	graph2 "github.com/danyukod/chave-pix-bff/internal/ui/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8083"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gateway := pix_key.NewPixKeyClient()
	findPixKeyByKeyUsecase := queries.NewFindPixKeyByKeyQuery(gateway)
	findPixKeyUsecase := queries.NewFindPixKeyQuery(gateway)
	createPixKeyUsecase := commands.NewCreatePixKeyService(gateway)
	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(graph2.Config{Resolvers: &graph2.Resolver{
		FindPixKeyUsecase:      findPixKeyUsecase,
		FindPixKeyByKeyUsecase: findPixKeyByKeyUsecase,
		CreatePixKeyUsecase:    createPixKeyUsecase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
