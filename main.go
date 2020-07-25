package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// this is required for the database connectio
	_ "github.com/lib/pq"
)

func main() {

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		port = 5432
	}

	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "localhost"
	}
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "none"
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = "default"
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		dbname = "db"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err) // hmm panic TODO
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Postgres connected successfully")

	http.HandleFunc("/", server)
	http.HandleFunc("/loan", loans)

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func server(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello You Have Made it to the server, %s \n", r.Host)
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
