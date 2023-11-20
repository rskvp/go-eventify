package config

import (
	"fmt"
)

var serverConfig ServerConfig

type ServerConfig struct {
	Port string
	Addr string
}

func GetServerConfig() ServerConfig {
	return serverConfig
}

func initServerConfig() {
	serverConfig.Port = getEnvOrDefault("SERVER_PORT", "8080")
	serverConfig.Addr = fmt.Sprintf("localhost:%s", serverConfig.Port)
}
