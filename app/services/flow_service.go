package services

import (
	"assalielmehdi/eventify/app/models"
	"assalielmehdi/eventify/app/repositories"
)

type FlowService struct {
	FlowRepository *repositories.FlowRepository
}

func NewFlowService(flowRepository *repositories.FlowRepository) *FlowService {
	return &FlowService{
		FlowRepository: flowRepository,
	}
}

func (sercice *FlowService) GetAll() []*models.Flow {
	return sercice.FlowRepository.GetAll()
}

func (service *FlowService) GetOneById(id string) (*models.Flow, error) {
	return service.FlowRepository.GetOneById(id)
}

func (service *FlowService) AddOne(payload *models.Flow) *models.Flow {
	return service.FlowRepository.AddOne(payload)
}

func (service *FlowService) UpdateOneById(payload *models.Flow) (*models.Flow, error) {
	return service.FlowRepository.UpdateOneById(payload)
}

func (service *FlowService) DeleteOneById(id string) (*models.Flow, error) {
	return service.FlowRepository.DeleteOneById(id)
}
