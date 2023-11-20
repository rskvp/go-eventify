package repositories

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/models"
)

const (
	DBTypeInMemory = 1
	DBTypeSqlite   = 2
)

type DB struct {
	*gorm.DB
	Type int
}

func NewDB(dbType int) *DB {
	db := &DB{
		Type: dbType,
	}

	switch dbType {
	case DBTypeInMemory:
		db.DB = openInMemoryDB()
	case DBTypeSqlite:
		db.DB = openSqliteDB()
	default:
		db.DB = openInMemoryDB()
	}

	db.migrate()

	return db
}

func openInMemoryDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func openSqliteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.GetDBConfig().Sqlite.File), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (db *DB) migrate() {
	db.AutoMigrate(
		&models.Flow{},
		&models.Event{},
		&models.Action{},
		&models.Execution{},
	)
}
