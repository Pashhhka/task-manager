package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Pashhhka/task-manager/internal/config"
	"github.com/Pashhhka/task-manager/internal/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresConnection(cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
