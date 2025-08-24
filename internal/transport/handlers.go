package transport

import "tracking/internal/services"

type HandlersTracking struct {
	endpoints services.Service
}

func NewHandlersTracking(endpoints services.Service) *HandlersTracking {
	h := &HandlersTracking{
		endpoints: endpoints,
	}
	return h
}
