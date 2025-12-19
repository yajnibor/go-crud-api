package main

import (
	"context"
	"fmt"
	"go-crud-api/db" // Import the generated package
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build connection string from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Use a connection pool for better performance
	pool, _ := pgxpool.New(ctx, connStr)

	// Initialize sqlc queries
	queries := db.New(pool)

	r := gin.Default()

	// 1. Load the templates folder
	r.LoadHTMLGlob("templates/*")

	// 2. Define the Homepage route
	r.GET("/", func(c *gin.Context) {
		// You could query your DB here to get the real count
		bookCount, err := queries.CountBooks(ctx)
		if err != nil {
			log.Printf("Error counting books: %v", err)
			bookCount = 0
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gopher's Bookstore",
			"count": bookCount,
		})
	})

	r.GET("/books", func(c *gin.Context) {
		books, err := queries.ListBooks(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, books)
	})

	r.POST("/books", func(c *gin.Context) {
		var b db.CreateBookParams
		c.ShouldBindJSON(&b)

		newBook, _ := queries.CreateBook(ctx, b)
		c.JSON(http.StatusCreated, newBook)
	})

	r.Run(":8080")
}
