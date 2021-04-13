package metrics

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/machinebox/graphql"
)

var Metrics map[string]prometheus.Gauge

func StartPrometheusHttpServer(port string, prefix string, network_name string, metrics_port string, metrics_poll_interval int) {
	go metricsLoop(port, prefix, metrics_poll_interval)

	Metrics = make(map[string]prometheus.Gauge)

	// Create gauges
	Metrics["latest_block_number"] = promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "latest_block_number",
		Help:        "latest block number in mongodb",
		ConstLabels: prometheus.Labels{"network_name": network_name},
	})

	// Start server
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+metrics_port, nil)
}

func metricsLoop(port string, prefix string, metrics_poll_interval int) {
	client := graphql.NewClient("http://localhost:" + port + prefix + "/query")

	for {
		time.Sleep(time.Duration(metrics_poll_interval) * time.Second)

		// Check latest block number
		req := graphql.NewRequest(`
			query {
				blocks(skip: 0, limit: 1){
					number
				}
			}	
		`)
		var respData interface{}
		err := client.Run(context.Background(), req, &respData)
		if err != nil {
			log.Println("ERROR: ", err.Error())

			Metrics["latest_block_number"].Set(0)
		} else {
			block_number := respData.(map[string]interface{})["blocks"].([]interface{})[0].(map[string]interface{})["number"].(float64)

			Metrics["latest_block_number"].Set(block_number)
		}
	}
}
