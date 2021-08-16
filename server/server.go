package main

import (
	"fmt"
	"os"

	"github.com/KawaiiWafu/apitask/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// GORM Connection (SQLite)
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}

	args := os.Args[1:]
	// Show help if no arguments were given
	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("  server start\t\tStarts Fiber server")
		fmt.Println("  server init\tInitializes SQLite database")
		return
	}

	// Initialize database
	if args[0] == "init" {
		db.AutoMigrate(&data.Customer{})
		db.AutoMigrate(&data.Product{})
		db.AutoMigrate(&data.OrderItem{})
		db.AutoMigrate(&data.Order{})
		return
	}

	// Start Fiber server
	if args[0] == "start" {
		app := fiber.New()
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello")
		})
		app.Listen(":3000")
	}
}
