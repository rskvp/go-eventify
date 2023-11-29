package models

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/models"
)

func setupEventsDB(t *testing.T) (*assert.Assertions, *app.DB, func()) {
	dbConfig := &config.DBConfig{
		Type: config.DBTypeSqlite,
		Sqlite: &config.DBSqliteConfig{
			File: "test_events.db",
		},
	}
	db := app.NewDB(dbConfig)

	return assert.New(t), db, func() {
		os.Remove(dbConfig.Sqlite.File)
	}
}

func TestEventColumnsConstrains(t *testing.T) {
	assert, db, teardown := setupEventsDB(t)
	defer teardown()

	db.Create(&models.Flow{ID: "flow_id"})

	db.Create([]*models.Event{
		{ID: "duplicate_id", FlowID: "flow_id"},
		{ID: uuid.NewString(), Name: "duplicate_name", FlowID: "flow_id"},
	})

	records := []*models.Event{
		{ID: "duplicate_id"},
		{ID: uuid.NewString(), Name: "duplicate_name"},
	}

	for _, record := range records {
		err := db.Create(record).Error

		assert.NotNil(err)
	}
}

func TestEventFlowIdFK(t *testing.T) {
	assert, db, teardown := setupEventsDB(t)
	defer teardown()

	records := []*models.Event{
		{ID: uuid.NewString(), Name: "name_1"},
		{ID: uuid.NewString(), Name: "name_2", FlowID: "not_existing"},
	}

	for _, record := range records {
		err := db.Create(record).Error

		assert.NotNil(err)
	}
}

func TestEventExecutionColumnsEventIdFK(t *testing.T) {
	assert, db, teardown := setupEventsDB(t)
	defer teardown()

	records := []*models.EventExecution{
		{ID: uuid.NewString()},
		{ID: uuid.NewString(), EventID: "not_existing"},
	}

	for _, record := range records {
		err := db.Create(record).Error

		assert.NotNil(err)
	}
}
