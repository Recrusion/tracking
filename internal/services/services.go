package services

import "tracking/internal/database"

type ServiceTracking struct {
	service *database.TrackingDatabase
}

func NewServiceTracking(service *database.TrackingDatabase) *ServiceTracking {
	s := &ServiceTracking{
		service: service,
	}
	return s
}
