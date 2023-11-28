package editor

import (
	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/models"
)

type EditorService struct {
	db *app.DB
}

func NewEditorService(db *app.DB) *EditorService {
	return &EditorService{
		db: db,
	}
}

func (service *EditorService) GetFlowById(id string) (*models.Flow, error) {
	var record models.Flow

	if err := service.db.Find(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (service *EditorService) UpdateFlowById(payload *models.Flow) error {
	return service.db.Save(payload).Error
}

func (service *EditorService) DeleteFlowById(id string) error {
	return service.db.Delete(&models.Flow{}, "id = ?", id).Error
}

func (service *EditorService) GetEventById(id string) (*models.Event, error) {
	var record models.Event

	if err := service.db.Find(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (service *EditorService) UpdateEventById(payload *models.Event) error {
	return service.db.Save(payload).Error
}

func (service *EditorService) DeleteEventById(id string) error {
	return service.db.Delete(&models.Event{}, "id = ?", id).Error
}
