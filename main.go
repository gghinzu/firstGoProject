package main

import (
	"firstGoProject/internal/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//using gin framework
	router := gin.Default()
	router.GET("/user-list", server.GetUsersHandler)
	//we are using '/' notation in url convention instead of other
	//special characters
	//here, '*' is for capturing values for 'id'
	//it will be shown like that -> http://localhost:3000/get-user-by-id/1
	router.GET("/get-user-by-id/*id", server.GetUsersByIDHandler)
	//validation
	fmt.Println("Server is running on http://localhost:3000")
	//the specific port that we want to work with
	router.Run(":3000")
}
