package main

import (
	"fmt"
	"os"

	"github.com/KawaiiWafu/apitask/data"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
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
	if args[0] == "init" || args[0] == "start" {
		db.AutoMigrate(&data.Customer{})
		db.AutoMigrate(&data.Product{})
		db.AutoMigrate(&data.Order{})
		db.AutoMigrate(&data.OrderItem{})
	}

	// Start Fiber server
	if args[0] == "start" {
		// Start sending reminders on midnight
		// cron library runs the function as goroutine
		// Should use real crontab in production
		c := cron.New()
		c.AddFunc("@midnight", sendReminders)
		c.Start()

		// Configure routes and start Fiber
		// Larger apps should use Prefork
		app := fiber.New()
		app.Post("/customer/new", createCustomer)
		app.Post("/product/new", createProduct)
		app.Post("/order/new", createOrder)
		app.Put("/order/add", addOrderItem)
		app.Put("/order/confirm", confirmOrder)
		app.Get("/order/:orderid", fetchOrder)
		app.Delete("/orderitem", deleteOrderItem)
		app.Put("/orderitem/amount", changeOrderItemAmount)
		app.Listen(":3000")
	}
}
