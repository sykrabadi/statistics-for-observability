package server

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

func writeResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	rsp, err := json.Marshal(map[string]string{
		"message": message,
	})
	if err != nil {
		log.Println(err)
	}
	w.Write(rsp)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	rn := rand.IntN(100)

	// status code
	sc := "200"

	// handler label
	l := r.URL.Path

	// treat random number <= 30 as error
	if rn <= 30 {
		sc = "500"
		s.Metrics.RequestTotal.WithLabelValues(l, sc).Inc()
		s.Metrics.RequestLatency.WithLabelValues(l, sc).Observe(float64(time.Since(t)))

		log.Println("encounter error")
		writeResponse(w, http.StatusInternalServerError, "encounter error")
		return
	}

	s.Metrics.RequestTotal.WithLabelValues(l, sc).Inc()
	s.Metrics.RequestLatency.WithLabelValues(l, sc).Observe(float64(time.Since(t)))

	writeResponse(w, http.StatusOK, "ok")
	return
}
