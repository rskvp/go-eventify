package config

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
)

const (
	ServerModeDev  = 1
	ServerModeProd = 2
)

type ServerConfig struct {
	Addr string
	CORS cors.Config
	Mode int
}

func GetEnvServerConfig() *ServerConfig {
	mode := ServerModeDev
	if os.Getenv("GIN_MODE") == "release" {
		mode = ServerModeProd
	}

	host, port := "localhost", getEnvOrDefault("SERVER_PORT", "8080")
	if mode == ServerModeProd {
		host = ""
	}

	return &ServerConfig{
		Addr: fmt.Sprintf("%s:%s", host, port),
		CORS: cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		},
		Mode: mode,
	}
}
