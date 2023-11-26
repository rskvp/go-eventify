package config

type DBSqliteConfig struct {
	File string
}

const (
	DBTypeInMemory = "inmem"
	DBTypeSqlite   = "sqlite"
)

type DBConfig struct {
	Type   string
	Sqlite *DBSqliteConfig
}

func GetEnvDBConfig() *DBConfig {
	return &DBConfig{
		Type: getEnvOrDefault("DB_TYPE", DBTypeInMemory),
		Sqlite: &DBSqliteConfig{
			File: getEnvOrDefault("DB_SQLITE_FILE", "eventify.db"),
		},
	}
}
