package config

import (
	"fmt"

	"github.com/gin-contrib/cors"
)

var serverConfig ServerConfig

type ServerConfig struct {
	Port string
	Addr string
	CORS cors.Config
}

func GetServerConfig() ServerConfig {
	return serverConfig
}

func initServerConfig() {
	serverConfig.Port = getEnvOrDefault("SERVER_PORT", "8080")
	serverConfig.Addr = fmt.Sprintf("localhost:%s", serverConfig.Port)
	serverConfig.CORS = cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}
}
