package repositories

import "assalielmehdi/eventify/app/models"

type FlowRepository struct {
	DB *DB
}

func NewFlowRepository(db *DB) *FlowRepository {
	return &FlowRepository{
		DB: db,
	}
}

func (repo *FlowRepository) GetAll() []*models.Flow {
	return []*models.Flow{}
}
