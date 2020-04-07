package svc

import (
	"net/http"

	"github.com/owncloud/ocis-settings/pkg/metrics"
)

// NewInstrument returns a service that instruments metrics.
func NewInstrument(next Service, metrics *metrics.Metrics) Service {
	return instrument{
		next:    next,
		metrics: metrics,
	}
}

type instrument struct {
	next    Service
	metrics *metrics.Metrics
}

// ServeHTTP implements the Service interface.
func (i instrument) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.next.ServeHTTP(w, r)
}

// Dummy implements the Service interface.
func (i instrument) Dummy(w http.ResponseWriter, r *http.Request) {
	i.next.Dummy(w, r)
}