package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/config"
)

type Router interface {
	Register(*gin.Engine)
}

type Server struct {
	Router *gin.Engine
	config *config.ServerConfig
}

func NewServer(serverConfig *config.ServerConfig) *Server {
	server := &Server{
		Router: gin.Default(),
		config: serverConfig,
	}

	server.Router.Use(cors.New(server.config.CORS))

	return server
}

func (server *Server) Register(router Router) {
	router.Register(server.Router)
}

func (server *Server) Serve() {
	server.Router.Run(server.config.Addr)
}
