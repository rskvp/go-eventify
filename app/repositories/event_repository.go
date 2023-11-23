package repositories

import (
	"github.com/google/uuid"

	"assalielmehdi/eventify/app/models"
)

type EventRepository struct {
	DB *DB
}

func NewEventRepository(db *DB) *EventRepository {
	return &EventRepository{
		DB: db,
	}
}

func (repo *EventRepository) AddOne(payload *models.Event) *models.Event {
	payload.ID = uuid.NewString()

	repo.DB.Create(payload)

	return payload
}
