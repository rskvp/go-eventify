package services

import (
	"assalielmehdi/eventify/app/models"
	"assalielmehdi/eventify/app/repositories"
)

type EventService struct {
	EventRepository *repositories.EventRepository
}

func NewEventService(eventRepository *repositories.EventRepository) *EventService {
	return &EventService{
		EventRepository: eventRepository,
	}
}

func (service *EventService) AddOne(payload *models.Event) *models.Event {
	return service.EventRepository.AddOne(payload)
}
