package config

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
)

type ServerConfig struct {
	Addr string
	CORS cors.Config
}

func GetEnvServerConfig() *ServerConfig {
	host, port := "localhost", getEnvOrDefault("SERVER_PORT", "8080")

	if os.Getenv("GIN_MODE") == "release" {
		host = ""
	}

	return &ServerConfig{
		Addr: fmt.Sprintf("%s:%s", host, port),
		CORS: cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		},
	}
}
