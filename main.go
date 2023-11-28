package main

import (
	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/editor"
	"assalielmehdi/eventify/app/explorer"
	"assalielmehdi/eventify/app/graph"
)

func main() {
	config.Init()
	serverConfig := config.GetEnvServerConfig()
	dbConfig := config.GetEnvDBConfig()

	server := app.NewServer(serverConfig)
	db := app.NewDB(dbConfig)

	explorerService := explorer.NewExplorerService(db)
	explorerHandler := explorer.NewExplorerHandler(explorerService)
	exploterRouter := explorer.NewExplorerRouter(explorerHandler)
	server.Register(exploterRouter)

	graphService := graph.NewGraphService(db)
	graphHandler := graph.NewGraphHandler(graphService)
	graphRouter := graph.NewGraphRouter(graphHandler)
	server.Register(graphRouter)

	editorService := editor.NewEditorService(db)
	editorHandler := editor.NewEditorHandler(editorService)
	editorRouter := editor.NewEditorRouter(editorHandler)
	server.Register(editorRouter)

	server.Serve()
}
