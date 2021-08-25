package main

import (
	"fmt"
	"os"
	"time"

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
		db.AutoMigrate(&data.Order{})
		db.AutoMigrate(&data.OrderItem{})
		return
	}

	// Start Fiber server
	if args[0] == "start" {
		app := fiber.New()
		app.Get("/customer/new/:name/:email", createCustomer)
		app.Get("/product/new/:name/:price", createProduct)
		app.Get("/order/new/:customer", createOrder)
		app.Get("/order/:orderid/add/:product/:amount", addOrderItem)
		app.Get("/order/:orderid/confirm", confirmOrder)
		app.Get("/order/:orderid", fetchOrder)
		app.Get("/orderitem/:itemid/delete", deleteOrderItem)
		app.Get("/orderitem/:itemid/amount/:amount", changeOrderItemAmount)
		app.Get("/emails", sendReminders) // Delete later
		app.Listen(":3000")
	}
}

// SHOULDN'T BE API REQUEST!
// Implement with cron
func sendReminders(c *fiber.Ctx) error {
	now := time.Now()
	// lastWeek := now.AddDate(0, 0, -7)
	lastWeek := now.AddDate(0, 0, 0)
	var orders []data.Order
	res := db.Where("confirmed = ? AND reminded = ? AND created_at <= ?", false, false, lastWeek).Find(&orders)
	if res.RowsAffected > 0 {
		tx := db.Begin()
		for _, order := range orders {
			// Send email
			fmt.Println(order)
			tx.Model(&order).Update("reminded", true)
		}
		tx.Commit()
	}
	return c.SendStatus(fiber.StatusOK)
}
