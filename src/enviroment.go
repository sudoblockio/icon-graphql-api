package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	MongodbUrl          string `envconfig:"ICON_GRAPHQL_API_MONGODB_URL" required:"true"`
	MongodbUser         string `envconfig:"ICON_GRAPHQL_API_MONGODB_USER" required:"true"`
	MongodbPass         string `envconfig:"ICON_GRAPHQL_API_MONGODB_PASS" required:"true"`
	Prefix              string `envconfig:"ICON_GRAPHQL_API_PREFIX" required:"false" default:""`
	Port                string `envconfig:"ICON_GRAPHQL_API_PORT" required:"false" default:"8000"`
	MetricsPort         string `envconfig:"ICON_GRAPHQL_API_METRICS_PORT" required:"false" default:"9400"`
	MetricsPollInterval int    `envconfig:"ICON_GRAPHQL_API_METRICS_POLL_INTERVAL" required:"false" default:"10"`
	NetworkName         string `envconfig:"ICON_GRAPHQL_API_NETWORK_NAME" required:"false" default:"mainnet"`
}

func getEnvironment() Environment {
	var env Environment
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("ERROR: envconfig - %s\n", err.Error())
	}
	return env
}
