package main

import (
	"strconv"

	"github.com/KawaiiWafu/apitask/data"
	"github.com/gofiber/fiber/v2"
)

// POST /customer/new/{name}/{email}
// Creates new customer and returns them
func createCustomer(c *fiber.Ctx) error {
	customer := data.Customer{
		Name:  c.Params("name"),
		Email: c.Params("email"),
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

// POST /product/new/{name}/{price}
// Creates new product and returns it
func createProduct(c *fiber.Ctx) error {
	price, err := strconv.ParseFloat(c.Params("price"), 64)
	if err != nil {
		return clientError(c, "Wrong product price format")
	}
	product := data.Product{
		Name:     c.Params("name"),
		PriceCZK: price,
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

// POST /order/new/{customer}
// Creates new order and returns it
func createOrder(c *fiber.Ctx) error {
	customerID, err := strconv.ParseInt(c.Params("customer"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong customer ID format")
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

// PUT /order/{orderid}/add/{product}/{amount}
// Adds product into order and returns it
func addOrderItem(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("orderid"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong order ID format")
	}
	var order data.Order
	result := db.Where("order_id = ?", orderID).First(&order)
	if result.Error == nil {
		productID, err := strconv.ParseInt(c.Params("product"), 10, 64)
		if err != nil {
			return clientError(c, "Wrong product ID format")
		}
		var product data.Product
		result := db.Where("product_id = ?", productID).First(&product)
		if result.Error == nil {
			amount, err := strconv.ParseInt(c.Params("amount"), 10, 64)
			if err != nil {
				return clientError(c, "Wrong order item amount format")
			}
			orderItem := data.OrderItem{
				Product:   product,
				ProductID: product.ProductID,
				Amount:    uint64(amount),
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

// DELETE /orderitem/{itemid}/delete
// Deletes order item
func deleteOrderItem(c *fiber.Ctx) error {
	itemID, err := strconv.ParseInt(c.Params("itemid"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong order item ID format")
	}
	res := db.Where("order_item_id = ?", itemID).Delete(data.OrderItem{})
	if res.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
		})
	}
	return clientError(c, "Order item does not exist")
}

// PUT /order/{orderid}/confirm
// Confirms order
func confirmOrder(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("orderid"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong order item ID format")
	}
	res := db.Model(&data.Order{}).Where("order_id = ?", orderID).Update("confirmed", true)
	if res.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
		})
	}
	return clientError(c, "Order does not exist")
}

// PUT /orderitem/{itemid}/amount/{amount}
// Changes amount of items of order item
func changeOrderItemAmount(c *fiber.Ctx) error {
	itemID, err := strconv.ParseInt(c.Params("itemid"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong order item ID format")
	}
	amount, err := strconv.ParseInt(c.Params("amount"), 10, 64)
	if err != nil {
		return clientError(c, "Wrong order item amount format")
	}
	res := db.Model(&data.OrderItem{}).Where("order_item_id = ?", itemID).Update("amount", amount)
	if res.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  fiber.StatusOK,
			"success": true,
		})
	}
	return clientError(c, "Order item does not exist")
}
