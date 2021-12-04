package internal

import (
	"database/sql" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"

	"github.com/ajeetyadav1991/go_postgres_crud/databases"

	_ "github.com/lib/pq" // postgres golang driver
)

// based on User struct defined in `models.go`

//---   DB queries and functions . Note that a `users` table is created in PSQL database ------
// insert one user in the `users` table in DB
func CreateUser(user User) int64 {

	// create the postgres db connection
	db := databases.CreateConnection() // CreateConnection function defined in db_migrate.go in databases package

	// close the db connection - defer statement run at the end of the function.
	defer db.Close()

	// create the insert sql query to add rows into `users` table of PSQL database
	// returning userid will return the id of the inserted user
	sqlStatement := "INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid"
	/*Not passing `userid` because `userid` is `SERIAL` type. Its range is 1 to 2,147,483,647. With each insertion it will increment.
	`RETURNING userid` means once it insert successfully in the db, return the `userid`.*/

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)
	// Using `Scan`, the `RETURNING userid` will decode to `id`.

	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
		//log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v \n", id)

	// return the inserted id
	return id
}

// get one user from the DB by its userid
func GetUser(id int64) (User, error) {
	// create the postgres db connection
	db := databases.CreateConnection() // CreateConnection function defined in db_migrate.go in databases package

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var user User

	// create the select sql query
	sqlStatement := "SELECT * FROM users WHERE userid=$1"

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

// get one user from the DB by its userid
func GetAllUsers() ([]User, error) {
	// create the postgres db connection
	db := databases.CreateConnection() // CreateConnection function defined in db_migrate.go in databases package

	// close the db connection
	defer db.Close()

	var users []User

	// create the select sql query
	sqlStatement := "SELECT * FROM users"

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user User

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

// update user in the DB
func UpdateUser(id int64, user User) int64 {
	// This star in *User is for the handlers.go file code: `updatedRows := UpdateUser(int64(id), &user)` for &user

	// create the postgres db connection
	db := databases.CreateConnection() // CreateConnection function defined in db_migrate.go in databases package

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := "UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1"

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

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

// delete user in the DB
func DeleteUser(id int64) int64 {

	// create the postgres db connection
	db := databases.CreateConnection() // CreateConnection function defined in db_migrate.go in databases package

	// close the db connection
	defer db.Close()

	// create the delete sql query from `users` atable created in PSQL database
	sqlStatement := "DELETE FROM users WHERE userid=$1"

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
