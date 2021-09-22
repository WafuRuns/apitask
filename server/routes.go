package main

import (
	"strconv"

	"github.com/WafuRuns/apitask/data"
	"github.com/gofiber/fiber/v2"
)

// POST /customer/new
// Creates new customer and returns them
func createCustomer(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong name or email format")
	}

	// Create customer
	customer := data.Customer{
		Name:  request.Name,
		Email: request.Email,
	}
	result := db.Create(&customer)
	if result.Error != nil {
		return serverError(c, "Customer creation failed")
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":   fiber.StatusCreated,
		"success":  true,
		"customer": customer,
	})
}

// POST /product/new
// Creates new product and returns it
func createProduct(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		Name     string  `json:"name"`
		PriceCZK float64 `json:"price"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong product name or price format")
	}

	// Create product
	product := data.Product{
		Name:     request.Name,
		PriceCZK: request.PriceCZK,
	}
	result := db.Create(&product)
	if result.Error != nil {
		return serverError(c, "Product creation failed")
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  fiber.StatusCreated,
		"success": true,
		"product": product,
	})
}

// POST /order/new
// Creates new order and returns it
func createOrder(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		CustomerID uint64 `json:"customer"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong customer ID format")
	}

	// Create order
	var customer data.Customer
	result := db.Where("customer_id = ?", request.CustomerID).First(&customer)
	if result.Error == nil {
		order := data.Order{
			Customer: customer,
			Items:    []data.OrderItem{},
		}
		result := db.Create(&order)
		if result.Error != nil {
			return serverError(c, "Order creation failed")
		}
		return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"status":  fiber.StatusCreated,
			"success": true,
			"order":   order,
		})
	}
	return clientError(c, "Customer does not exist")
}

// PUT /order/add
// Adds product into order and returns it
func addOrderItem(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		OrderID   uint64 `json:"orderid"`
		ProductID uint64 `json:"product"`
		Amount    uint64 `json:"amount"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong order/product ID or amount format")
	}

	// Add order item
	var order data.Order
	result := db.Where("order_id = ?", request.OrderID).First(&order)
	if result.Error == nil {
		var product data.Product
		result := db.Where("product_id = ?", request.ProductID).First(&product)
		if result.Error == nil {
			orderItem := data.OrderItem{
				Product:   product,
				ProductID: product.ProductID,
				Amount:    request.Amount,
			}
			db.Model(&order).Association("Items").Append(&orderItem)
			return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
				"status":    fiber.StatusCreated,
				"success":   true,
				"orderItem": orderItem,
			})
		}
		return clientError(c, "Product does not exist")
	}
	return clientError(c, "Order does not exist")
}

// GET /order/{orderid}
// Returns complete order information
func fetchOrder(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("orderid"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong order ID format")
	}
	var order data.Order
	result := db.Where("order_id = ?", orderID).First(&order)
	if result.Error == nil {
		db.Find(&order.Items, "order_id = ?", order.OrderID)
		db.Find(&order.Customer, "customer_id = ?", order.CustomerID)
		for i, orderItem := range order.Items {
			db.Find(&order.Items[i].Product, "product_id = ?", orderItem.ProductID)
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
			"order":   order,
		})
	}
	return clientError(c, "Order does not exist")
}

// DELETE /orderitem
// Deletes order item
func deleteOrderItem(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		ItemID uint64 `json:"itemid"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong order item ID format")
	}

	// Delete order item
	res := db.Where("order_item_id = ?", request.ItemID).Delete(data.OrderItem{})
	if res.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
		})
	}
	return clientError(c, "Order item does not exist")
}

// PUT /order/confirm
// Confirms order
func confirmOrder(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		OrderID uint64 `json:"orderid"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong order ID format")
	}

	// Confirm order
	res := db.Model(&data.Order{}).Where("order_id = ?", request.OrderID).Update("confirmed", true)
	if res.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
		})
	}
	return clientError(c, "Order does not exist")
}

// PUT /orderitem/amount
// Changes amount of items of order item
func changeOrderItemAmount(c *fiber.Ctx) error {
	// Process request body
	request := struct {
		ItemID uint64 `json:"itemid"`
		Amount uint64 `json:"amount"`
	}{}
	err := c.BodyParser(&request)
	if err != nil {
		return clientError(c, "Wrong order item ID or amount format")
	}

	// Change order item amount
	res := db.Model(&data.OrderItem{}).Where("order_item_id = ?", request.ItemID).Update("amount", request.Amount)
	if res.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
		})
	}
	return clientError(c, "Order item does not exist")
}
