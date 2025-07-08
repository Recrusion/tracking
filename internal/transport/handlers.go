package transport

import "tracking/internal/services"

type HandlersTracking struct {
	handlers *services.ServiceTracking
}

func NewHandlersTracking(handlers *services.ServiceTracking) *HandlersTracking {
	h := &HandlersTracking{
		handlers: handlers,
	}
	return h
}
