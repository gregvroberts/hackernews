package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gregvroberts/hackernews/app/domain/repository/link"
	"github.com/gregvroberts/hackernews/app/domain/repository/user"
	"github.com/gregvroberts/hackernews/app/generated"
	"github.com/gregvroberts/hackernews/app/infrastructure/db"
	"github.com/gregvroberts/hackernews/app/infrastructure/persistence"
	"github.com/gregvroberts/hackernews/app/infrastructure/util"
	"github.com/gregvroberts/hackernews/app/interfaces"
	"github.com/joho/godotenv"
)

func init() {
	// Loads the values fromt he .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port := "8080"
	config, err := util.LoadConfig(".", "dev")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := db.OpenSqlxDB(&config)
	if err != nil {
		log.Fatal(err)
	}

	var linkService link.LinkService
	var userService user.UserService

	linkService = persistence.NewLink(conn)
	userService = persistence.NewUser(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		LinkService: linkService,
		UserService: userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
