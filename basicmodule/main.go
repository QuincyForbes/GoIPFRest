package main

import (
	"basicmodule/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/tokens/:cid", handlers.GetItemByCID) 
	r.GET("/tokens", handlers.GetAllItems)       

	r.Run() 
}

