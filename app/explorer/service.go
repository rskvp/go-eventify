package explorer

import (
	"fmt"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/models"
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
			Title:    flow.Name,
			Children: make([]*ExplorerTreeNode, 0),
			IsLeaf:   false,
		}

		for _, event := range flow.Events {
			flowNode.Children = append(flowNode.Children, &ExplorerTreeNode{
				Key:      fmt.Sprintf("/flows/%s/events/%s", flow.ID, event.ID),
				Title:    event.Name,
				Children: make([]*ExplorerTreeNode, 0),
				IsLeaf:   true,
			})
		}

		tree = append(tree, flowNode)
	}

	return tree, nil
}
