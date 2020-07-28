package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gregless22/lab/models"
)

// Response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// UserPersist interface
type UserPersist interface {
	CreateUser(user models.User) int64
}

// Database will save the data
type Database struct {
}

// Init .. initialises the database tables
func Init() {

	// get the sql statement to intialise the table
	sqlStatement := models.InitUserCommand()

	// create the postgres db connection
	db := connect()
	defer db.Close()

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

}

// Connect gets connection details from env variables and returns a pointer to the database
func connect() (db *sql.DB) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		port = 5432
		log.Println("Error setting PORT from env, default used.")
	}

	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "localhost"
		log.Println("Error setting HOST from env, default used.")
	}
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
		log.Println("Error setting USER from env, default used.")
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {

		password = "password"
		log.Println("Error setting PASSWORD from env, default used.")
	}
	dbname, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		dbname = "lab"
		log.Println("Error setting DATABASE from env, default used.")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	fmt.Println(psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to Database %s", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging Database %s", err)
		return
	}

	fmt.Println("Postgres connected successfully")
	return
}

// CreateUser accepts a User struct and returns the ID if successfully returned
func (d Database) CreateUser(user models.User) int64 {

	// create the postgres db connection
	db := connect()
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// GetUser returns a user and error from the database by its userid
func (d Database) GetUser(id int64) (models.User, error) {
	// create the postgres db connection
	db := connect()
	defer db.Close()

	// create a user of models.User type
	var user models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// GetAllUsers returns all of the users from the DB
func (d Database) GetAllUsers() ([]models.User, error) {
	// create the postgres db connection
	db := connect()
	defer db.Close()

	var users []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err
}

// UpdateUser in the DB, the User is the first arguement adn returns the row that its been delete from
func (d Database) UpdateUser(user models.User) int64 {

	// create the postgres db connection
	db := connect()
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, user.ID, user.Name, user.Location, user.Age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// DeleteUser in the DB takes the user ID and returns the row it is deleted from
func (d Database) DeleteUser(id int64) int64 {

	// create the postgres db connection
	db := connect()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
