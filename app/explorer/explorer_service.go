package explorer

import (
	"assalielmehdi/eventify/app/repositories"
	"fmt"
)

type ExplorerService struct {
	db *repositories.DB
}

func NewExplorerService(db *repositories.DB) *ExplorerService {
	return &ExplorerService{
		db: db,
	}
}

func (service *ExplorerService) GetTree() (ExplorerTree, error) {
	query := `
		SELECT flows.id as flowId, flows.name as flowName, events.id as eventId, events.name as eventName
		FROM flows
		LEFT JOIN events
		ON events.flow_id = flows.id;
	`

	rows, err := service.db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tree := make(ExplorerTree, 0)
	var curFlow *ExplorerTreeNode

	for rows.Next() {
		var flowId, flowName, eventId, eventName string

		rows.Scan(&flowId, &flowName, &eventId, &eventName)

		if curFlow == nil || curFlow.Key != fmt.Sprintf("/flows/%s", flowId) {
			if curFlow != nil {
				tree = append(tree, curFlow)
			}

			curFlow = &ExplorerTreeNode{
				Key:      fmt.Sprintf("/flows/%s", flowId),
				Title:    flowName,
				Children: make([]*ExplorerTreeNode, 0),
				IsLeaf:   false,
			}
		}

		if eventId != "" {
			event := &ExplorerTreeNode{
				Key:      fmt.Sprintf("%s/events/%s", curFlow.Key, eventId),
				Title:    eventName,
				Children: make([]*ExplorerTreeNode, 0),
				IsLeaf:   true,
			}

			curFlow.Children = append(curFlow.Children, event)
		}
	}

	if curFlow != nil {
		tree = append(tree, curFlow)
	}

	return tree, nil
}
