package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sscarry2/ginapi/configs"
	"github.com/sscarry2/ginapi/routes"
)

func main() {
	
	router := setupRouter()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv(("PORT"))
	router.Run(":" + port)
	
}

func setupRouter() *gin.Engine {
	//load .env
	godotenv.Load(".env")

	//connect DB
	configs.ConnectDB()

	
	router := gin.Default()
	//cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		// MaxAge: 12 * time.Hour,
	  }))

	apiV1 := router.Group("/api/v1")

	routes.InitHomeRoute(apiV1)
	routes.InitUserRoute(apiV1)


	return router
}