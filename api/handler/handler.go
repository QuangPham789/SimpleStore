package handler

import (
	"context"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/repositories"
	"graphdemo/pkg/services"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

// Routes initialize the handlers for the router
func Routes() *chi.Mux {
	router := chi.NewRouter()
	queryHandler := graphQLHandler()
	router.Post("/query", repositories.DataloaderMiddleware(context.Background(), queryHandler))

	router.Get("/playground", playgroundQLHandler("/v1/query"))

	return router
}

func graphQLHandler() http.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &services.Resolver{}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	//endpoint argument must be same as graphql handler path

	playgroundHandler := playground.Handler("GraphQL", endpoint)

	return playgroundHandler
}
