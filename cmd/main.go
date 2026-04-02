package main

import (
	"log"

	"statistics-for-observability/server"
	"statistics-for-observability/utils"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	reg := prometheus.NewRegistry()
	metrics, err := utils.NewMetrics(reg)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(metrics)
	srv.RunServer(reg)
}
