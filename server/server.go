package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/KawaiiWafu/apitask/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// GORM Connection (SQLite)
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
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
		app.Get("/customer/new/:name/:email", createCustomer)
		app.Get("/product/new/:name/:price", createProduct)
		app.Listen(":3000")
	}
}

func createCustomer(c *fiber.Ctx) error {
	customer := data.Customer{
		Name:  c.Params("name"),
		Email: c.Params("email"),
	}
	result := db.Create(&customer)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	} else {
		return c.SendStatus(fiber.StatusCreated)
	}
}

func createProduct(c *fiber.Ctx) error {
	price, err := strconv.ParseFloat(c.Params("price"), 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product := data.Product{
		Name:     c.Params("name"),
		PriceCZK: price,
	}
	result := db.Create(&product)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	} else {
		return c.SendStatus(fiber.StatusCreated)
	}
}
