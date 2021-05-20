package server

import (
	"context"
	"log"
	"net/http"

	// "github.com/gin-gonic/autotls"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/Oasixer/ShopifyChallenge2021/config"
	"github.com/Oasixer/ShopifyChallenge2021/db"
	"github.com/Oasixer/ShopifyChallenge2021/handler"
	"github.com/Oasixer/ShopifyChallenge2021/resolvers"
	"github.com/Oasixer/ShopifyChallenge2021/schema"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func SetupRouter(db *db.DB) *gin.Engine{
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema_ptr := schema.NewSchema()
	s := graphql.MustParseSchema(*schema_ptr, &resolvers.Resolvers{DB: db}, opts...)

	context.Background()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(NoFlocMiddleware())

	//////////////////
	// serving single page application stuff
	//////////////////

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../frontend/public", true)))

	// if no route then serve the root file
	router.NoRoute(func(c *gin.Context) {
		c.File("../frontend/public/index.html")
	})

	// graphql routes and route handling
	// only enable graphql playground on dev mode
	if gin.Mode() == gin.DebugMode {
		router.GET("/graphql", gin.WrapH(handler.GraphiQL{}))
	}
	router.POST("/graphql", gin.WrapH(handler.Authenticate(&handler.GraphQL{Schema: s})))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.POST("/upload-file", handler.SaveFile)
	api.GET("/download-file", handler.GetFile)
	// serve and run

	return router
}

func Serve(db *db.DB) {
	router := SetupRouter(db)
	log.Fatal(router.Run(":" + config.CONFIG.Port))
}
