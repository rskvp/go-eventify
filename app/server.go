package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/config"
)

type ServerResponseSuccess struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type ServerEmptyResponseSuccess struct {
	Type string `json:"type"`
}

type ServerResponseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewServerResponseSuccess(payload any) *ServerResponseSuccess {
	return &ServerResponseSuccess{
		Type:    "OK",
		Payload: payload,
	}
}

func NewServerEmptyResponseSuccess() *ServerEmptyResponseSuccess {
	return &ServerEmptyResponseSuccess{
		Type: "OK",
	}
}

func NewServerResponseError(err error) *ServerResponseError {
	return &ServerResponseError{
		Type:    "NOK",
		Message: err.Error(),
	}
}

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

	if serverConfig.Mode == config.ServerModeProd {
		server.handleStatic()
	}

	return server
}

func (server *Server) Register(router Router) {
	router.Register(server.Router)
}

func (server *Server) Serve() {
	server.Router.Run(server.config.Addr)
}

func (server *Server) handleStatic() {
	server.Router.Static("/static/", "static/")
	server.Router.LoadHTMLGlob("templates/*")
	server.Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
