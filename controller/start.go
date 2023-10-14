package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	id := c.Param("id")
	if id == "1" {
		fmt.Println("1")
	} else {
		fmt.Print("2")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
