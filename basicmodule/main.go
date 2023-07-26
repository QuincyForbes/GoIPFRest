package main

import (
	// Custom module import
	"basicmodule/handlers"

	// Third-party package import
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router
	r := gin.Default()

	// Define route for getting an item by CID
	r.GET("/tokens/:cid", handlers.GetItemByCID) 

	// Define route for getting all items
	r.GET("/tokens", handlers.GetAllItems)       

	// Start the gin server on the default port (8080)
	r.Run() 
}
