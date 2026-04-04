package server

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
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
	// status code
	statusCode := http.StatusOK

	// handler label
	l := r.URL.Path
	
	// handler method
	m := r.Method

	// random delay
	maxRd := 800
	minRd := 500
	rd := rand.IntN(maxRd-minRd) + minRd

	// tail latency: introduce 5% slowest request latency
	if rand.IntN(100) < 5{
		rd = 2000 + rand.IntN(1000)
	}

	time.Sleep(time.Millisecond * time.Duration(rd))

	defer func() {
		sc := strconv.Itoa(statusCode)
		s.Metrics.RequestTotal.WithLabelValues(m, l, sc).Inc()
		s.Metrics.RequestLatency.WithLabelValues(m, l, sc).Observe((time.Since(t)).Seconds())
	}()

	// treat random number <= 5 as error
	if rand.IntN(100) <= 5 {
		statusCode = http.StatusInternalServerError
		log.Println("encounter error")
		writeResponse(w, statusCode, "encounter error")
		return
	}

	writeResponse(w, http.StatusOK, "ok")

	return
}
