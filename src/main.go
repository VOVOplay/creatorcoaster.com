package main

import (
	"log"
	"net/http"

	"github.com/VOVOplay/creatorcoaster.com/src/config"
	"github.com/VOVOplay/creatorcoaster.com/src/database"
	"github.com/VOVOplay/creatorcoaster.com/src/handlers"
)

func main() {
	config := config.GetConfig()

	router := configureRouter(config)

	database.GetStaffMemberList()

	err := http.ListenAndServe(config.Port, router)
	if err != nil {
		log.Fatal("Server crashed: ", err)
	}
}

func configureRouter(config config.Config) *http.ServeMux {
	router := http.NewServeMux()

	// Static
	fileServer := http.FileServer(http.Dir("static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	pageHandler := handlers.NewPageHandler()
	wikiHandler := handlers.NewWikiHandler()
	blogHandler := handlers.NewBlogHandler()
	adminHandler := handlers.NewAdminHandler(config.IsProduction)

	router.HandleFunc("GET /", pageHandler.HandleHome)
	router.HandleFunc("GET /about/", pageHandler.HandleAbout)
	router.HandleFunc("GET /about-this-website/", pageHandler.HandleAboutWebsite)
	router.HandleFunc("GET /privacy/", pageHandler.HandlePrivacy)
	router.HandleFunc(("GET /resources/"), pageHandler.HandleResources)

	router.HandleFunc("GET /wiki/", wikiHandler.HandleWiki)

	router.HandleFunc("GET /blog/", blogHandler.HandleBlog)

	router.HandleFunc("GET /admin/", adminHandler.HandleAdmin)
	router.HandleFunc("GET /auth/discord/callback", adminHandler.HandleCallback)
	router.HandleFunc("GET /auth/discord/logout", adminHandler.HandleAdminLogout)

	return router
}
