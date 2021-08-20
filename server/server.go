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
		app.Get("/orderitem/:itemid/delete", deleteOrderItem)
		app.Get("/orderitem/:itemid/amount/:amount", changeOrderItemAmount)
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
	}
	return c.SendString(strconv.Itoa(int(customer.CustomerID)))
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
	}
	return c.SendString(strconv.Itoa(int(product.ProductID)))
}

func createOrder(c *fiber.Ctx) error {
	customerID, err := strconv.ParseInt(c.Params("customer"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var customer data.Customer
	result := db.Where("customer_id = ?", customerID).First(&customer)
	if result.Error == nil {
		order := data.Order{
			Customer: customer,
			Items:    []data.OrderItem{},
		}
		result := db.Create(&order)
		if result.Error != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.SendString(strconv.Itoa(int(order.OrderID)))
	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func addOrderItem(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("orderid"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var order data.Order
	result := db.Where("order_id = ?", orderID).First(&order)
	if result.Error == nil {
		productID, err := strconv.ParseInt(c.Params("product"), 10, 64)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		var product data.Product
		result := db.Where("product_id = ?", productID).First(&product)
		if result.Error == nil {
			amount, err := strconv.ParseInt(c.Params("amount"), 10, 64)
			if err != nil {
				return c.SendStatus(fiber.StatusBadRequest)
			}
			orderItem := data.OrderItem{
				Product:   product,
				ProductID: product.ProductID,
				Amount:    uint64(amount),
			}
			db.Model(&order).Association("Items").Append(&orderItem)
			// Getting order items
			var items []data.OrderItem
			db.Find(&items, "order_id = ?", order.OrderID)
			fmt.Println(items)
			return c.SendStatus(fiber.StatusOK)
		}
	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func deleteOrderItem(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadRequest)
}

func changeOrderItemAmount(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadRequest)
}
