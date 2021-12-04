package databases

import (
	"database/sql" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// create connection with postgres db
func CreateConnection() *sql.DB {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open the connection - read the keys from your .env file or app.env file using `ConfigStruct` struct
	//var config ConfigStruct
	db, err := sql.Open(config.DB_USERNAME, config.DB_POSTGRES_URL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err != nil {
		panic(err)
	}

	/*Hoping our connection details are validated, next we’re going to call Ping() method on sql.DB object to test our connection.
	  db.ping() will force open a database connection to confirm if we are successfully connected to the database.*/
	err = db.Ping() // check the connection

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// return the connection
	return db
}
