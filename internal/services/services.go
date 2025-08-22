package services

import "tracking/internal/storage"

type ServiceTracking struct {
	service *storage.TrackingDatabase
}

func NewServiceTracking(service *storage.TrackingDatabase) *ServiceTracking {
	s := &ServiceTracking{
		service: service,
	}
	return s
}
