package repositories

import (
	"time"

	"github.com/google/uuid"

	"assalielmehdi/eventify/app/models"
)

type FlowRepository struct {
	DB *DB
}

func NewFlowRepository(db *DB) *FlowRepository {
	return &FlowRepository{
		DB: db,
	}
}

func (repo *FlowRepository) GetAll() []*models.Flow {
	var records []*models.Flow

	repo.DB.Find(&records)

	return records
}

func (repo *FlowRepository) AddOne(record *models.Flow) *models.Flow {
	record.ID = uuid.NewString()
	record.Events = make([]*models.Event, 0)
	record.LastExecAt = time.Date(1996, time.February, 18, 0, 0, 0, 0, time.UTC)

	repo.DB.Create(record)

	return record
}
