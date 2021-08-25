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
	if args[0] == "init" {
		db.AutoMigrate(&data.Customer{})
		db.AutoMigrate(&data.Product{})
		db.AutoMigrate(&data.Order{})
		db.AutoMigrate(&data.OrderItem{})
		return
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
		app.Post("/customer/new/:name/:email", createCustomer)
		app.Post("/product/new/:name/:price", createProduct)
		app.Post("/order/new/:customer", createOrder)
		app.Put("/order/:orderid/add/:product/:amount", addOrderItem)
		app.Post("/order/:orderid/confirm", confirmOrder)
		app.Get("/order/:orderid", fetchOrder)
		app.Delete("/orderitem/:itemid/delete", deleteOrderItem)
		app.Put("/orderitem/:itemid/amount/:amount", changeOrderItemAmount)
		app.Listen(":3000")
	}
}
