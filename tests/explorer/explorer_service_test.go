package explorer_test

import (
	"assalielmehdi/eventify/app/explorer"
	"assalielmehdi/eventify/app/models"
	"assalielmehdi/eventify/app/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepare() (*repositories.DB, *explorer.ExplorerService) {
	db := repositories.NewDB(repositories.DBTypeInMemory)
	service := explorer.NewExplorerService(db)

	return db, service
}

func TestGetTree(t *testing.T) {
	assert := assert.New(t)
	db, service := prepare()

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
