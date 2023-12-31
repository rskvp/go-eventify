package graph

import (
	"fmt"

	"gorm.io/gorm"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/models"
)

type GraphService struct {
	db *app.DB
}

func NewGraphService(db *app.DB) *GraphService {
	return &GraphService{
		db: db,
	}
}

func (service *GraphService) GetFlowGraph(flowId string) (*FlowGraph, error) {
	var events []*models.Event

	err := service.db.Where("flow_id	= ?", flowId).Find(&events).Error
	if err != nil {
		return nil, err
	}

	graph := &FlowGraph{
		Nodes: make([]*FlowGraphNode, 0),
		Edges: make([]*FlowGraphEdges, 0),
	}

	for _, event := range events {
		graph.Nodes = append(graph.Nodes, &FlowGraphNode{
			Id:             event.ID,
			Position:       FlowGraphNodePosition{X: event.PositionX, Y: event.PositionY},
			Data:           FlowGraphNodeData{event.Name},
			TargetPosition: NodeTargetPosition,
			SourcePosition: NodeSourcePosition,
			Type:           extractNodeType(event),
		})

		if event.PrevID != "" {
			graph.Edges = append(graph.Edges, &FlowGraphEdges{
				Id:        fmt.Sprintf("%s->%s", event.PrevID, event.ID),
				Source:    event.PrevID,
				Target:    event.ID,
				MarkerEnd: FlowGraphEdgeMarker{EdgeMarkerEndType},
			})
		}
	}

	return graph, nil
}

func extractNodeType(event *models.Event) string {
	if event.IsInput {
		return NodeTypeInput
	}

	if event.IsOutput {
		return NodeTypeOutput
	}

	return NodeTypeDefault
}

func (service *GraphService) UpdateEventsPositions(updates []*FlowGraphNodePositionUpdate) error {
	return service.db.Transaction(func(tx *gorm.DB) error {
		for _, update := range updates {
			event := &models.Event{
				ID:        update.EventId,
				PositionX: update.NewPosition.X,
				PositionY: update.NewPosition.Y,
			}

			if err := service.db.Model(event).Select("PositionX", "PositionY").Updates(event).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
