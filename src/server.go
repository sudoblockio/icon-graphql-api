package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"pranavt61/icon-graphql-api/graph"
	"pranavt61/icon-graphql-api/graph/generated"
	"pranavt61/icon-graphql-api/metrics"
	"pranavt61/icon-graphql-api/mongodb"
)

const defaultPort = "8000"

func main() {

	// Get enviroment variables
	env := getEnvironment()

	// Start prometheus client
	go metrics.StartPrometheusHttpServer(env.Port, env.Prefix, env.NetworkName, env.MetricsPort, env.MetricsPollInterval)

	err := mongodb.ConnectClient(env.MongodbUrl, env.MongodbUser, env.MongodbPass)
	if err != nil {
		log.Printf("Error connecting mongo client: %s", err.Error())
		os.Exit(1)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle(env.Prefix+"/", playground.Handler("GraphQL playground", env.Prefix+"/query"))
	http.Handle(env.Prefix+"/query", srv)

	log.Printf("connect to http://localhost:%s%s/ for GraphQL playground", env.Port, env.Prefix)
	log.Fatal(http.ListenAndServe(":"+env.Port, nil))

}
