package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/machinebox/graphql"
)

var Metrics map[string]prometheus.Counter

func StartPrometheusHttpServer(port string, network_name string) {
	Metrics = make(map[string]prometheus.Counter)

	// Create gauges
	Metrics["latest_block_number"] = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "websockets_bytes_written",
		Help:        "number of bytes sent over through websockets",
		ConstLabels: prometheus.Labels{"network_name": network_name},
	})

	// Start server
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}

func metricsLoop(port string, prefix string, poll_interval int) {
	client := graphql.NewClient("localhost:" + port + prefix + "/query")

	for {
		time.Sleep(poll_interval * time.Second)

		// Check latest block number
		req := graphql.NewRequest(`
				query ($key: String!) {
						items (id:$key) {
								field1
								field2
								field3
						}
				}
		`)

	}
}
