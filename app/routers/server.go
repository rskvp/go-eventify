package routers

import (
	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/config"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (server *Server) Serve() {
	server.Router.Run(config.GetServerConfig().Addr)
}
