package explorer

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/explorer"
	"assalielmehdi/eventify/app/models"
)

func setup(t *testing.T) (*assert.Assertions, *app.DB, *explorer.ExplorerService, func()) {
	dbConfig := &config.DBConfig{
		Type: config.DBTypeSqlite,
		Sqlite: &config.DBSqliteConfig{
			File: "test.db",
		},
	}
	db := app.NewDB(dbConfig)
	service := explorer.NewExplorerService(db)

	return assert.New(t), db, service, func() {
		os.Remove(dbConfig.Sqlite.File)
	}
}

func TestGetTree(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	flows := []*models.Flow{
		{
			ID:   "1",
			Name: "Flow #1",
		},
		{
			ID:   "2",
			Name: "Flow #2",
		},
	}
	events := []*models.Event{
		{
			ID:     "1",
			Name:   "Event #1",
			FlowID: flows[0].ID,
		},
		{
			ID:     "2",
			Name:   "Event #2",
			FlowID: flows[0].ID,
		},
	}

	db.Create(flows)
	db.Create(events)

	tree, _ := service.GetTree()

	assert.Len(tree, 2)

	assert.Equal(tree[0].Key, fmt.Sprintf("/flows/%s", flows[0].ID))
	assert.Equal(tree[0].Label, flows[0].Name)
	assert.Len(tree[0].Children, 2)

	assert.Equal(tree[0].Children[0].Key, fmt.Sprintf("/flows/%s/events/%s", flows[0].ID, events[0].ID))
	assert.Equal(tree[0].Children[0].Label, events[0].Name)
	assert.Len(tree[0].Children[0].Children, 0)

	assert.Equal(tree[0].Children[1].Key, fmt.Sprintf("/flows/%s/events/%s", flows[0].ID, events[1].ID))
	assert.Equal(tree[0].Children[1].Label, events[1].Name)
	assert.Len(tree[0].Children[1].Children, 0)

	assert.Equal(tree[1].Key, fmt.Sprintf("/flows/%s", flows[1].ID))
	assert.Equal(tree[1].Label, flows[1].Name)
	assert.Len(tree[1].Children, 0)
}

func TestAddFlow(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	payload := &explorer.AddFlowRequest{
		Name: "flow",
	}

	service.AddFlow(payload)

	var record models.Flow
	db.Where("name = ?", payload.Name).Find(&record)

	assert.NotEmpty(record.ID)
	assert.Equal(record.Name, payload.Name)
}

func TestAddEvent(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	db.Create(&models.Flow{ID: "flow_id"})

	payload := &explorer.AddEventRequest{
		Name:   "event",
		FlowId: "flow_id",
	}

	service.AddEvent(payload)

	var record models.Event
	db.Where("name = ?", payload.Name).Find(&record)

	assert.NotEmpty(record.ID)
	assert.Equal(record.Name, payload.Name)
	assert.Equal(record.FlowID, payload.FlowId)
}
