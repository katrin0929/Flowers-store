package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katrin0929/Flowers-store/back/internal/handler"
	"github.com/katrin0929/Flowers-store/back/internal/model"
	"github.com/katrin0929/Flowers-store/back/internal/repository"
	"github.com/katrin0929/Flowers-store/back/internal/service"
	"github.com/katrin0929/Flowers-store/back/pkg/database"
	"log"
)

func main() {
	// Инициализация БД одной строкой
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to init DB:", err)
	}

	// Автомиграция
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Инициализация зависимостей
	authHandler := handler.NewAuthHandler(
		service.NewAuthService(
			repository.NewUserRepository(db),
		),
	)

	// Запуск сервера
	router := gin.Default()
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	//port := database.GetEnv("APP_PORT", "8080")
	log.Printf("Server running on :%s", "8080")
	if err := router.Run(":" + "8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
