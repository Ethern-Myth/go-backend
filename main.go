package main

import (
	"backend_sqlite/controllers/user"
	"backend_sqlite/data"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()
	//For Production
	// gin.SetMode(gin.ReleaseMode)

	data.ConnectDatabase()
	server.GET("/user", user.GetUsers)
	server.GET("/user/:id", user.GetUser)
	server.POST("/user", user.CreateUser)
	server.PUT("/user/:id",user.UpdateUser)
	server.DELETE("/user/:id",user.DeleteUser)

	server.Run(":5000")
}