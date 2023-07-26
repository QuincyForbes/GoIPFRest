package main

import (
	"basicmodule/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/tokens/:cid", handlers.GetItemByCID) // Just reference the function, don't call it
	r.GET("/tokens", handlers.GetAllItems)       // Same here

	r.Run() // Default runs on :8080
}

