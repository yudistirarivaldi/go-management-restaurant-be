package main

import (
	"os"
	"maanagement-restaurant/database"
	"maanagement-restaurant/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.User
	

}