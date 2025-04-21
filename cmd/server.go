package main

import (
	"fmt"
	"go_server/config"
	gql_endpoint "go_server/router/gql"
	http_router "go_server/router/http"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// @title Go Server API
// @version 1.0
// @description Example Gin GraphQL/REST server with kitchen sink of features

// @host localhost:8081
// @BasePath /v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	pg_cfg := config.PgConfigFromEnv().SslDisabled()

	db, err := sqlx.Connect("postgres", pg_cfg.GetDsn())
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	// Create a new HTTP client
	httpClient := &http.Client{}

	// Create Gin Engine
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Setup http routes
	http_router.SetupRouter(r)

	// Setup gql endpoint
	gql_endpoint.SetupEndpointAndPlayground(r, db, httpClient)

	log.Printf("Connect to http://localhost%s for GraphQL playground\n", port)
	log.Println(r.Run(port))

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// Handle graceful shutdown
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
