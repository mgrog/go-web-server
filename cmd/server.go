package main

import (
	gql_endpoint "go_server/router/gql"
	http_router "go_server/router/http"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8081"

	// Create Gin Engine
	r := gin.Default()

	// Setup http routes
	http_router.SetupRouter(r)

	// Setup gql routes
	gql_endpoint.SetupEndpointAndPlayground(r)

	log.Printf("Connect to http://localhost%s for GraphQL playground\n", port)
	log.Println(r.Run(port))

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")

}
