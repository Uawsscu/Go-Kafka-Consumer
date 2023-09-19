package main

import (
	"fmt"
	"main/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	config.StartKafkaConsumer()
	config.ConnectDB()
	router := gin.Default()
	router.Use(cors.Default())
	initRoutes(router)
	router.Run(":8081")

}

func initRoutes(r *gin.Engine) {
	// r.GET("/anime/:id", anime.GetAnimeHandler)
	// r.POST("/anime/test", anime.animePost)
	// r.DELETE("/anime/delete", anime.animeDELETE)
	// r.PUT("/anime/update", anime.animeUpdate)

}
