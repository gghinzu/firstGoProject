package main

import (
	"firstGoProject/internal/server"
	"fmt"
	"net/http"
)

func main() {
	//go automatically works on localhost, we assigned the 3000th port to it, and we provided the path here
	//invoking the handler functions in server package
	http.HandleFunc("/user-list", server.GetUsersHandler)
	http.HandleFunc("/get-user-by-id", server.GetUsersByIDHandler)

	//validation
	fmt.Println("Server is running on http://localhost:3000")

	//handling the error in starting
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("An error occurred while starting the server.", err)
	}
}
