package main

import (
	"fmt"
	"go/http/configs"
	"go/http/internal/auth"
	"go/http/internal/link"
	"go/http/pkg/db"
	"go/http/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}
	fmt.Println("Server is listening on port 8081")
	_ = server.ListenAndServe()
}
