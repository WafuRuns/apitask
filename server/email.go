package main

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/KawaiiWafu/apitask/data"
)

func sendReminders() {
	// SMTP configuartion
	from := "reminder@example.com"
	password := "123456"
	host := "127.0.0.1"
	port := "587"
	auth := smtp.PlainAuth("", from, password, host)

	// Get time before 7 days
	now := time.Now()
	lastWeek := now.AddDate(0, 0, -7)
	var orders []data.Order

	// Find orders older than 7 days
	// Only unconfirmed (unpaid) orders and orders that weren't reminded before
	res := db.Where("confirmed = ? AND reminded = ? AND created_at <= ?", false, false, lastWeek).Find(&orders)
	if res.RowsAffected > 0 {
		tx := db.Begin()
		for _, order := range orders {
			var customer data.Customer
			res := db.Where("customer_id = ?", order.CustomerID).Find(&customer)
			if res.RowsAffected > 0 {
				to := []string{customer.Email}
				msg := []byte(
					"Hello " + customer.Name + ", your order has not been confirmed for 7 days",
				)
				err := smtp.SendMail(host+":"+port, auth, from, to, msg)
				if err != nil {
					println(order.OrderID)
					fmt.Println(err)
				} else {
					tx.Model(&order).Update("reminded", true)
				}
			}
		}
		tx.Commit()
	}
}
