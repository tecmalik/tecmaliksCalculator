package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/calculate", calculate)

	router.Run("localhost:9090")
}
