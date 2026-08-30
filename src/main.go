package main

import (
	"log"
	"net/http"

	"github.com/VOVOplay/creatorcoaster.com/src/config"
	"github.com/VOVOplay/creatorcoaster.com/src/handlers"
)

func main() {
	config := config.GetConfig()

	router := configureRouter()

	err := http.ListenAndServe(config.Port, router)
	if err != nil {
		log.Fatal("Server crashed: ", err)
	}
}

func configureRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Static
	fileServer := http.FileServer(http.Dir("src/static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	pageHandler := handlers.NewPageHandler()
	router.HandleFunc("GET /", pageHandler.HandleHome)

	return router
}
