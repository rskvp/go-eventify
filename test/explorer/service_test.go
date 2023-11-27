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

func setup() (*app.DB, *explorer.ExplorerService, func()) {
	dbConfig := &config.DBConfig{
		Type: config.DBTypeSqlite,
		Sqlite: &config.DBSqliteConfig{
			File: "test.db",
		},
	}
	db := app.NewDB(dbConfig)
	service := explorer.NewExplorerService(db)

	return db, service, func() {
		os.Remove(dbConfig.Sqlite.File)
	}
}

func TestGetTree(t *testing.T) {
	db, service, teardown := setup()
	defer teardown()

	assert := assert.New(t)

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
	assert.Equal(tree[0].Title, flows[0].Name)
	assert.False(tree[0].IsLeaf)
	assert.Len(tree[0].Children, 2)

	assert.Equal(tree[0].Children[0].Key, fmt.Sprintf("/flows/%s/events/%s", flows[0].ID, events[0].ID))
	assert.Equal(tree[0].Children[0].Title, events[0].Name)
	assert.True(tree[0].Children[0].IsLeaf)
	assert.Len(tree[0].Children[0].Children, 0)

	assert.Equal(tree[0].Children[1].Key, fmt.Sprintf("/flows/%s/events/%s", flows[0].ID, events[1].ID))
	assert.Equal(tree[0].Children[1].Title, events[1].Name)
	assert.True(tree[0].Children[1].IsLeaf)
	assert.Len(tree[0].Children[1].Children, 0)

	assert.Equal(tree[1].Key, fmt.Sprintf("/flows/%s", flows[1].ID))
	assert.Equal(tree[1].Title, flows[1].Name)
	assert.False(tree[1].IsLeaf)
	assert.Len(tree[1].Children, 0)

}
