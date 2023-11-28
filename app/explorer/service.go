package explorer

import (
	"fmt"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/models"

	"github.com/google/uuid"
)

type ExplorerService struct {
	db *app.DB
}

func NewExplorerService(db *app.DB) *ExplorerService {
	return &ExplorerService{
		db: db,
	}
}

func (service *ExplorerService) GetTree() (ExplorerTree, error) {
	var flows []*models.Flow
	err := service.db.Preload("Events").Find(&flows).Error
	if err != nil {
		return nil, err
	}

	tree := make(ExplorerTree, 0)

	for _, flow := range flows {
		flowNode := &ExplorerTreeNode{
			Key:      fmt.Sprintf("/flows/%s", flow.ID),
			Label:    flow.Name,
			Children: make([]*ExplorerTreeNode, 0),
			Type:     flow.Type,
		}

		for _, event := range flow.Events {
			flowNode.Children = append(flowNode.Children, &ExplorerTreeNode{
				Key:      fmt.Sprintf("/flows/%s/events/%s", flow.ID, event.ID),
				Label:    event.Name,
				Children: make([]*ExplorerTreeNode, 0),
			})
		}

		tree = append(tree, flowNode)
	}

	return tree, nil
}

func (service *ExplorerService) AddFlow(payload *AddFlowRequest) (*models.Flow, error) {
	record := &models.Flow{
		ID:   uuid.NewString(),
		Name: payload.Name,
		Type: models.FlowTypeDefault,
	}

	err := service.db.Create(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (service *ExplorerService) AddEvent(payload *AddEventRequest) (*models.Event, error) {
	record := &models.Event{
		ID:     uuid.NewString(),
		Name:   payload.Name,
		FlowID: payload.FlowId,
	}

	err := service.db.Create(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}
