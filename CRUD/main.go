package main

import (
	"crud/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ExampleDB()
	
	router := gin.Default()

	
	//Listen in port 5000
	if err := router.Run(":5000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}