package editor

import (
	"os"
	"testing"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/editor"
	"assalielmehdi/eventify/app/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*assert.Assertions, *app.DB, *editor.EditorService, func()) {
	dbConfig := &config.DBConfig{
		Type: config.DBTypeSqlite,
		Sqlite: &config.DBSqliteConfig{
			File: "test.db",
		},
	}
	db := app.NewDB(dbConfig)
	service := editor.NewEditorService(db)

	return assert.New(t), db, service, func() {
		os.Remove(dbConfig.Sqlite.File)
	}
}

func TestGetFlowById(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	record := &models.Flow{
		ID:   uuid.NewString(),
		Name: "flow",
	}

	db.Create(record)

	existingRecord, _ := service.GetFlowById(record.ID)

	assert.Equal(record.ID, existingRecord.ID)
	assert.Equal(record.Name, existingRecord.Name)
}

func TestUpdateFlowById(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	record := &models.Flow{
		ID:   uuid.NewString(),
		Name: "flow",
	}

	db.Create(record)

	record.Name = "flow 1"

	service.UpdateFlowById(record)

	var existingRecord models.Flow
	db.Find(&existingRecord, "id = ?", record.ID)

	assert.Equal(record.ID, existingRecord.ID)
	assert.Equal(record.Name, existingRecord.Name)
}

func TestDeleteFlowById(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	record := &models.Flow{
		ID:   uuid.NewString(),
		Name: "flow",
	}

	db.Create(record)

	service.DeleteFlowById(record.ID)

	var existingRecord models.Flow
	db.Find(&existingRecord, "id = ?", record.ID)

	assert.NotEqual(record.ID, existingRecord.ID)
}

func TestGetEventById(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	record := &models.Event{
		ID:   uuid.NewString(),
		Name: "event",
	}

	db.Create(record)

	existingRecord, _ := service.GetEventById(record.ID)

	assert.Equal(record.ID, existingRecord.ID)
	assert.Equal(record.Name, existingRecord.Name)
}

func TestUpdateEventById(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	record := &models.Event{
		ID:   uuid.NewString(),
		Name: "event",
	}

	db.Create(record)

	record.Name = "event 1"

	service.UpdateEventById(record)

	var existingRecord models.Event
	db.Find(&existingRecord, "id = ?", record.ID)

	assert.Equal(record.ID, existingRecord.ID)
	assert.Equal(record.Name, existingRecord.Name)
}

func TestDeleteEventById(t *testing.T) {
	assert, db, service, teardown := setup(t)
	defer teardown()

	record := &models.Event{
		ID:   uuid.NewString(),
		Name: "event",
	}

	db.Create(record)

	service.DeleteEventById(record.ID)

	var existingRecord models.Event
	db.Find(&existingRecord, "id = ?", record.ID)

	assert.NotEqual(record.ID, existingRecord.ID)
}
