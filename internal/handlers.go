package internal

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api

	"strconv" // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

// a custom response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// Define a HomePage handler that will show the homepage of the blog
func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: HomePageHandler")
	// if client uses anything such as "/abc", "/tag", then return error
	if r.URL.Path != "/" {
		http.NotFound(w, r) // 404 page not found
	}
	w.Write([]byte("Welcome Golang"))

}

// CreateUser create a user in the postgres db
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateUserHandler")

	w.Header().Set("Content-Type", "application/json")

	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	// `User` dfined in model.go is in Uppercase. It means it is public or exported. It can be accessed by the other packages
	var user User

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// call CreateUser function and pass the user
	insertID := CreateUser(user)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetUserHandler")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getUser function with user id to retrieve a single user
	user, err := GetUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(user)
}

// GetAllUser will return all the users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAllUsersHandler")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	users, err := GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

// UpdateUser update user's detail in the postgres db
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateUserHandler")

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty user of type models.User:
	var user User
	/* --- from here, i tried to update a single field  */
	// What if we need to update a only one field/variable  of a row in database
	//  to update a partial field of the User struct, we will declare an
	// user_input struct to hold the expected data from the client.

	var user_input struct { // the format must be same that we have in User struct in models.go file
		ID       int64   `json:"id"`
		Name     *string `json:"name"`
		Location *string `json:"location"`
		Age      *int64  `json:"age"`
	}
	// decode the JSON request to user
	err = json.NewDecoder(r.Body).Decode(&user_input)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// If the user_input.Name value is nil then we know that no corresponding "Name" key/value pair
	// was provided in the JSON request body. So we move on and leave the
	// user.Name value unchanged. Otherwise, we update the user.Name value with the new `Name``
	// value. Importantly, because user_input.Name is a now a pointer to a string, we need
	// to dereference the pointer using the * operator to get the underlying value
	// before assigning it to our user record
	if user_input.Name != nil {
		user.Name = *user_input.Name
	}

	// We also do the same for the other fields in the input struct.
	if user_input.Location != nil {
		user.Location = *user_input.Location
	}
	if user_input.Age != nil {
		user.Age = *user_input.Age
	}
	/* But seems like the partial field codes are not working... */

	// call `UpdateUser` sql query to get the values from the request body and update them
	// to the appropriate fields of the user record in user  table of PSQL.
	updatedRows := UpdateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteUserHandler")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteUser, convert the int to int64
	deletedRows := DeleteUser(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
