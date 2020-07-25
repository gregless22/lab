package main

import (
	"fmt"
	"log"
	"net/http"

	// this is required for the database connectio
	_ "github.com/lib/pq"
)

type server struct {
	db *database
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/loan", loans)

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func home(w http.ResponseWriter, r *http.Request) {

	// sqlStatement := `CREATE TABLE users ( id SERIAL, age INT, first_name VARCHAR(255), last_name VARCHAR(255), email TEXT );`

	// _, err := db.Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }

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
