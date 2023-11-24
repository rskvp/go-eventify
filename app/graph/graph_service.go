package graph

import (
	"assalielmehdi/eventify/app/models"
	"assalielmehdi/eventify/app/repositories"
	"fmt"
	"strconv"
	"strings"
)

type GraphService struct {
	db *repositories.DB
}

func NewGraphService(db *repositories.DB) *GraphService {
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
			Position:       extractNodePosition(event),
			Data:           FlowGraphNodeData{event.Name},
			TargetPosition: NodeTargetPosition,
			SourcePosition: NodeSourcePosition,
			Type:           extractNodeType(event),
		})

		if event.PrevEventID != "" {
			graph.Edges = append(graph.Edges, &FlowGraphEdges{
				Id:        fmt.Sprintf("%s->%s", event.PrevEventID, event.ID),
				Source:    event.PrevEventID,
				Target:    event.ID,
				MarkerEnd: FlowGraphEdgeMarker{EdgeMarkerEndType},
			})
		}
	}

	return graph, nil
}

func (service *GraphService) UpdateEventPosition(eventId string, payload *FlowGraphNodePosition) (*FlowGraphNodePosition, error) {
	query := `
		UPDATE events
		SET position = ?
		WHERE id = ?;
	`
	newPosition := fmt.Sprintf("%d|%d", payload.X, payload.Y)

	err := service.db.Exec(query, newPosition, eventId).Error
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func extractNodePosition(event *models.Event) FlowGraphNodePosition {
	tokens := strings.Split(event.Position, "|")

	if len(tokens) != 2 {
		return FlowGraphNodePosition{
			X: 0,
			Y: 0,
		}
	}

	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])

	return FlowGraphNodePosition{
		X: x,
		Y: y,
	}
}

func extractNodeType(event *models.Event) string {
	if event.IsInput {
		return NodeTypeInput
	}

	if event.IsOutput {
		return NodeTypeOutput
	}

	return ""
}
