package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ountyrio/gql-bookstore/internal/graph/generated"
	"github.com/ountyrio/gql-bookstore/internal/graph/resolvers"
	"github.com/ountyrio/gql-bookstore/internal/service"
	"github.com/ountyrio/gql-bookstore/internal/service/repository"
	pg "github.com/ountyrio/gql-bookstore/utils"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	db, err := pg.OpenDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	port := viper.GetString("server.port")

	// init service implementation
	authorRepository := repository.NewAuthorRepository(db)
	bookRepository := repository.NewBookRepository(db)
	genreRepository := repository.NewGenreRepository(db)

	authorService := service.NewAuthorService(authorRepository)
	bookService := service.NewBookService(authorRepository, bookRepository, genreRepository)
	genreService := service.NewGenreService(genreRepository)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: resolvers.NewResolver(authorService, bookService, genreService),
		}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
