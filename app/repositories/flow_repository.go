package repositories

import (
	"errors"
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

func (repo *FlowRepository) GetOneById(id string) (*models.Flow, error) {
	var record *models.Flow

	repo.DB.First(&record, "id = ?", id)

	if record == nil {
		return nil, errors.New("not found")
	}

	return record, nil
}

func (repo *FlowRepository) AddOne(record *models.Flow) *models.Flow {
	record.ID = uuid.NewString()
	record.Events = make([]*models.Event, 0)
	record.LastExecAt = time.Date(1996, time.February, 18, 0, 0, 0, 0, time.UTC)

	repo.DB.Create(record)

	return record
}

func (repo *FlowRepository) UpdateOneById(record *models.Flow) (*models.Flow, error) {
	record, err := repo.GetOneById(record.ID)
	if err != nil {
		return nil, err
	}

	repo.DB.Save(&record)

	return record, nil
}

func (repo *FlowRepository) DeleteOneById(id string) (*models.Flow, error) {
	record, err := repo.GetOneById(id)
	if err != nil {
		return nil, err
	}

	repo.DB.Delete(&record)

	return record, nil
}
