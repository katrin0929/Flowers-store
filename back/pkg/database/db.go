package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("POSTGRES_HOST", "postgres"),
		GetEnv("POSTGRES_PORT", "5432"),
		GetEnv("POSTGRES_USER", "postgres"),
		GetEnv("POSTGRES_PASSWORD", "postgres"),
		GetEnv("POSTGRES_DB", "postgres"),
	)

	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Connection attempt %d failed: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	return db, err
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
