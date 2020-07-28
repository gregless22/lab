package main

import (
	"fmt"
	"log"
	"net/http"

	// this is required for the database connectio

	"github.com/gregless22/lab/database"
	_ "github.com/lib/pq"
)

func main() {
	// initilise the database
	database.Init()

	http.HandleFunc("/", home)
	http.HandleFunc("/loan", loans)

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	db := database.Database{}
	results, err := db.GetAllUsers()
	// this is where unpacking will occur
	if err != nil {
		fmt.Printf("Database output %s \n", err)
	}
	fmt.Printf("Database output %+v \n", results)
	fmt.Fprintf(w, "Hello You Have Made it to the server")
}

func loans(w http.ResponseWriter, r *http.Request) {

	// switch the cases
	switch r.Method {
	case "GET":
		//get the data from the database
		fmt.Fprintf(w, "get method")
	case "POST":
		//add new data
		fmt.Fprintf(w, "post method")
	}

}
