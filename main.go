package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"assalielmehdi/eventify/event"
	"assalielmehdi/eventify/flow"
	"assalielmehdi/eventify/handlers"
	"assalielmehdi/eventify/runners"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}
	addr := fmt.Sprintf(":%s", port)

	e := event.NewEvent()
	f := flow.NewHttpFlow(e)

	log.Println(f.Id)

	r := runners.NewHttpRunner()

	r.Register(f)

	rh := handlers.NewRunHandler(r)
	mux := http.NewServeMux()

	mux.Handle("/api/flow/run/", rh)

	log.Printf("Listening on %s\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err == nil {
		log.Fatal(err)
	}
}
