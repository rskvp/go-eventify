package config

import (
	"fmt"

	"github.com/gin-contrib/cors"
)

type ServerConfig struct {
	Addr string
	CORS cors.Config
}

func GetEnvServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr: fmt.Sprintf("localhost:%s", getEnvOrDefault("SERVER_PORT", "8080")),
		CORS: cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		},
	}
}
