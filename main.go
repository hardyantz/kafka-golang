package main

import (
	"fmt"
	"hardyantz/kafka-golang/producer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("server producer ready from API...")
	router := gin.Default()
	router.POST("/send-email", SendEmail)
	router.Run("localhost:8088")
}

func SendEmail(c *gin.Context) {
	var newMail producer.Message
	response := map[string]string{}

	response["status"] = "success"

	if err := c.BindJSON(&newMail); err != nil {
		response["error"] = err.Error()
		c.IndentedJSON(http.StatusCreated, response)
		return
	}

	err := producer.SendMessage(newMail)
	if err != nil {
		response["error"] = err.Error()
		c.IndentedJSON(http.StatusCreated, response)
		return
	}
	c.IndentedJSON(http.StatusCreated, response)
}
