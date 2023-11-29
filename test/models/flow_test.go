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

func setupFlowsDB(t *testing.T) (*assert.Assertions, *app.DB, func()) {
	dbConfig := &config.DBConfig{
		Type: config.DBTypeSqlite,
		Sqlite: &config.DBSqliteConfig{
			File: "test_flows.db",
		},
	}
	db := app.NewDB(dbConfig)

	return assert.New(t), db, func() {
		os.Remove(dbConfig.Sqlite.File)
	}
}

func TestFlowColumnsConstrains(t *testing.T) {
	assert, db, teardown := setupFlowsDB(t)
	defer teardown()

	db.Create([]*models.Flow{
		{ID: "duplicate_id"},
		{ID: uuid.NewString(), Name: "duplicate_name"},
	})

	records := []*models.Flow{
		{ID: "duplicate_id"},
		{ID: uuid.NewString(), Name: "duplicate_name"},
	}

	for _, record := range records {
		err := db.Create(record).Error

		assert.NotNil(err)
	}
}

func TestFlowExecutionColumnsEventIdFK(t *testing.T) {
	assert, db, teardown := setupEventsDB(t)
	defer teardown()

	records := []*models.FlowExecution{
		{ID: uuid.NewString()},
		{ID: uuid.NewString(), FlowID: "not_existing"},
	}

	for _, record := range records {
		err := db.Create(record).Error

		assert.NotNil(err)
	}
}
