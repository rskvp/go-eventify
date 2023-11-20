package config

var dbConfig DBConfig

type DBSqliteConfig struct {
	File string
}

type DBConfig struct {
	Sqlite DBSqliteConfig
}

func GetDBConfig() DBConfig {
	return dbConfig
}

func initDBConfig() {
	dbConfig = DBConfig{
		Sqlite: DBSqliteConfig{
			File: getEnvOrDefault("DB_SQLITE_FILE", "eventify.db"),
		},
	}
}
