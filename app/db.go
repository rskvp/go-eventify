package app

import (
	"log"

	"github.com/glebarez/sqlite"
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
	config *config.DBConfig
}

func NewDB(dbConfig *config.DBConfig) *DB {
	db := &DB{
		config: dbConfig,
	}

	switch dbConfig.Type {
	case config.DBTypeInMemory:
		db.openInMemoryDB()
	case config.DBTypeSqlite:
		db.openSqliteDB()
	default:
		db.openInMemoryDB()
	}

	db.migrate()

	return db
}

func (db *DB) openInMemoryDB() {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.DB = gormDB
}

func (db *DB) openSqliteDB() {
	gormDB, err := gorm.Open(sqlite.Open(db.config.Sqlite.File), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.DB = gormDB
}

func (db *DB) migrate() {
	db.AutoMigrate(
		&models.Flow{},
		&models.Event{},
		&models.Action{},
		&models.Execution{},
	)
}
