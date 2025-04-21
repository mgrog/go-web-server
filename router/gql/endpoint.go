package gql_endpoint

import (
	"context"
	"go_server/graph"
	"go_server/graph/dataloader"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	lru "github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/vektah/gqlparser/v2/ast"
)

func SetupEndpointAndPlayground(r *gin.Engine, db *sqlx.DB, httpClient *http.Client) {
	h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db, HttpClient: httpClient}}))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	h.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		return next(ctx)
	})

	r.Use(dataloader.Middleware(httpClient))

	r.GET("/", func(c *gin.Context) {
		h := playground.Handler("GraphQL playground", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}
