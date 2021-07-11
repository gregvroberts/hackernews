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
)

var config util.Config

func init() {

}

func main() {
	tmpConfig, err := util.LoadConfig(".", "dev")
	// Loads the values from the .env into the system
	if err != nil {
		log.Fatal("No .env file found")
	}

	config = *tmpConfig

	var (
		defaultAppPort = "8080"
		appPort        = config.AppPort
	)

	// use default port if not specified in the .env file
	if appPort == "" {
		appPort = defaultAppPort
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
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))

}
