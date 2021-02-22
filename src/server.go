package main

import (
	"log"
	"net/http"
	"os"
	"pranavt61/icon-graphql-api/graph"
	"pranavt61/icon-graphql-api/graph/generated"
	"pranavt61/icon-graphql-api/mongodb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8000"

func main() {

	mongodb_url_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_URL")
	mongodb_user_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_USER")
	mongodb_pass_env := os.Getenv("ICON_GRAPHQL_API_MONGODB_PASS")
	prefix_env := os.Getenv("ICON_GRAPHQL_API_PREFIX")
	port_env := os.Getenv("ICON_GRAPHQL_API_PORT")

	if mongodb_url_env == "" {
		panic("ERROR: missing required enviroment variable: ICON_GRAPHQL_API_MONGODB_URL")
	}
	if mongodb_user_env == "" {
		panic("ERROR: missing required enviroment variable: ICON_GRAPHQL_API_MONGODB_USER")
	}
	if mongodb_pass_env == "" {
		panic("ERROR: missing required enviroment variable: ICON_GRAPHQL_API_MONGODB_PASS")
	}
	if port_env == "" {
		port_env = defaultPort
	}
	if port_env == "" {
		port_env = ""
	}

	mongodb.ConnectClient(mongodb_url_env, mongodb_user_env, mongodb_pass_env)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle(prefix_env+"/", playground.Handler("GraphQL playground", prefix_env+"/query"))
	http.Handle(prefix_env+"/query", srv)

	log.Printf("connect to http://localhost:%s%s/ for GraphQL playground", port_env, prefix_env)
	log.Fatal(http.ListenAndServe(":"+port_env, nil))

}
