package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Pashhhka/task-manager/internal/config"
	"github.com/Pashhhka/task-manager/internal/database"
	"github.com/Pashhhka/task-manager/internal/handler"
	"github.com/Pashhhka/task-manager/internal/repository"
	"github.com/Pashhhka/task-manager/internal/service"
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
		c.JSON(200, gin.H{"message": "pong"})
	})

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
