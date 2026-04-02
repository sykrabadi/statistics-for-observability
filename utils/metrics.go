package utils

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pkg/errors"
)

type Metrics struct{
	RequestTotal *prometheus.CounterVec
	RequestLatency *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) (*Metrics, error){
	m := &Metrics{
		RequestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "total_request",
				Help: "http total requests",
			},
			[]string{"handler", "status_code"},
		),
		RequestLatency: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "request_latency",
				Help: "http request latency in second",
			},
			[]string{"handler", "status_code"},
		),
	}

	err := reg.Register(m.RequestTotal)
	if err != nil {
		return nil, errors.Wrap(err, "fail register request total")
	}
	err = reg.Register(m.RequestLatency)	
	if err != nil {
		return nil, errors.Wrap(err, "fail register request latency")
	}
	
	return m, nil
}