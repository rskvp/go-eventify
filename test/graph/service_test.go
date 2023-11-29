package graph_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/graph"
	"assalielmehdi/eventify/app/models"
)

func setup() (*app.DB, graph.GraphService, func()) {
	dbConfig := &config.DBConfig{
		Type: config.DBTypeSqlite,
		Sqlite: &config.DBSqliteConfig{
			File: "test.db",
		},
	}
	db := app.NewDB(dbConfig)
	service := graph.NewGraphService(db)

	return db, *service, func() {
		os.Remove(dbConfig.Sqlite.File)
	}
}

func TestGetFlowGraph(t *testing.T) {
	db, service, teardown := setup()
	defer teardown()

	assert := assert.New(t)

	flowId := "1"
	eventIds := []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}
	events := []*models.Event{
		{
			ID:        eventIds[0],
			Name:      "event-1",
			IsInput:   true,
			IsOutput:  false,
			PositionX: 0,
			PositionY: 0,
			FlowID:    flowId,
		},
		{
			ID:        eventIds[1],
			Name:      "event-2",
			IsInput:   false,
			IsOutput:  false,
			PositionX: 50,
			PositionY: 50,
			PrevID:    eventIds[0],
			FlowID:    flowId,
		},
		{
			ID:        eventIds[2],
			Name:      "event-3",
			IsInput:   false,
			IsOutput:  true,
			PositionX: 100,
			PositionY: 100,
			PrevID:    eventIds[1],
			FlowID:    flowId,
		},
	}

	db.Create(&models.Flow{ID: flowId})
	db.Create(&events)

	flowGraph, _ := service.GetFlowGraph(flowId)

	assert.Len(flowGraph.Nodes, 3)

	assert.Equal(flowGraph.Nodes[0].Id, events[0].ID)
	assert.Equal(flowGraph.Nodes[0].Position, graph.FlowGraphNodePosition{X: 0, Y: 0})
	assert.Equal(flowGraph.Nodes[0].Data, graph.FlowGraphNodeData{Label: events[0].Name})
	assert.Equal(flowGraph.Nodes[0].TargetPosition, graph.NodeTargetPosition)
	assert.Equal(flowGraph.Nodes[0].SourcePosition, graph.NodeSourcePosition)
	assert.Equal(flowGraph.Nodes[0].Type, graph.NodeTypeInput)

	assert.Equal(flowGraph.Nodes[1].Id, events[1].ID)
	assert.Equal(flowGraph.Nodes[1].Position, graph.FlowGraphNodePosition{X: 50, Y: 50})
	assert.Equal(flowGraph.Nodes[1].Data, graph.FlowGraphNodeData{Label: events[1].Name})
	assert.Equal(flowGraph.Nodes[1].TargetPosition, graph.NodeTargetPosition)
	assert.Equal(flowGraph.Nodes[1].SourcePosition, graph.NodeSourcePosition)
	assert.Equal(flowGraph.Nodes[1].Type, graph.NodeTypeDefault)

	assert.Equal(flowGraph.Nodes[2].Id, events[2].ID)
	assert.Equal(flowGraph.Nodes[2].Position, graph.FlowGraphNodePosition{X: 100, Y: 100})
	assert.Equal(flowGraph.Nodes[2].Data, graph.FlowGraphNodeData{Label: events[2].Name})
	assert.Equal(flowGraph.Nodes[2].TargetPosition, graph.NodeTargetPosition)
	assert.Equal(flowGraph.Nodes[2].SourcePosition, graph.NodeSourcePosition)
	assert.Equal(flowGraph.Nodes[2].Type, graph.NodeTypeOutput)

	assert.Len(flowGraph.Edges, 2)

	assert.Equal(flowGraph.Edges[0].Id, fmt.Sprintf("%s->%s", events[0].ID, events[1].ID))
	assert.Equal(flowGraph.Edges[0].Source, events[0].ID)
	assert.Equal(flowGraph.Edges[0].Target, events[1].ID)
	assert.Equal(flowGraph.Edges[0].MarkerEnd, graph.FlowGraphEdgeMarker{Type: graph.EdgeMarkerEndType})

	assert.Equal(flowGraph.Edges[1].Id, fmt.Sprintf("%s->%s", events[1].ID, events[2].ID))
	assert.Equal(flowGraph.Edges[1].Source, events[1].ID)
	assert.Equal(flowGraph.Edges[1].Target, events[2].ID)
	assert.Equal(flowGraph.Edges[1].MarkerEnd, graph.FlowGraphEdgeMarker{Type: graph.EdgeMarkerEndType})
}

func TestUpdateEventPosition(t *testing.T) {
	db, service, teardown := setup()
	defer teardown()

	assert := assert.New(t)

	db.Create(&models.Flow{ID: "flow_id"})

	event := models.Event{
		ID:        uuid.NewString(),
		FlowID:    "flow_id",
		PositionX: 0,
		PositionY: 0,
	}
	payload := []*graph.FlowGraphNodePositionUpdate{
		{
			EventId: event.ID,
			NewPosition: &graph.FlowGraphNodePosition{
				X: 10.5,
				Y: 20.5,
			},
		},
	}

	db.Create(&event)

	service.UpdateEventsPositions(payload)

	db.Where("id = ?", event.ID).First(&event)

	assert.Equal(event.PositionX, payload[0].NewPosition.X)
	assert.Equal(event.PositionY, payload[0].NewPosition.Y)
}
