package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	paymentGroup := r.Group("/api")
	{
		paymentGroup.GET("/payment", handlePayment)
	}

	r.Run(":9000")

}

func handlePayment(c *gin.Context) {}
