package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/config"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		Router: gin.Default(),
	}

	server.Router.Use(cors.New(config.GetServerConfig().CORS))

	return server
}

func (server *Server) Serve() {
	server.Router.Run(config.GetServerConfig().Addr)
}
