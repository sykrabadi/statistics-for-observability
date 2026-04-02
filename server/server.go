package server

import (
	"log"
	"net/http"

	"statistics-for-observability/utils"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	Metrics *utils.Metrics
}

func NewServer(metrics *utils.Metrics) *Server {
	return &Server{
		Metrics: metrics,
	}
}

func (s *Server) RunServer(reg *prometheus.Registry) {
	http.HandleFunc("/home", s.homeHandler)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	port := ":8083"
	log.Println("running http server at: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("fail running at: %v, with error: %v", port, err)
	}
}
