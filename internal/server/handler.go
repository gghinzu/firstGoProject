package server

import (
	"firstGoProject/internal/person"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// the url that the program is working on
// we don't need to define it explicitly because goland already works with localhost as host by default
// const myUrl = "http://localhost:3000/user-list"

// we will get all the users and display them as they are with this function
// using get method
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//for simple text output, doesn't contain any formatting
		w.Header().Set("Content-Type", "text/plain")
		for _, p := range person.GetAllUsers() {
			fmt.Fprintf(w, "Full name: %s %s\n\n", p.Name, p.Surname)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// we will get just only the given user & their details with this function
func GetUsersByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// getting the parameters from the url
		// here is the get-user-by-id?id=1 version of the url
		queryParams := r.URL.Query()
		// extracting only id value from the url
		idStr := queryParams.Get("id")
		// if there's no id, display all the users again, redirecting user-list
		if idStr == "" {
			http.Redirect(w, r, "/user-list", http.StatusSeeOther)
			return
		}
		// converting the given string into a numeric value (int) with strconv.Atoi
		// also deleting the whitespaces with strings.TrimSpace
		ID, err := strconv.Atoi(strings.TrimSpace(idStr))
		// checking the error and if there's an error, displaying an error msg
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		// if id is okay, send this id to GetUserDetail method & try to find given id
		detail, err := person.GetUserDetail(ID)
		// checking the error and if there's an error, displaying an error msg
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		// displaying the outputs as plain text
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Here are the user's details:\n\n")
		fmt.Fprintf(w, "Full name: %s %s\n", detail.Name, detail.Surname)
		fmt.Fprintf(w, "Age: %d\n", detail.Age)
		fmt.Fprintf(w, "Gender: %s\n", detail.Gender)
		fmt.Fprintf(w, "Education: %s\n", detail.Education)
	} else {
		// if the method is not GET
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
