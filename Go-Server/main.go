package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"main/routers"
)

func Server() {
	router := gin.Default()
	entrance := router.Group("/entrance", routers.Cors())
	{
		entrance.POST("/login", routers.Login)
		entrance.POST("/register", routers.Regsiter)
	}
	home := router.Group("/api", routers.BasicHandle())
	{
		home.GET("/user", routers.GetData)
	}
	err := router.Run(":8080")
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	Server()
}
