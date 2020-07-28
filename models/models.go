package models

// User schema for the database
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}

// InitUserCommand will return the psql command to insert in to the database
func InitUserCommand() string {
	return `CREATE TABLE if NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		age INT,
		name TEXT UNIQUE NOT NULL,
		location TEXT )
		`
}
